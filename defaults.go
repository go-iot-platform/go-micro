package micro

import (
	"github.com/panovateam/go-micro/client"
	"github.com/panovateam/go-micro/debug/trace"
	"github.com/panovateam/go-micro/server"
	"github.com/panovateam/go-micro/store"

	// set defaults
	gcli "github.com/panovateam/go-micro/client/grpc"
	memTrace "github.com/panovateam/go-micro/debug/trace/memory"
	gsrv "github.com/panovateam/go-micro/server/grpc"
	memoryStore "github.com/panovateam/go-micro/store/memory"
)

func init() {
	// default client
	client.DefaultClient = gcli.NewClient()
	// default server
	server.DefaultServer = gsrv.NewServer()
	// default store
	store.DefaultStore = memoryStore.NewStore()
	// set default trace
	trace.DefaultTracer = memTrace.NewTracer()
}
