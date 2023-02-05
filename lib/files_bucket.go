package lib

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/gridfs"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type FilesBucket struct {
	*gridfs.Bucket
}

func NewFilesBucket(env Env) FilesBucket {

	opts := options.Client()
	opts.ApplyURI(env.BUCKET_URL)
	opts.SetMaxPoolSize(5)

	client, err := mongo.Connect(context.Background(), opts)
	if err != nil {
		log.Fatal("Error connecting to mongo database: ", err.Error())
	}

	bucket, err := gridfs.NewBucket(client.Database("mobile_db"))
	if err != nil {
		log.Fatal("Error connecting to bucket: ", err.Error())
	}

	return FilesBucket{
		Bucket: bucket,
	}
}
