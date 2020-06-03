package paws

import (
	"github.com/aws/aws-sdk-go/aws"
	"reflect"
	"testing"
)

func TestApi_StartInstances(t *testing.T) {
	type args struct {
		instanceIds InstanceIds
	}
	tests := []struct {
		name    string
		args    args
		want    InstanceIds
		wantErr bool
	}{
		{
			name: "",
			args: args{
				instanceIds: []*string{
					aws.String(""),
				},
			},
			want: []*string{
				aws.String(""),
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			a, err := New(
				AwsAccount{
					AccountId:  "",
					RoleName:   "",
					ExternalId: "",
				},
				"ap-northeast-1",
			)

			got, err := a.StartInstances(tt.args.instanceIds)
			if (err != nil) != tt.wantErr {
				t.Errorf("StartInstances() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("StartInstances() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestApi_StopInstances(t *testing.T) {
	type args struct {
		instanceIds InstanceIds
	}
	tests := []struct {
		name    string
		args    args
		want    InstanceIds
		wantErr bool
	}{
		{
			name: "",
			args: args{
				instanceIds: []*string{
					aws.String(""),
				},
			},
			want: []*string{
				aws.String(""),
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			a, err := New(
				AwsAccount{
					AccountId:  "",
					RoleName:   "",
					ExternalId: "",
				},
				"ap-northeast-1",
			)

			got, err := a.StopInstances(tt.args.instanceIds)
			if (err != nil) != tt.wantErr {
				t.Errorf("StopInstances() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("StopInstances() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestApi_TerminateInstances(t *testing.T) {
	type args struct {
		instanceIds InstanceIds
	}
	tests := []struct {
		name    string
		args    args
		want    InstanceIds
		wantErr bool
	}{
		{
			name: "",
			args: args{
				instanceIds: []*string{
					aws.String(""),
				},
			},
			want: []*string{
				aws.String(""),
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			a, err := New(
				AwsAccount{
					AccountId:  "",
					RoleName:   "",
					ExternalId: "",
				},
				"ap-northeast-1",
			)

			got, err := a.TerminateInstances(tt.args.instanceIds)
			if (err != nil) != tt.wantErr {
				t.Errorf("TerminateInstances() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("TerminateInstances() got = %v, want %v", got, tt.want)
			}
		})
	}
}
