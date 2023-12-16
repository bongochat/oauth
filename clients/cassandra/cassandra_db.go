package cassandra

import (
	"github.com/bongochat/bongochat-oauth/config"
	"github.com/gocql/gocql"
)

var (
	session *gocql.Session
	conf    = config.GetConfig()
)

func init() {
	cluster := gocql.NewCluster(conf.Host)
	cluster.Keyspace = conf.Keyspace
	cluster.Consistency = gocql.Quorum

	var err error
	if session, err = cluster.CreateSession(); err != nil {
		panic(err)
	}
}

func GetSession() *gocql.Session {
	return session
}
