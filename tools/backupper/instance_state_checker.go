package main

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/ec2"
)

type InstanceStateChecker struct {
	client            *ec2.EC2
	instanceIDs       []*string
	instanceStateName string
}

func NewInstanceChecker(client *ec2.EC2, instanceIDs []*string, instanceStateName string) *InstanceStateChecker {
	return &InstanceStateChecker{client: client, instanceIDs: instanceIDs, instanceStateName: instanceStateName}
}

func (c *InstanceStateChecker) Check() (ok bool, err error) {
	output, err := c.client.DescribeInstances(&ec2.DescribeInstancesInput{
		Filters: []*ec2.Filter{
			{
				Name:   aws.String("instance-state-name"),
				Values: []*string{aws.String(c.instanceStateName)},
			},
			{
				Name:   aws.String("instance-id"),
				Values: c.instanceIDs,
			},
		},
		InstanceIds: nil,
	})
	if err != nil {
		return false, err
	}

	var gotIds BackupInstanceIds
	for _, r := range output.Reservations {
		for _, i := range r.Instances {
			gotIds = append(gotIds, i.InstanceId)
		}
	}

	return len(c.instanceIDs) == len(gotIds), nil
}
