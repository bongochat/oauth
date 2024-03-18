package cassandra

import (
	"crypto/tls"
	"log"
	"os"

	"github.com/aws/aws-sigv4-auth-cassandra-gocql-driver-plugin/sigv4"
	"github.com/gocql/gocql"
	"github.com/joho/godotenv"
)

var (
	session *gocql.Session
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	host := os.Getenv("DB_HOST")
	keyspace := os.Getenv("KEYSPACE")
	access_key := os.Getenv("AWS_ACCESS_KEY_ID")
	secret_key := os.Getenv("AWS_SECRET_ACCESS_KEY")
	region := os.Getenv("AWS_REGION")

	cluster := gocql.NewCluster(host)

	if os.Getenv("ENV") == "PRODUCTION" {
		cluster.Port = 9142
		cluster.Keyspace = keyspace
		cluster.Consistency = gocql.LocalQuorum

		var auth sigv4.AwsAuthenticator = sigv4.NewAwsAuthenticator()
		auth.Region = region
		auth.AccessKeyId = access_key
		auth.SecretAccessKey = secret_key

		cluster.ProtoVersion = 4
		cluster.SslOpts = &gocql.SslOptions{Config: &tls.Config{ServerName: host, InsecureSkipVerify: true}}

		cluster.Authenticator = auth
		cluster.DisableInitialHostLookup = false

	} else {
		host := os.Getenv("DB_HOST")
		keyspace := os.Getenv("KEYSPACE")

		log.Println(host, keyspace)
		cluster.Keyspace = keyspace
		cluster.Consistency = gocql.Quorum
	}

	if session, err = cluster.CreateSession(); err != nil {
		panic(err)
	}
}

func GetSession() *gocql.Session {
	return session
}
