package main

import (
	v1 "github.com/taheralfayad/portfolio_v2/api/v1"

	"os"
	"context"
	"log"
	"errors"

	"fmt"
	"database/sql"
	
	"github.com/aws/smithy-go"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/cors"
	_ "github.com/lib/pq"
)

func main() {
	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept"},
		AllowCredentials: true,
	}))

	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")
	access_key_id := os.Getenv("RUSTFS_ACCESS_KEY_ID")
	secret_access_key := os.Getenv("RUSTFS_SECRET_ACCESS_KEY")
	endpoint := os.Getenv("RUSTFS_ENDPOINT_URL")

	connStr := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=disable",
		user,
		password,
		host,
		port,
		dbName,
	)

	cfg := aws.Config{
			EndpointResolver: aws.EndpointResolverFunc(func(service, region string) (aws.Endpoint, error) {
					return aws.Endpoint{
							URL: endpoint,
					}, nil
			}),
			Credentials: aws.NewCredentialsCache(credentials.NewStaticCredentialsProvider(access_key_id, secret_access_key, "")),
	}


	ctx := context.Background()

	client := s3.NewFromConfig(cfg, func(o *s3.Options) {
			o.UsePathStyle = true
	})

	_, err := client.CreateBucket(ctx, &s3.CreateBucketInput{
			Bucket: aws.String("portfolio-assets"),
	})

	if err != nil {
			var apiErr smithy.APIError
			if errors.As(err, &apiErr) {
					if apiErr.ErrorCode() != "BucketAlreadyOwnedByYou" {
						log.Fatalf("create bucket failed: %v", err)
					}
			}
	}

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	if err := db.Ping(); err != nil {
		panic(err)
	}

	r.POST("/work-experiences", func(c *gin.Context) {
		v1.AddWorkExperience(c, db)
	})

	r.POST("/projects", func(c *gin.Context) {
		v1.AddProject(c, db, ctx, client)
	})

	r.POST("/users", func(c *gin.Context) {
		v1.AddUser(c, db)
	})

	r.POST("/skills", func(c *gin.Context) {
		v1.AddSkill(c, db)
	})

	r.GET("/all-tables", func(c *gin.Context) {
		v1.GetAllTables(c, db)
	})

	r.GET("/users", func(c *gin.Context) {
		v1.GetUsers(c, db)
	})

	r.GET("/work-experiences", func(c *gin.Context) {
		v1.GetWorkExperiences(c, db)
	})

	r.PUT("/work-experiences", func(c *gin.Context) {
		v1.EditWorkExperience(c, db)
	})

	r.PUT("/users", func(c *gin.Context) {
		v1.EditUser(c, db)
	})

	r.Run()
}
