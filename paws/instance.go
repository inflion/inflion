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
	"log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/ec2"
)

type InstanceID *string
type InstanceIds []*string

type AwsInstance struct {
	InstanceID       string
	Name             string
	PrivateAddress   string
	PublicAddress    string
	Tags             Tags
	SecurityGroupIds []string
	Status           string
}

type CreateInstanceParameter struct {
	Ami    string
	Family string
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
		log.Println("could not create instance", err)
		return nil, err
	}

	return runResult.Instances[0].InstanceId, nil
}

func (a *Api) StartInstances(instanceIds InstanceIds) (InstanceIds, error) {
	result, err := a.conn.StartInstances(&ec2.StartInstancesInput{
		AdditionalInfo: nil,
		DryRun:         aws.Bool(false),
		InstanceIds:    instanceIds,
	})
	if err != nil {
		log.Println("could not start instances", err)
		return nil, err
	}

	var startingInstanceIds InstanceIds
	for _, startingInstance := range result.StartingInstances {
		startingInstanceIds = append(startingInstanceIds, startingInstance.InstanceId)
	}

	return startingInstanceIds, nil
}

func (a *Api) StopInstances(instanceIds InstanceIds) (InstanceIds, error) {
	result, err := a.conn.StopInstances(&ec2.StopInstancesInput{
		DryRun:      aws.Bool(false),
		Force:       nil,
		Hibernate:   nil,
		InstanceIds: instanceIds,
	})
	if err != nil {
		log.Println("could not stop instances", err)
		return nil, err
	}

	var stoppingInstanceIds InstanceIds
	for _, stoppingInstance := range result.StoppingInstances {
		stoppingInstanceIds = append(stoppingInstanceIds, stoppingInstance.InstanceId)
	}

	return stoppingInstanceIds, nil
}

func (a *Api) TerminateInstances(instanceIds InstanceIds) (InstanceIds, error) {
	result, err := a.conn.TerminateInstances(&ec2.TerminateInstancesInput{
		DryRun:      aws.Bool(false),
		InstanceIds: instanceIds,
	})
	if err != nil {
		log.Println("could not terminate instances", err)
		return nil, err
	}

	var terminatingInstanceIds InstanceIds
	for _, terminatingInstance := range result.TerminatingInstances {
		terminatingInstanceIds = append(terminatingInstanceIds, terminatingInstance.InstanceId)
	}

	return terminatingInstanceIds, nil
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
			ec2Instance.SecurityGroupIds = a.convertSecurityGroup(instance.SecurityGroups)

			awsTags, err := a.conn.DescribeTags(createInput(ec2Instance.InstanceID))
			if err != nil {
				log.Println(err)
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

func (a *Api) convertSecurityGroup(securityGroups []*ec2.GroupIdentifier) []string {
	var result = []string{}
	for _, sg := range securityGroups {
		result = append(result, aws.StringValue(sg.GroupId))
	}
	return result
}
