// Copyright 2020 The Inflion Authors.
//
// Use of this software is governed by the Business Source License
// included in the file licenses/BSL.txt.
//
// As of the Change Date specified in that file, in accordance with
// the Business Source License, use of this software will be governed
// by the Apache License, Version 2.0, included in the file
// licenses/APL.txt.

package paws

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ec2"
	"strings"
)

const fullOpen = "0.0.0.0/0"

type AwsSecurityGroup struct {
	ec2 *ec2.EC2
}

type OpenPorts struct {
	ports []int64
}

func (o *OpenPorts) ToString() string {
	p := []string{}
	for _, v := range o.ports {
		p = append(p, string(v))
	}

	log.Println(p)
	return strings.Join(p, ",")
}

func (o *OpenPorts) Length() int {
	return len(o.ports)
}

type PawsSg struct {
	Id               string
	Name             string
	rawSecurityGroup *ec2.SecurityGroup
}

func (s *PawsSg) GetOpenPorts() *OpenPorts {
	var openPorts []int64
	for _, p := range s.rawSecurityGroup.IpPermissions {
		for _, r := range p.IpRanges {
			if *r.CidrIp == fullOpen {
				openPorts = append(openPorts, *p.ToPort)
			}
		}
	}
	return &OpenPorts{ports: openPorts}
}

func (s *PawsSg) HasOpenPorts() bool {
	return s.GetOpenPorts().Length() != 0
}

func NewAwsSecurityGroup(awsAccount AwsAccount, region string) AwsSecurityGroup {
	sess, err := session.NewSession()
	if err != nil {
		log.Println(err)
	}

	conf := CreateConfig(awsAccount, region, sess)
	return AwsSecurityGroup{ec2: ec2.New(sess, &conf)}
}

func (s *AwsSecurityGroup) GetSecurityGroups() ([]PawsSg, error) {
	var securityGroups []PawsSg

	result, err := s.ec2.DescribeSecurityGroups(nil)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			default:
				log.Println(aerr.Error())
			}
		} else {
			log.Println(err.Error())
		}
		return nil, err
	}

	for _, sg := range result.SecurityGroups {
		securityGroups = append(securityGroups,
			PawsSg{
				Id:               aws.StringValue(sg.GroupId),
				Name:             aws.StringValue(sg.GroupName),
				rawSecurityGroup: sg,
			})
	}

	return securityGroups, nil
}
