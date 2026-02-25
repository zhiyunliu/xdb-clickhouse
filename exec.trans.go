package xdbclickhouse

import (
	"context"

	"github.com/ClickHouse/clickhouse-go/v2/lib/driver"

	contribxdb "github.com/zhiyunliu/glue/contrib/xdb"
	"github.com/zhiyunliu/glue/xdb"
)

var _ xdb.ITrans = (*clickhouseTrans)(nil)

type clickhouseTrans struct {
	cfg   *contribxdb.Setting
	trans driver.Conn
}

func (db *clickhouseTrans) Query(ctx context.Context, sql string, input any, opts ...xdb.TemplateOption) (data xdb.Rows, err error) {
	err = NotImplemented
	return
}

func (db *clickhouseTrans) Multi(ctx context.Context, sql string, input any, opts ...xdb.TemplateOption) (data []xdb.Rows, err error) {
	err = NotImplemented
	return
}

func (db *clickhouseTrans) First(ctx context.Context, sql string, input any, opts ...xdb.TemplateOption) (data xdb.Row, err error) {
	err = NotImplemented
	return
}

func (db *clickhouseTrans) Scalar(ctx context.Context, sql string, input any, opts ...xdb.TemplateOption) (data interface{}, err error) {
	err = NotImplemented
	return
}

func (db *clickhouseTrans) Exec(ctx context.Context, sql string, input any, opts ...xdb.TemplateOption) (r xdb.Result, err error) {
	err = NotImplemented
	return
}

func (db *clickhouseTrans) QueryAs(ctx context.Context, sql string, input any, result any, opts ...xdb.TemplateOption) (err error) {
	return NotImplemented
}

func (db *clickhouseTrans) FirstAs(ctx context.Context, sql string, input any, result any, opts ...xdb.TemplateOption) (err error) {
	return NotImplemented
}

func (d *clickhouseTrans) Rollback() (err error) {
	return
}

func (d *clickhouseTrans) Commit() (err error) {
	return
}
