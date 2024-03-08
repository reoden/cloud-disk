package test

import (
	"bytes"
	"cloud-disk/core/models"
	"encoding/json"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"testing"
	"xorm.io/xorm"
)

func TestXormTest(t *testing.T) {
	engine, err := xorm.NewEngine("mysql", "root:123456@tcp(localhost:3306)/cloud-disk?charset=utf8mb4&parseTime=True&loc=Local")

	if err != nil {
		t.Fatal(err)
	}

	data := make([]*models.UserBasic, 0)
	err = engine.Find(&data)
	if err != nil {
		t.Fatal(err)
	}

	marshal, err := json.Marshal(data)
	if err != nil {
		t.Fatal(err)
	}
	dst := new(bytes.Buffer)
	err = json.Indent(dst, marshal, "", "  ")
	if err != nil {
		t.Fatal(err)
	}

	fmt.Println(dst.String())
}
