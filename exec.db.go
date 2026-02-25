package xdbclickhouse

import (
	"context"

	"github.com/ClickHouse/clickhouse-go/v2/lib/driver"

	"fmt"

	contribxdb "github.com/zhiyunliu/glue/contrib/xdb"
	"github.com/zhiyunliu/glue/xdb"
)

// error:
//   - 当成功实现 xdb.IDB 接口时返回 nil
//   - 当未实现接口时返回包含错误信息的 error 对象
func Test() error {
	// 创建临时 clickhousedb 实例并赋值给空接口变量
	var tmpdb any = &clickhousedb{}

	// 尝试将临时对象断言为 xdb.IDB 接口
	// 成功断言说明已实现接口，返回 nil
	if _, ok := tmpdb.(xdb.IDB); ok {
		return nil
	}

	// 断言失败说明未实现接口，返回错误信息
	return fmt.Errorf("clickhousedb.Test():xdb.IDB not implemented")
}

var _ xdb.IDB = (*clickhousedb)(nil)

type clickhousedb struct {
	cfg *contribxdb.Setting
	db  driver.Conn
}

func (db *clickhousedb) Query(ctx context.Context, sql string, input any, opts ...xdb.TemplateOption) (data xdb.Rows, err error) {
	err = NotImplemented
	return
}

func (db *clickhousedb) Multi(ctx context.Context, sql string, input any, opts ...xdb.TemplateOption) (data []xdb.Rows, err error) {
	err = NotImplemented
	return
}

func (db *clickhousedb) First(ctx context.Context, sql string, input any, opts ...xdb.TemplateOption) (data xdb.Row, err error) {
	err = NotImplemented
	return
}

func (db *clickhousedb) Scalar(ctx context.Context, sql string, input any, opts ...xdb.TemplateOption) (data interface{}, err error) {
	err = NotImplemented
	return
}

func (db *clickhousedb) Exec(ctx context.Context, sql string, input any, opts ...xdb.TemplateOption) (r xdb.Result, err error) {
	err = NotImplemented
	return
}

func (db *clickhousedb) QueryAs(ctx context.Context, sql string, input any, result any, opts ...xdb.TemplateOption) (err error) {
	return NotImplemented
}

func (db *clickhousedb) FirstAs(ctx context.Context, sql string, input any, result any, opts ...xdb.TemplateOption) (err error) {
	return NotImplemented
}

func (db *clickhousedb) Begin() (trans xdb.ITrans, err error) {
	return &clickhouseTrans{
		cfg:   db.cfg,
		trans: db.db,
	}, nil

}

func (db *clickhousedb) BeginTx(context.Context) (trans xdb.ITrans, err error) {
	err = NotImplemented
	return
}

func (db *clickhousedb) Close() (err error) {
	if db.db != nil {
		return db.db.Close()
	}
	return
}

func (db *clickhousedb) GetImpl() (impl any) {
	return db.db
}

func (db *clickhousedb) Transaction(ctx context.Context, callback xdb.TransactionCallback) (err error) {
	return NotImplemented
}
