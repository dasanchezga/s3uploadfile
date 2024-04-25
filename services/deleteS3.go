package services

import (
	"errors"
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

// ErrFileNotFound es un error que se utiliza cuando el archivo no se encuentra en S3.
var ErrFileNotFound = errors.New("archivo no encontrado en S3")

func DeleteFromS3(bucketName, key string) error {

	// Create a new AWS session
    sess := session.Must(session.NewSession(&aws.Config{
        Region: aws.String("sa-east-1"), // región AWS deseada
        Credentials: credentials.NewEnvCredentials(),
    }))

    // Create an S3 service client
    svc := s3.New(sess)

	// Verificar si el archivo existe en S3
    exists, err := FileExistsInS3(bucketName, key)
    if err != nil {
        return err
    }
    // Si el archivo no existe, informar y salir sin hacer nada más
    if !exists {
        fmt.Printf("El archivo %s no existe en el bucket %s\n", key, bucketName)
        return ErrFileNotFound
    }

    // Create input for DeleteObject operation
    input := &s3.DeleteObjectInput{
        Bucket: aws.String(bucketName),
        Key:    aws.String(key),
    }

    // Delete the object from S3
    _, err = svc.DeleteObject(input)
    if err != nil {
        return err
    }

    return nil
}
