package houyi

import (
	_ "go.uber.org/automaxprocs"

	"github.com/aide-cloud/moon/cmd/server/houyi/internal/houyiconf"
	"github.com/aide-cloud/moon/cmd/server/houyi/internal/server"
	"github.com/aide-cloud/moon/pkg/conn"
	"github.com/aide-cloud/moon/pkg/env"
	sLog "github.com/aide-cloud/moon/pkg/log"
	"github.com/aide-cloud/moon/pkg/types"

	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/config"
	"github.com/go-kratos/kratos/v2/config/file"
	"github.com/go-kratos/kratos/v2/encoding/json"
	"github.com/go-kratos/kratos/v2/log"
	"google.golang.org/protobuf/encoding/protojson"
)

func init() {
	//增加这段代码
	json.MarshalOptions = protojson.MarshalOptions{
		UseEnumNumbers: true, // 将枚举值作为数字发出，默认为枚举值的字符串
	}
}

func newApp(c *houyiconf.Bootstrap, srv *server.Server, logger log.Logger) *kratos.App {
	opts := []kratos.Option{
		kratos.ID(env.ID()),
		kratos.Name(env.Name()),
		kratos.Version(env.Version()),
		kratos.Metadata(env.Metadata()),
		kratos.Logger(logger),
		kratos.Server(srv.GetServers()...),
	}
	registerConf := c.GetServer().GetRegistry()
	if !types.IsNil(registerConf) {
		register, err := conn.NewRegister(c.GetServer().GetRegistry())
		if !types.IsNil(err) {
			log.Warnw("register error", err)
			panic(err)
		}
		opts = append(opts, kratos.Registrar(register))
	}

	return kratos.New(opts...)
}

func Run(flagconf string) {
	c := config.New(config.WithSource(file.NewSource(flagconf)))
	defer c.Close()

	if err := c.Load(); !types.IsNil(err) {
		panic(err)
	}

	var bc houyiconf.Bootstrap
	if err := c.Scan(&bc); !types.IsNil(err) {
		panic(err)
	}

	env.SetName(bc.GetServer().GetName())
	env.SetMetadata(bc.GetServer().GetMetadata())
	env.SetEnv(bc.GetEnv())

	logger := sLog.GetLogger()
	log.SetLogger(logger)
	app, cleanup, err := wireApp(&bc, logger)
	if !types.IsNil(err) {
		panic(err)
	}
	defer cleanup()

	// start and wait for stop signal
	if err = app.Run(); !types.IsNil(err) {
		panic(err)
	}
}