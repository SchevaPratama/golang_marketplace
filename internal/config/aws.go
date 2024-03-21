package config

import (
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/spf13/viper"
	"log"
	"os"
)

func NewAws(v *viper.Viper) *s3.Client {

	fmt.Println("S3 id :: ", os.Getenv("S3_ID"))
	fmt.Println("S3 secret :: ", os.Getenv("S3_SECRET_KEY"))
	fmt.Println("S3 region :: ", os.Getenv("S3_REGION"))

	cfg, err := config.LoadDefaultConfig(context.TODO(),
		config.WithCredentialsProvider(credentials.NewStaticCredentialsProvider(
			os.Getenv("S3_ID"),         //v.GetString("aws.id"),
			os.Getenv("S3_SECRET_KEY"), // v.GetString("aws.secret"),
			"",
		)),
		config.WithRegion(os.Getenv("S3_REGION")),
	)
	if err != nil {
		log.Fatalf("Failed connect aws: %v", err)
		return nil
	}

	client := s3.NewFromConfig(cfg)
	return client
}
