package main

import (
	"flag"
	"fmt"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ec2"
	"os"
)

var confPath = flag.String("config", "", "a path of config yaml file")

func main() {
	flag.Parse()

	conf, err := newConf(confPath)
	if err != nil {
		os.Exit(2)
	}

	err = NewBackupper(
		*conf,
		ec2.New(session.Must(session.NewSessionWithOptions(session.Options{
			Profile:           conf.BackupTargetConfig.Profile,
			SharedConfigState: session.SharedConfigEnable,
		}))),
		ec2.New(session.Must(session.NewSessionWithOptions(session.Options{
			Profile:           conf.AmiLaunchTestConfig.Profile,
			SharedConfigState: session.SharedConfigEnable,
		})))).Run()
	if err != nil {
		fmt.Printf("err: %v", err)
		os.Exit(2)
	}
}
