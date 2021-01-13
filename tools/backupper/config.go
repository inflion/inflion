package main

import (
	"errors"
	"github.com/aws/aws-sdk-go/service/ec2"
	"gopkg.in/yaml.v3"
	"io/ioutil"
)

type Config struct {
	BackupTargetFilers        []*ec2.Filter `yaml:"BackupTargetFilers"`
	AmiTags                   []*ec2.Tag    `yaml:"AmiTags"`
	LaunchPermissionAddUserId string        `yaml:"LaunchPermissionAddUserId"`
}

func newConf(confPath *string) (*Config, error) {
	if *confPath == "" {
		return nil, errors.New("config file required, but not set")
	}

	buf, err := ioutil.ReadFile(*confPath)
	if err != nil {
		return nil, errors.New("failed to read confPath file")
	}

	conf := &Config{}
	err = yaml.Unmarshal(buf, conf)
	if err != nil {
		return nil, errors.New("failed to read config file")
	}

	return conf, nil
}
