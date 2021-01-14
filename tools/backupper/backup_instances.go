package main

type BackupInstances struct {
	ids       BackupInstanceIds
	imageIds  *ImageIds
	instances []*BackupInstance
}

type BackupInstanceIds []*string

type BackupInstance struct {
	id                   string
	name                 string
	backupExcludeVolumes []string
	imageId              *string
}
