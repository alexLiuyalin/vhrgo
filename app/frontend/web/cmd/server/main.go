package main

import (
	"flag"
	"fmt"
	"os"

	"vhrgo/app/frontend/web/internal/conf"

	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/config"
	"github.com/go-kratos/kratos/v2/config/file"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/tracing"
	"github.com/go-kratos/kratos/v2/transport/http"
)

const Version = "1.0.0" // 版本号
// go build -ldflags "-X main.Version=x.y.z"
var (
	// Name is the name of the compiled software.
	Name string
	// Version is the version of the compiled software.
	// flagconf is the config flag.
	flagconf string
	// 通过编译时传入的参数设置
	BuildTime string // 编译时间
	GitBranch string // git branch
	GitHash   string // git hash
	GoVersion string // go version
	// 命令行参数
	flagVersion bool // 不执行程序 只打印版本信息

	id, _ = os.Hostname()
)

func init() {
	// flag.StringVar(&flagconf, "conf", "../configs", "config path, eg: -conf config.yaml")
	flag.StringVar(&flagconf, "conf", "../../configs", "config path, eg: -conf config.yaml")
	flag.BoolVar(&flagVersion, "version", false, "show version, eg: -version")
}

func newApp(logger log.Logger,
	cs *conf.Server,
	hs *http.Server,
) *kratos.App {
	return kratos.New(
		kratos.ID(id),
		kratos.Name(Name),
		kratos.Version(Version),
		kratos.Metadata(map[string]string{"BuildTime": BuildTime,
			"GitBranch": GitBranch,
			"GitHash":   GitHash,
			"GoVersion": GoVersion}),
		kratos.Logger(logger),
		kratos.Server(
			hs,
		),
	)
}

func main() {
	flag.Parse()

	// 输出版本信息后退出
	if flagVersion {
		fmt.Printf(`BuildTime: %s
		GitBranch: %s
		GitHash: %s
		GoVersion: %s`, BuildTime, GitBranch, GitHash, GoVersion)
		os.Exit(0)
	}

	logger := log.With(log.NewStdLogger(os.Stdout),
		"ts", log.DefaultTimestamp,
		"caller", log.DefaultCaller,
		"service.id", id,
		"service.name", Name,
		"service.version", Version,
		"trace.id", tracing.TraceID(),
		"span.id", tracing.SpanID(),
	)
	c := config.New(
		config.WithSource(
			file.NewSource(flagconf),
		),
	)
	defer c.Close()

	if err := c.Load(); err != nil {
		panic(err)
	}

	var bc conf.Bootstrap
	if err := c.Scan(&bc); err != nil {
		panic(err)
	}
	app, cleanup, err := wireApp(bc.Server, bc.Data, bc.Microservices, bc.Management, logger)
	if err != nil {
		panic(err)
	}
	defer cleanup()

	// start and wait for stop signal
	if err := app.Run(); err != nil {
		panic(err)
	}
}
