package database

import (
	"github.com/couchbase/gocb/v2"
	"log"
	"time"
)

const (
	connectionString = "couchbase"
	bucketName       = "passgw"
	username         = "admin"
	password         = "password"
)

func InitCouchbaseBucket() *gocb.Bucket {
	cluster, err := gocb.Connect("couchbase://"+connectionString, gocb.ClusterOptions{
		Authenticator: gocb.PasswordAuthenticator{
			Username: username,
			Password: password,
		},
	})
	if err != nil {
		log.Fatalln(err.Error())
	}

	bucket := cluster.Bucket(bucketName)
	err = bucket.WaitUntilReady(5*time.Second, nil)
	if err != nil {
		log.Fatal(err.Error())
	}
	return bucket
}
