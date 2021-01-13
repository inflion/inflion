package main

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/ec2"
)

type ImageStateChecker struct {
	client         *ec2.EC2
	targetImageIds ImageIds
}
type ImageIds []*string

func NewImageStateChecker(client *ec2.EC2, targetImageIds ImageIds) *ImageStateChecker {
	return &ImageStateChecker{client: client, targetImageIds: targetImageIds}
}

func (c *ImageStateChecker) Check() (bool, error) {
	output, err := c.client.DescribeImages(&ec2.DescribeImagesInput{
		Filters: []*ec2.Filter{
			{
				Name:   aws.String("state"),
				Values: []*string{aws.String(ec2.ImageStateAvailable)},
			},
		},
		ImageIds: c.targetImageIds,
	})
	if err != nil {
		return false, err
	}
	return len(output.Images) == len(c.targetImageIds), nil
}
