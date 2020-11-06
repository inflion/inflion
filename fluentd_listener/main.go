package main

import (
	"bufio"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	pb "github.com/inflion/inflion/inflionserver/byteevent/byteeventpb"
	"github.com/inflion/inflion/internal/logger"
	"github.com/spf13/cobra"
	"google.golang.org/grpc"
	_log "log"
	"net/http"
	"os"
)

var (
	name    = "fluentd-listener"
	version = "master"
	rootCmd = &cobra.Command{
		Use:     name,
		Short:   name,
		RunE:    runRootCmd,
		Version: version,
	}
	log     logger.Logger
	project string // NOTE: The term project will be abolished, so this variable is tentatively defined as a package scope variable.
)

func init() {
	var err error
	log, err = logger.NewZapLogger(&logger.Configuration{Level: logger.InfoLevel})
	if err != nil {
		_log.Fatalln(err.Error())
	}
}

func init() {
	rootCmd.Flags().String("endpoint", "127.0.0.1:50051", "inflionserver address")
	rootCmd.Flags().String("listen", "0.0.0.0:8000", "listen address")
	rootCmd.Flags().String("project", "", "project name [required]")
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}

func runRootCmd(cmd *cobra.Command, _ []string) error {
	endpoint, err := cmd.Flags().GetString("endpoint")
	if err != nil {
		return err
	}
	addr, err := cmd.Flags().GetString("listen")
	if err != nil {
		return err
	}
	project, err = cmd.Flags().GetString("project")
	if err != nil {
		return err
	}
	if len(project) == 0 {
		return errors.New("project flag is required and must not be empty")
	}
	logRecordHandler, err := newInflionLogRecordHandler(endpoint)
	if err != nil {
		return err
	}
	mux := newServeMux(logRecordHandler.sendToInflionServer, log)
	log.Info(fmt.Sprintf("Listening on address %s", addr))
	return http.ListenAndServe(addr, mux)
}

type logRecord = map[string]interface{}
type logRecordHandler = func(r logRecord)

func newServeMux(handleLogRecord logRecordHandler, log logger.Logger) *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
		log.InfoWith("request", "method", req.Method, "path", req.RequestURI)
		if req.Method != http.MethodPost {
			res.WriteHeader(http.StatusMethodNotAllowed)
			//goland:noinspection GoUnhandledErrorResult
			res.Write([]byte("Only POST method supported."))
			return
		}
		if req.Header.Get("Content-Type") != "application/x-ndjson" {
			res.WriteHeader(http.StatusBadRequest)
			//goland:noinspection GoUnhandledErrorResult
			res.Write([]byte("Only content-type 'application/x-ndjson' supported."))
			return
		}

		scanner := bufio.NewScanner(req.Body)
		failedAtLeastOnce := false
		succeededAtLeastOnce := false
		type lineError = struct {
			record []byte
			err    error
		}
		var lineErrors []lineError

		// NOTE: json.Decoder.More() cannot be used this `for` statement, since it hangs when a record that cannot be parsed as JSON comes.
		for scanner.Scan() {
			line := scanner.Bytes()
			if len(line) == 0 {
				continue
			}

			var logRecord logRecord
			if err := json.Unmarshal(line, &logRecord); err != nil {
				failedAtLeastOnce = true
				lineErrors = append(lineErrors, lineError{record: line, err: err})
			} else {
				succeededAtLeastOnce = true
				handleLogRecord(logRecord)
			}
		}

		statusCode := http.StatusAccepted // Because the logRecordHandler might be asynchronous.
		if failedAtLeastOnce {
			if succeededAtLeastOnce {
				statusCode = 207
				for _, lineError := range lineErrors {
					log.ErrorWith(
						"Malformed log record found",
						"record", string(lineError.record),
						"err", lineError.err.Error(),
					)
				}
			} else {
				statusCode = http.StatusBadRequest
			}
		}
		res.WriteHeader(statusCode)
	})
	return mux
}

type inflionLogRecordHandler struct {
	con    *grpc.ClientConn
	client pb.ByteEventClient
}

func newInflionLogRecordHandler(endpoint string) (*inflionLogRecordHandler, error) {
	log.Info(fmt.Sprintf("Connecting to gRPC endpoint %s", endpoint))
	con, err := grpc.Dial(endpoint, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}
	log.Info("Established gRPC connection")
	client := pb.NewByteEventClient(con)
	return &inflionLogRecordHandler{con: con, client: client}, nil
}

func (h *inflionLogRecordHandler) sendToInflionServer(record map[string]interface{}) {
	event, err := json.Marshal(record)
	if err != nil {
		log.Error(err)
		return
	}
	res, err := h.client.Put(context.Background(), &pb.PutByteEventRequest{Project: project, Event: event})
	if err != nil {
		log.ErrorWith("Failed to send event", "err", err, "event", event)
		return
	}
	log.InfoWith("Event sent", "res", res)
}

func (h *inflionLogRecordHandler) close() error {
	log.Info("Closing gRPC connection")
	return h.con.Close()
}
