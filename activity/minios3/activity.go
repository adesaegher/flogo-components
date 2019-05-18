// Package amazons3 uploads or downloads files from minio Simple Storage Service (S3)
package minios3

import (
	"context"
	"github.com/TIBCOSoftware/flogo-lib/core/activity"
	"github.com/TIBCOSoftware/flogo-lib/logger"
	"github.com/minio/minio-go"
)

const (
	ivAction             = "action"
	ivS3Endpoint         = "s3Endpoint"
	ivAwsAccessKeyID     = "awsAccessKeyID"
	ivAwsSecretAccessKey = "awsSecretAccessKey"
	ivUseSsl             = "useSsl"
	ivAwsRegion          = "awsRegion"
	ivS3BucketName       = "s3BucketName"
	ivLocalLocation      = "localLocation"
	ivS3Location         = "s3Location"
	ivS3NewLocation      = "s3NewLocation"
	ovResult             = "result"
)

// log is the default package logger
var log = logger.GetLogger("activity-minios3")

// MyActivity is a stub for your Activity implementation
type MyActivity struct {
	metadata *activity.Metadata
}

// NewActivity creates a new activity
func NewActivity(metadata *activity.Metadata) activity.Activity {
	return &MyActivity{metadata: metadata}
}

// Metadata implements activity.Activity.Metadata
func (a *MyActivity) Metadata() *activity.Metadata {
	return a.metadata
}

// Eval implements activity.Activity.Eval
func (a *MyActivity) Eval(context activity.Context) (done bool, err error) {

	// Get the action
	action := context.GetInput(ivAction).(string)
	s3Endpoint := context.GetInput(ivS3Endpoint).(string)
	awsRegion := context.GetInput(ivAwsRegion).(string)
	useSsl := context.GetInput(ivUseSsl).(bool)
	s3BucketName := context.GetInput(ivS3BucketName).(string)
	// localLocation is a file when uploading a file or a directory when downloading a file
	localLocation := context.GetInput(ivLocalLocation).(string)
	s3Location := context.GetInput(ivS3Location).(string)
	s3NewLocation := context.GetInput(ivS3NewLocation).(string)

	// AWS Credentials, only if needed
	var awsAccessKeyID, awsSecretAccessKey = "", ""
	if context.GetInput(ivAwsAccessKeyID) != nil {
		awsAccessKeyID = context.GetInput(ivAwsAccessKeyID).(string)
	}
	if context.GetInput(ivAwsSecretAccessKey) != nil {
		awsSecretAccessKey = context.GetInput(ivAwsSecretAccessKey).(string)
	}

	// See which action needs to be taken
	var s3err error
	switch action {
	case "download":
		s3err = downloadFileFromS3(s3Endpoint, awsAccessKeyID, awsSecretAccessKey, useSsl, awsRegion, localLocation, s3Location, s3BucketName)
	case "upload":
		s3err = uploadFileToS3(s3Endpoint, awsAccessKeyID, awsSecretAccessKey, useSsl, awsRegion, localLocation, s3Location, s3BucketName)
	case "delete":
		s3err = deleteFileFromS3(s3Endpoint, awsAccessKeyID, awsSecretAccessKey, useSsl, awsRegion, s3Location, s3BucketName)
	case "copy":
		s3err = copyFileOnS3(s3Endpoint, awsAccessKeyID, awsSecretAccessKey, useSsl, awsRegion, s3Location, s3BucketName, s3NewLocation)
	}
	if s3err != nil {
		// Set the output value in the context
		context.SetOutput(ovResult, s3err.Error())
		return true, s3err
	}

	// Set the output value in the context
	context.SetOutput(ovResult, "OK")

	return true, nil
}

// Function to download a file from an S3 bucket
func downloadFileFromS3(s3Endpoint string, awsAccessKeyID string, awsSecretAccessKey string, useSsl bool, awsRegion string, directory string, s3Location string, s3BucketName string) error {
	// Create an instance of the Minio Client
        minioClient, err := minio.New(s3Endpoint, awsAccessKeyID, awsSecretAccessKey, useSsl)
        if err != nil {
                return err
        }

	// Download the file to disk
        err = minioClient.FGetObject(s3BucketName, s3Location, directory, minio.GetObjectOptions{})
	if err != nil {
		return err
	}

	return nil
}

// Function to delete a file from an S3 bucket
func deleteFileFromS3(s3Endpoint string, awsAccessKeyID string, awsSecretAccessKey string, useSsl bool, awsRegion string, s3Location string, s3BucketName string) error {
	// Create an instance of the Minio Client
        minioClient, err := minio.New(s3Endpoint, awsAccessKeyID, awsSecretAccessKey, useSsl)
        if err != nil {
                return err
        }

	// Delete the file from S3
	err = minioClient.RemoveObject(s3BucketName, s3Location)
	if err != nil {
		return err
	}

	return nil
}

// Function to upload a file from an S3 bucket
func uploadFileToS3(s3Endpoint string, awsAccessKeyID string, awsSecretAccessKey string, useSsl bool, awsRegion string, localFile string, s3Location string, s3BucketName string) error {
	// Create an instance of the Minio Client
        minioClient, err := minio.New(s3Endpoint, awsAccessKeyID, awsSecretAccessKey, useSsl)
        if err != nil {
                return err
        }

	// Upload the file
	_, err = minioClient.FPutObject(s3BucketName, s3Location, localFile, minio.PutObjectOptions{});
	if err != nil {
		return err
	}

	return nil
}

// Function to copy a file in an S3 bucket
func copyFileOnS3(s3Endpoint string, awsAccessKeyID string, awsSecretAccessKey string, useSsl bool, awsRegion string, s3Location string, s3BucketName string, s3NewLocation string) error {
	// Create an instance of the Minio Client
        minioClient, err := minio.New(s3Endpoint, awsAccessKeyID, awsSecretAccessKey, useSsl)
        if err != nil {
                return err
        }

	// Prepare the copy object
        src := minio.NewSourceInfo(s3BucketName, s3BucketName, nil)
        dst, err := minio.NewDestinationInfo(s3BucketName, s3NewLocation, nil, nil)

	// Copy the object
	err = minioClient.CopyObject(dst, src)
	if err != nil {
		return err
	}

	return nil
}
