package xdbclickhouse

import (
	"errors"
	"fmt"
	"regexp"
	"time"

	clickhouse "github.com/ClickHouse/clickhouse-go/v2"
	"github.com/zhiyunliu/glue/config"
	contribxdb "github.com/zhiyunliu/glue/contrib/xdb"
	"github.com/zhiyunliu/glue/xdb"
)

func init() {
	xdb.Register(&clickhouseResolver{})
	//tpl.Register(New(Proto, ArgumentPrefix))
}

const (
	Proto          = "clickhouse"
	ArgumentPrefix = ""
)

type clickhouseResolver struct {
}

func (s *clickhouseResolver) Name() string {
	return Proto
}

func (s *clickhouseResolver) Resolve(connName string, setting config.Config, opts ...xdb.Option) (dbObj interface{}, err error) {
	cfg := contribxdb.NewConfig(connName)
	err = setting.ScanTo(cfg.Cfg)
	if err != nil {
		return nil, fmt.Errorf("can't read clickhouse config:%w", err)
	}
	// 解析clickhouse链接
	reg := regexp.MustCompile(`^clickhouse://(.*):(.*)@(.*)/(.*)\?(.*)`)
	params := reg.FindSubmatch([]byte(cfg.Cfg.Conn))
	options := clickhouse.Options{
		DialTimeout:     time.Duration(cfg.Cfg.LifeTime) * time.Second,
		MaxOpenConns:    cfg.Cfg.MaxOpen,
		MaxIdleConns:    cfg.Cfg.MaxIdle,
		ConnMaxLifetime: time.Duration(cfg.Cfg.LifeTime) * time.Second,
		Settings: clickhouse.Settings{
			"max_execution_time": 1000,
		},
		ConnOpenStrategy: clickhouse.ConnOpenInOrder,
		Compression: &clickhouse.Compression{
			Method: clickhouse.CompressionLZ4,
		},
	}
	if len(params) < 5 {
		return nil, errors.New("clickhouse config is incorrect")
	}
	for index, param := range params {
		paramsString := string(param)
		switch index {
		case 1:
			options.Auth.Username = paramsString
		case 2:
			options.Auth.Password = paramsString
		case 3:
			options.Addr = []string{paramsString}
		case 4:
			options.Auth.Database = paramsString
		}
	}
	conn, err := clickhouse.Open(&options)
	if err != nil {
		return nil, err
	}
	return &clickhousedb{
		cfg: cfg,
		db:  conn,
	}, nil

}
