package main

import (
	"errors"
	"github.com/aws/aws-sdk-go/service/ec2"
	"gopkg.in/yaml.v3"
	"io/ioutil"
)

type Config struct {
	BackupTargetConfig  BackupTargetConfig  `yaml:"BackupTargetConfig"`
	AmiLaunchTestConfig AmiLaunchTestConfig `yaml:"AMILaunchTestConfig"`
}

type BackupTargetConfig struct {
	Profile            string        `yaml:"profile"`
	BackupTargetFilers []*ec2.Filter `yaml:"target_filers"`
	AmiTags            []*ec2.Tag    `yaml:"ami_tags"`
}

type AmiLaunchTestConfig struct {
	Profile                   string    `yaml:"profile"`
	LaunchPermissionAddUserId string    `yaml:"launch_permission_add_user_id"`
	InstanceType              string    `yaml:"instance_type"`
	SubnetId                  string    `yaml:"subnet_id"`
	KeyName                   string    `yaml:"key_name"`
	SecurityGroupIds          []*string `yaml:"security_group_ids"`
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
