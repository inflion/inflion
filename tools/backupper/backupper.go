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
	config          Config
	client          *ec2.EC2
	BackupInstances BackupInstances
}

func NewBackupper(config Config, client *ec2.EC2) *Backupper {
	return &Backupper{config: config, client: client}
}

func (h *Backupper) Run() error {
	fmt.Println("Getting instance information...")
	err := h.getBackupTargetInstances(h.config.BackupTargetFilers)
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
	imageIds, err := h.createAmi()
	if err != nil {
		return err
	}

	fmt.Println("Waiting for AMIs to be available...")
	wait(NewImageStateChecker(h.client, imageIds))

	fmt.Println("Adding name tag to snapshots...")
	err = h.addNameTagToSnapShots(imageIds)
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
	h.addPermissionToAmi(imageIds, h.config.LaunchPermissionAddUserId)

	fmt.Println("Finished")

	return nil
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

			h.BackupInstances.instances = append(h.BackupInstances.instances, instance)
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

func (h *Backupper) createAmi() (ImageIds, error) {

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
			return nil, err
		}
		imageIds = append(imageIds, output.ImageId)
		fmt.Println("id:" + *output.ImageId + " name: " + i.name)
	}
	return imageIds, nil
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
