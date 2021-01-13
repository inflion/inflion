package main

type BackupInstances struct {
	ids       BackupInstanceIds
	instances []BackupInstance
}

type BackupInstanceIds []*string

type BackupInstance struct {
	id                   string
	name                 string
	backupExcludeVolumes []string
}
