package main

import (
	"errors"
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/ec2"
	"log"
	"os"
	"strings"
	"time"
)

type Backupper struct {
	config           Config
	client           *ec2.EC2
	launchTestClient *ec2.EC2
	BackupInstances  BackupInstances
}

func NewBackupper(config Config, client *ec2.EC2, launchTestClient *ec2.EC2) *Backupper {
	return &Backupper{config: config, client: client, launchTestClient: launchTestClient}
}

func (h *Backupper) Run() error {
	fmt.Println("Getting instance information...")
	err := h.getBackupTargetInstances(h.config.BackupTargetConfig.BackupTargetFilers)
	if err != nil {
		return err
	}
	h.displayInstances()

	fmt.Println("Stop instances to get AMIs. Do you want to run it? yes/no")
	fmt.Print("Enter a value: ")
	ok := askForConfirmation()
	if !ok {
		os.Exit(0)
	}

	fmt.Println("Stopping instances...")
	err = h.stopInstances()
	if err != nil {
		return err
	}

	fmt.Println("Waiting for all instances to be stopped...")
	wait(NewInstanceChecker(h.client, h.BackupInstances.ids, ec2.InstanceStateNameStopped))

	fmt.Println("Creating AMIs...")
	err = h.createAmi()
	if err != nil {
		return err
	}

	fmt.Println("Waiting for AMIs to be available...")
	wait(NewImageStateChecker(h.client, *h.BackupInstances.imageIds))

	fmt.Println("Adding name tag to snapshots...")
	err = h.addNameTagToSnapShots(*h.BackupInstances.imageIds)
	if err != nil {
		return err
	}

	fmt.Println("Start instances...")
	err = h.startInstances()
	if err != nil {
		return err
	}

	fmt.Println("Waiting for the state of all instances to be running...")
	wait(NewInstanceChecker(h.client, h.BackupInstances.ids, ec2.InstanceStateNameRunning))
	fmt.Println("All instances running")

	fmt.Println("Adding AMI launch permissions to account...")
	h.addPermissionToAmi(*h.BackupInstances.imageIds, h.config.AmiLaunchTestConfig.LaunchPermissionAddUserId)

	fmt.Println("Launching AMI with account " + h.config.AmiLaunchTestConfig.Profile)
	h.runLaunchTestInstance()

	fmt.Println("Finished")

	return nil
}

func (h *Backupper) runLaunchTestInstancesInput(instance *BackupInstance) *ec2.RunInstancesInput {
	return &ec2.RunInstancesInput{
		ImageId:          instance.imageId,
		InstanceType:     aws.String(h.config.AmiLaunchTestConfig.InstanceType),
		KeyName:          aws.String(h.config.AmiLaunchTestConfig.KeyName),
		MaxCount:         aws.Int64(1),
		MinCount:         aws.Int64(1),
		SecurityGroupIds: h.config.AmiLaunchTestConfig.SecurityGroupIds,
		SubnetId:         aws.String(h.config.AmiLaunchTestConfig.SubnetId),
		TagSpecifications: []*ec2.TagSpecification{
			{
				ResourceType: aws.String("instance"),
				Tags: []*ec2.Tag{
					{
						Key:   aws.String("Name"),
						Value: aws.String("launch_test_" + instance.name),
					},
					{
						Key:   aws.String("LaunchTest"),
						Value: aws.String("true" + instance.name),
					},
				},
			},
			{
				ResourceType: aws.String("volume"),
				Tags: []*ec2.Tag{
					{
						Key:   aws.String("Name"),
						Value: aws.String("launch_test_" + instance.name),
					},
					{
						Key:   aws.String("LaunchTest"),
						Value: aws.String("true" + instance.name),
					},
				},
			},
		},
	}
}
func (h *Backupper) runLaunchTestInstance() {
	const suffix = ".private.test"
	const strRepeatLength = 40
	var instanceIds []*string

	fmt.Println(strings.Repeat("-", strRepeatLength))
	for _, i := range h.BackupInstances.instances {
		reservation, err := h.launchTestClient.RunInstances(h.runLaunchTestInstancesInput(i))
		if err != nil {
			fmt.Println(err)
		}

		for _, i := range reservation.Instances {
			fmt.Printf("Host %v\n  HostName %v\n", getTagValue(i.Tags, "Name")+suffix, *i.PrivateIpAddress)
			instanceIds = append(instanceIds, i.InstanceId)
		}

	}
	fmt.Println(strings.Repeat("-", strRepeatLength))

	fmt.Println("Waiting for the state of all test instances to be running...")
	wait(NewInstanceChecker(h.launchTestClient, instanceIds, ec2.InstanceStateNameRunning))

}

func (h *Backupper) getBackupTargetInstances(filters []*ec2.Filter) error {
	output, err := h.client.DescribeInstances(&ec2.DescribeInstancesInput{
		Filters: filters,
	})
	if err != nil {
		return err
	}

	for _, r := range output.Reservations {
		for _, i := range r.Instances {
			instance := BackupInstance{
				id:                   *i.InstanceId,
				name:                 getAmiName(getTagValue(i.Tags, "Name")),
				backupExcludeVolumes: getBackupExcludeVolumes(i.Tags),
			}

			h.BackupInstances.instances = append(h.BackupInstances.instances, &instance)
			h.BackupInstances.ids = append(h.BackupInstances.ids, i.InstanceId)
		}
	}

	return nil
}

func (h *Backupper) addPermissionToAmi(imageIds ImageIds, userId string) {
	for _, imageId := range imageIds {
		_, err := h.client.ModifyImageAttribute(&ec2.ModifyImageAttributeInput{
			ImageId: imageId,
			LaunchPermission: &ec2.LaunchPermissionModifications{
				Add: []*ec2.LaunchPermission{{
					UserId: aws.String(userId),
				}},
			},
		})
		if err != nil {
			fmt.Printf("err: %v\n", err.Error())
			continue
		}
		fmt.Printf("image_id: %v user_id: %v\n", *imageId, userId)
	}
}

func (h *Backupper) stopInstances() error {
	_, err := h.client.StopInstances(&ec2.StopInstancesInput{
		InstanceIds: h.BackupInstances.ids,
	})
	return err
}

func wait(checker StateChecker) {
	const waitInterval = 3 * time.Second
	for {
		ok, err := checker.Check()
		if ok {
			fmt.Print("\n")
			break
		}
		if err != nil {
			log.Println(err)
		}

		fmt.Print("*")
		time.Sleep(waitInterval)
	}
}

func (h *Backupper) createAmi() error {

	var imageIds ImageIds
	for _, i := range h.BackupInstances.instances {

		var mapping []*ec2.BlockDeviceMapping
		for _, v := range i.backupExcludeVolumes {
			m := &ec2.BlockDeviceMapping{
				DeviceName: aws.String(v),
				NoDevice:   aws.String(""),
			}
			mapping = append(mapping, m)
		}

		output, err := h.client.CreateImage(&ec2.CreateImageInput{
			BlockDeviceMappings: mapping,
			InstanceId:          aws.String(i.id),
			Name:                aws.String(i.name),
			TagSpecifications: []*ec2.TagSpecification{{
				ResourceType: aws.String("image"),
				Tags: []*ec2.Tag{{
					Key:   aws.String("Name"),
					Value: aws.String(i.name),
				}},
			}},
		})
		if err != nil {
			return err
		}
		imageIds = append(imageIds, output.ImageId)
		i.imageId = output.ImageId

		fmt.Println("id:" + *output.ImageId + " name: " + i.name)
	}
	h.BackupInstances.imageIds = &imageIds

	return nil
}

func (h *Backupper) addTagsToAmi(imageIds ImageIds, tags []*ec2.Tag) error {
	_, err := h.client.CreateTags(&ec2.CreateTagsInput{
		Resources: imageIds,
		Tags:      tags,
	})
	if err != nil {
		return err
	}
	return nil
}

func (h *Backupper) startInstances() error {
	input := &ec2.StartInstancesInput{
		InstanceIds: h.BackupInstances.ids,
	}
	_, err := h.client.StartInstances(input)
	if err != nil {
		return err
	}

	return nil
}

func (h *Backupper) addNameTagToSnapShots(imageIds ImageIds) error {
	o, err := h.client.DescribeImages(&ec2.DescribeImagesInput{
		ImageIds: imageIds,
	})
	if err != nil {
		return err
	}

	for _, i := range o.Images {
		for _, mapping := range i.BlockDeviceMappings {
			// SnapshotId will be nil before AMI is available
			if mapping.Ebs.SnapshotId == nil {
				return errors.New("there was no snapshot")
			}
			err := h.addTagsToAmi(ImageIds{
				mapping.Ebs.SnapshotId,
			}, []*ec2.Tag{
				{
					Key:   aws.String("Name"),
					Value: aws.String(getTagValue(i.Tags, "Name")),
				},
			})
			if err != nil {
				return err
			}

			fmt.Println("id: " + *mapping.Ebs.SnapshotId + " Name: " + getTagValue(i.Tags, "Name"))
		}
	}
	return nil
}

func (h *Backupper) displayInstances() {
	for _, i := range h.BackupInstances.instances {
		fmt.Printf("%+v\n", i)
	}
}

func askForConfirmation() bool {
	var response string

	_, err := fmt.Scanln(&response)
	if err != nil {
		log.Fatal(err)
	}

	switch strings.ToLower(response) {
	case "yes":
		return true
	case "no":
		return false
	default:
		fmt.Println("I'm sorry but I didn't get what you meant, please type yes or no and then press enter:")
		return askForConfirmation()
	}
}

func getTagValue(tags []*ec2.Tag, key string) string {
	for _, tag := range tags {
		if *tag.Key == key {
			return *tag.Value
		}
	}
	return ""
}

func getBackupExcludeVolumes(tags []*ec2.Tag) []string {
	v := getTagValue(tags, "BackupExcludeVolume")
	return strings.Split(v, ",")
}
func getAmiName(instanceName string) string {
	return instanceName + "_" + time.Now().Format("2006-01-02-1504")
}
