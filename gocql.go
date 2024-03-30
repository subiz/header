package header

import (
	"context"
	"fmt"
	"time"

	"github.com/gocql/gocql"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

var attempCount = promauto.NewCounterVec(prometheus.CounterOpts{
	Name: "gocql_count_2",
	Help: "The total number of query attempts",
}, []string{
	"statement",
})

var attempErrorCount = promauto.NewCounterVec(prometheus.CounterOpts{
	Name: "gocql_error_count",
	Help: "The total number of query attempts that failed",
}, []string{
	"statement",
})

type queryObserver struct{}

func (qo *queryObserver) ObserveQuery(ctx context.Context, q gocql.ObservedQuery) {
	attempCount.WithLabelValues(q.Statement).Inc()
	if q.Err != nil {
		attempErrorCount.WithLabelValues(q.Statement).Inc()
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
