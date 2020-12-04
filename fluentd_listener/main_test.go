package main

import (
	"bytes"
	"github.com/inflion/inflion/logger"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func Test_newServeMux(t *testing.T) {
	t.Run("posts two log records", func(t *testing.T) {
		recorder, records := newLogRecorder()
		mux := newServeMux(recorder, logger.MuteLogger)
		req := httptest.NewRequest("POST", "http://127.0.0.1/", bytes.NewBufferString(`{"message": "log 1"}
{"message": "log 2"}`))
		req.Header.Set("Content-Type", "application/x-ndjson")
		w := httptest.NewRecorder()
		mux.ServeHTTP(w, req)
		res := w.Result()
		assert.Equal(t, http.StatusAccepted, res.StatusCode)
		assert.Equal(t, []logRecord{{"message": "log 1"}, {"message": "log 2"}}, *records)
	})

	t.Run("send non-POST request", func(t *testing.T) {
		mux := newServeMux(nullLogHandler, logger.MuteLogger)
		req := httptest.NewRequest("GET", "http://127.0.0.1/", nil)
		w := httptest.NewRecorder()
		mux.ServeHTTP(w, req)
		res := w.Result()
		assert.Equal(t, http.StatusMethodNotAllowed, res.StatusCode)
		body, _ := ioutil.ReadAll(res.Body)
		assert.Equal(t, "Only POST method supported.", string(body))
	})

	t.Run("send non-ndjson content-type", func(t *testing.T) {
		mux := newServeMux(nullLogHandler, logger.MuteLogger)
		req := httptest.NewRequest("POST", "http://127.0.0.1/", bytes.NewBufferString(`{"message": "log 1"}`))
		req.Header.Set("Content-Type", "application/json")
		w := httptest.NewRecorder()
		mux.ServeHTTP(w, req)
		res := w.Result()
		assert.Equal(t, http.StatusBadRequest, res.StatusCode)
		body, _ := ioutil.ReadAll(res.Body)
		assert.Equal(t, "Only content-type 'application/x-ndjson' supported.", string(body))
	})

	t.Run("send partially-broken ndjson", func(t *testing.T) {
		recorder, records := newLogRecorder()
		mux := newServeMux(recorder, logger.MuteLogger)
		req := httptest.NewRequest("POST", "http://127.0.0.1/", bytes.NewBufferString(`{"message": "log 1"}
broken log 2
{"message": "log 3"}`))
		req.Header.Set("Content-Type", "application/x-ndjson")
		w := httptest.NewRecorder()
		mux.ServeHTTP(w, req)
		res := w.Result()
		assert.Equal(t, 207, res.StatusCode)
		assert.Equal(t, []logRecord{{"message": "log 1"}, {"message": "log 3"}}, *records)
	})

	t.Run("send fully-broken ndjson", func(t *testing.T) {
		recorder, records := newLogRecorder()
		mux := newServeMux(recorder, logger.MuteLogger)
		req := httptest.NewRequest("POST", "http://127.0.0.1/", bytes.NewBufferString(`broken log1
broken log2
broken log3`))
		req.Header.Set("Content-Type", "application/x-ndjson")
		w := httptest.NewRecorder()
		mux.ServeHTTP(w, req)
		res := w.Result()
		assert.Equal(t, http.StatusBadRequest, res.StatusCode)
		assert.Empty(t, *records)
	})
}

func newLogRecorder() (logRecordHandler, *[]logRecord) {
	records := &[]logRecord{}
	return func(record logRecord) {
		*records = append(*records, record)
	}, records
}

var nullLogHandler func(r logRecord)
