package server

import (
	"log"

	"github.com/go-kratos/kratos/contrib/registry/etcd/v2"
	"github.com/go-kratos/kratos/v2/registry"
	"github.com/google/wire"
	clientv3 "go.etcd.io/etcd/client/v3"
	"google.golang.org/grpc"
	"vhrgo/app/frontend/web/internal/conf"
)

// ProviderSet is server providers.
var ProviderSet = wire.NewSet(NewHTTPServer, NewClient, NewRegistrar)

func NewClient(conf *conf.Management) *clientv3.Client {
	client, err := clientv3.New(clientv3.Config{
		Endpoints:            conf.Registry.Endpoints,
		AutoSyncInterval:     conf.Registry.AutoSyncInterval.AsDuration(),
		DialTimeout:          conf.Registry.DialTimeout.AsDuration(),
		DialKeepAliveTime:    conf.Registry.DialKeepAliveTime.AsDuration(),
		DialKeepAliveTimeout: conf.Registry.DialKeepAliveTimeout.AsDuration(),
		MaxCallSendMsgSize:   int(conf.Registry.MaxCallSendMsgSize),
		MaxCallRecvMsgSize:   int(conf.Registry.MaxCallRecvMsgSize),
		Username:             conf.Registry.Username,
		Password:             conf.Registry.Password,
		RejectOldCluster:     conf.Registry.RejectOldCluster,
		PermitWithoutStream:  conf.Registry.PermitWithoutStream,
		// TLS:                  &tls.Config{},
		DialOptions: []grpc.DialOption{
			grpc.WithBlock(),
		},
	})
	if err != nil {
		log.Panic(err)
	}
	return client
}

func NewRegistrar(client *clientv3.Client,
	conf *conf.Management) registry.Registrar {
	return etcd.New(client,
		etcd.Namespace(conf.Registry.Namespace),
		etcd.RegisterTTL(conf.Registry.RegisterTtl.AsDuration()),
		etcd.MaxRetry(int(conf.Registry.MaxRetry)))
}
