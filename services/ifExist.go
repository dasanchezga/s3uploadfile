package services

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

func FileExistsInS3(bucketName, key string) (bool, error) {
    // Create a new AWS session
    sess := session.Must(session.NewSession(&aws.Config{
        Region: aws.String("sa-east-1"), // región AWS deseada
        Credentials: credentials.NewEnvCredentials(),
    }))

    // Create an S3 service client
    svc := s3.New(sess)

    // Crea el input para la operación HeadObject
    input := &s3.HeadObjectInput{
        Bucket: aws.String(bucketName),
        Key:    aws.String(key),
    }

    // Realiza la operación HeadObject para verificar la existencia del archivo
    _, err := svc.HeadObject(input)
    if err != nil {
        if awsErr, ok := err.(awserr.Error); ok && awsErr.Code() == "NotFound" {
            // El objeto no existe en S3
            return false, nil
        }
        // Ocurrió un error diferente, devolverlo
        return false, err
    }

    // El objeto existe en S3
    return true, nil
}