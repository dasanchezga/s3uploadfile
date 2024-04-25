package services

import (
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

func UploadToS3(filePath string, bucketName string, contentType string, key string) error {
	// Create a new AWS session
	// Configurar las credenciales de AWS desde las variables de entorno
    sess := session.Must(session.NewSession(&aws.Config{
        Region: aws.String("sa-east-1"), // región AWS deseada
        Credentials: credentials.NewEnvCredentials(),
    }))

	// Create a new cliente de servicio S3
	svc := s3.New(sess)

	// Open the file to upload
	file, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	// Upload the file to S3
	_, err = svc.PutObject(&s3.PutObjectInput{
		Bucket: aws.String(bucketName),
		Key:    aws.String(key),
		Body:   file,
		ContentType: aws.String(contentType),
	})
	if err != nil {
		return err
	}

	return nil
}
