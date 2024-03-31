package header

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/gocql/gocql"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

var attempCount = promauto.NewCounterVec(prometheus.CounterOpts{
	Name: "gocql_count_3",
	Help: "The total number of query attempts",
}, []string{
	"statement", "type",
})

var attempErrorCount = promauto.NewCounterVec(prometheus.CounterOpts{
	Name: "gocql_error_count_3",
	Help: "The total number of query attempts that failed",
}, []string{
	"statement", "type",
})

type queryObserver struct{}

func (qo *queryObserver) ObserveQuery(ctx context.Context, q gocql.ObservedQuery) {
	typ := "select"
	if strings.HasPrefix(q.Statement, "INSERT") || strings.HasPrefix(q.Statement, "insert") {
		typ = "insert"
	} else if strings.HasPrefix(q.Statement, "DELETE") || strings.HasPrefix(q.Statement, "delete") {
		typ = "delete"
	}

	attempCount.WithLabelValues(q.Statement, typ).Inc()
	if q.Err != nil {
		attempErrorCount.WithLabelValues(q.Statement, typ).Inc()
	}
}

func ConnectDB(seeds []string, keyspace string) *gocql.Session {
	var session *gocql.Session
	cluster := gocql.NewCluster(seeds...)
	cluster.QueryObserver = &queryObserver{}
	cluster.Timeout = 30 * time.Second
	cluster.ConnectTimeout = 30 * time.Second
	cluster.Keyspace = keyspace
	var err error
	for i := 0; i < 10; i++ {
		if session, err = cluster.CreateSession(); err == nil {
			return session
		}
		fmt.Println("cassandra", err, ". Retring after 5sec...")
		time.Sleep(15 * time.Second)
	}
	panic(err)
}
