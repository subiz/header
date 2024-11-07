package header

import (
	"context"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/gocql/gocql"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

var _hostname string

func init() {
	_hostname, _ = os.Hostname()
	_hostname = strings.Split(_hostname, "-")[0]
}

var attempCount = promauto.NewCounterVec(prometheus.CounterOpts{
	Name: "gocql_count_3",
	Help: "The total number of query attempts",
}, []string{
	"statement", "type",
})

var opsBandwidth = promauto.NewCounterVec(prometheus.CounterOpts{
	Name: "gocql_bandwidth",
	Help: "Gocql bandwidth",
}, []string{
	"host",
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

type frameObserver struct{}

func (qo *frameObserver) ObserveFrameHeader(ctx context.Context, h gocql.ObservedFrameHeader) {
	opsBandwidth.WithLabelValues(_hostname).Add(float64(h.Length))
}

func ConnectDB(seeds []string, keyspace string) *gocql.Session {
	var session *gocql.Session
	cluster := gocql.NewCluster(seeds...)
	cluster.QueryObserver = &queryObserver{}
	cluster.FrameHeaderObserver = &frameObserver{}
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
