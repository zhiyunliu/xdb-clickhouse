package xdbclickhouse_test

import (
	"testing"

	xdbclickhouse "github.com/zhiyunliu/xdb-clickhouse"
)

func TestTest(t *testing.T) {

	gotErr := xdbclickhouse.Test()
	if gotErr != nil {
		t.Errorf("Test() failed: %v", gotErr)
	}
}
