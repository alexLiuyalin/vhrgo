package data

import (
	"context"
	"time"
	// apicore "vhrgo/api/core/service/v1"
	"vhrgo/app/frontend/web/internal/conf"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/registry"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	_ "github.com/go-sql-driver/mysql"
	"github.com/google/wire"
	"github.com/jmoiron/sqlx"
	redis "github.com/redis/go-redis/v9"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(
	NewData,
	NewDB,
	NewRedis,
	NewUserRepo,
	NewMenuRepo,
	NewEmployeeRepo,
	NewCaptchaRepo,
	NewAuthRepo,
)

// Data .
type Data struct {
	log *log.Helper
	// user apicore.UserClient

	db        *sqlx.DB
	cache     redis.UniversalClient
	cacheTime uint32
}

// NewData 创建数据库对象
func NewDB(conf *conf.Data, logger log.Logger) *sqlx.DB {
	log := log.NewHelper(log.With(logger, "module", "data/newDB"))

	// 使用 sql.DB 对象
	db, err := sqlx.Connect(conf.Database.Driver, conf.Database.Source)
	if err != nil {
		log.Fatalf("failed opening connection to db: %v", err)
	}
	db.SetMaxIdleConns(int(conf.Database.MaxIdleConns))               // 设置空闲状态下的最大连接数
	db.SetMaxOpenConns(int(conf.Database.MaxOpenConns))               // 设置最大打开连接数
	db.SetConnMaxLifetime(conf.Database.ConnMaxLifetime.AsDuration()) // 设置连接的最大可复用时间
	return db
}

// NewRedis 创建缓存对象
func NewRedis(cs *conf.Server, cd *conf.Data,
	logger log.Logger) redis.UniversalClient {
	log := log.NewHelper(log.With(logger, "module", "data/redis"))
	client := redis.NewClient(&redis.Options{
		Addr:         cd.Redis.Addr,
		ClientName:   cs.Name,
		Password:     cd.Redis.Password,
		ReadTimeout:  cd.Redis.ReadTimeout.AsDuration(),
		WriteTimeout: cd.Redis.WriteTimeout.AsDuration(),
		DialTimeout:  cd.Redis.DialTimeout.AsDuration(),
		PoolSize:     int(cd.Redis.PoolSize),
		DB:           int(cd.Redis.Db),
	})

	timeout, cancelFunc := context.WithTimeout(context.Background(), time.Second*2)
	defer cancelFunc()
	err := client.Ping(timeout).Err()
	if err != nil {
		log.Fatalf("redis connect error: %v", err)
	}
	return client
}

// NewData .
func NewData(mysqlx *sqlx.DB,
	redisCmd redis.UniversalClient,
	r registry.Registrar,
	cs *conf.Server,
	mic *conf.Microservices,
	conf *conf.Management,
	logger log.Logger) (*Data, func(), error) {
	log := log.NewHelper(log.With(logger, "merchant", "data/data"))

	coreConn, err := grpc.DialInsecure(
		context.Background(),
		grpc.WithEndpoint("discovery:///"+mic.Core),
		grpc.WithDiscovery(r.(registry.Discovery)), // 服务发现
		grpc.WithTimeout(conf.Registry.DialTimeout.AsDuration()),
		grpc.WithMiddleware(
			// tracing.Client(tracing.WithTracerProvider(tp)),// 增加链路追踪
			recovery.Recovery(),
		),
	)
	if err != nil {
		log.Errorf("grpc.DialInsecure error(%v)", err)
		return nil, nil, err
	}

	d := &Data{
		db:    mysqlx,
		cache: redisCmd,
		// user:  apicore.NewUserClient(coreConn),
	}
	return d, func() {
		if err := d.db.Close(); err != nil {
			log.Error(err)
		}

		if err := d.cache.Close(); err != nil {
			log.Error(err)
		}
		coreConn.Close()
	}, nil
}
