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
	"errors"
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/service/ec2"
)

type InstanceID *string

type AwsInstance struct {
	InstanceID     string
	Name           string
	PrivateAddress string
	PublicAddress  string
	Tags           Tags
	SecurityGroups []SecurityGroup
	Status         string
}

type CreateInstanceParameter struct {
	Ami    string
	Family string
}

type SecurityGroup struct {
	Id   string
	Name string
}

func createInput(instanceId string) *ec2.DescribeTagsInput {
	input := &ec2.DescribeTagsInput{
		Filters: []*ec2.Filter{
			{
				Name: aws.String("resource-id"),
				Values: []*string{
					aws.String(instanceId),
				},
			},
		},
	}

	return input
}

func (a *Api) GetInstance(instanceId string) (*AwsInstance, error) {
	params := &ec2.DescribeInstancesInput{
		Filters: []*ec2.Filter{
			{
				Name: aws.String("instance-id"),
				Values: []*string{
					aws.String(instanceId),
				},
			},
		},
	}

	result, err := a.conn.DescribeInstances(params)
	if err != nil {
		return nil, err
	}

	instances := a.convertAwsInstances(result.Reservations)
	if len(instances) != 0 {
		return instances[0], nil
	}

	return nil, errors.New("instance not found")
}

func (a *Api) DescribeSecurityGroup(groupIds []*string) {
	input := &ec2.DescribeSecurityGroupsInput{
		GroupIds: groupIds,
	}
	result, err := a.conn.DescribeSecurityGroups(input)

	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			default:
				log.Error(aerr.Error())
			}
		} else {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			log.Error(err.Error())
		}
		return
	}

	fmt.Print(result)
}

type FilterCondition struct {
	All      bool
	TagName  string
	TagValue string
}

func NewEmptyFilterCondition() FilterCondition {
	return FilterCondition{All: true}
}

func (a *Api) GetInstances(cond FilterCondition) ([]*AwsInstance, error) {
	var params *ec2.DescribeInstancesInput
	if cond.All {
		params = nil
	} else {
		params = &ec2.DescribeInstancesInput{
			Filters: []*ec2.Filter{
				{
					Name: aws.String("tag:" + cond.TagName),
					Values: []*string{
						aws.String(cond.TagValue),
					},
				},
			},
		}
	}

	result, err := a.conn.DescribeInstances(params)
	if err != nil {
		return nil, err
	}

	instances := a.convertAwsInstances(result.Reservations)
	return instances, nil
}

func (a *Api) CreateInstance(parameter CreateInstanceParameter) (InstanceID, error) {
	runResult, err := a.conn.RunInstances(&ec2.RunInstancesInput{
		ImageId:      aws.String(parameter.Ami),
		InstanceType: aws.String(parameter.Family),
		MinCount:     aws.Int64(1),
		MaxCount:     aws.Int64(1),
	})

	if err != nil {
		log.Error("could not create instance", err)
		return nil, err
	}

	return runResult.Instances[0].InstanceId, nil
}

func (a *Api) TerminateInstance(instanceId InstanceID) (InstanceID, error) {
	result, err := a.conn.TerminateInstances(&ec2.TerminateInstancesInput{
		DryRun:      aws.Bool(false),
		InstanceIds: []*string{instanceId},
	})
	if err != nil {
		log.Error("could not terminate instance", err)
		return nil, err
	}

	return result.TerminatingInstances[0].InstanceId, nil
}

func (a *Api) convertAwsInstances(reservations []*ec2.Reservation) []*AwsInstance {
	var instances []*AwsInstance

	for _, res := range reservations {
		for _, instance := range res.Instances {
			ec2Instance := AwsInstance{InstanceID: *instance.InstanceId}

			if instance.PrivateIpAddress != nil {
				ec2Instance.PrivateAddress = *instance.PrivateIpAddress
			}
			if instance.PublicIpAddress != nil {
				ec2Instance.PublicAddress = *instance.PublicIpAddress
			}

			ec2Instance.Status = aws.StringValue(instance.State.Name)
			ec2Instance.SecurityGroups = a.convertSecurityGroup(instance.SecurityGroups)

			awsTags, err := a.conn.DescribeTags(createInput(ec2Instance.InstanceID))
			if err != nil {
				log.Error(err)
			}

			tags := a.convertAwsTagsToTags(awsTags)
			ec2Instance.Name = tags.FindValueOrElse("Name", "")
			ec2Instance.Tags = tags

			instances = append(instances, &ec2Instance)
		}
	}

	return instances
}

func (a *Api) convertAwsTagsToTags(awsTags *ec2.DescribeTagsOutput) Tags {
	var tags Tags
	for _, tag := range awsTags.Tags {
		if tag.Key != nil {
			tags.append(Tag{Key: aws.StringValue(tag.Key), Value: aws.StringValue(tag.Value)})
		}
	}
	return tags
}

func (a *Api) convertSecurityGroup(securityGroups []*ec2.GroupIdentifier) []SecurityGroup {
	var result []SecurityGroup

	for _, sg := range securityGroups {
		result = append(result, SecurityGroup{
			Id:   *sg.GroupId,
			Name: *sg.GroupName,
		})
	}
	return result
}
