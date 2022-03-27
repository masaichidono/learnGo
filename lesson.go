package main

import (
	"database/sql"
	"fmt"
)

func my() error {
	db, err := sql.Open("mysql", "")

	if err != nil {
		//抛出错误
		panic("connect to mysql error")
	}

	query := "" //查询语句

	info, err := db.Exec(query) //查询数据
	if err == sql.ErrNoRows {
		//wrap这个error到上层，因为需要记录一些错误信息，方便定位
		return fmt.Errorf("data not found, sql: %s", query)
	}
	//处理业务逻辑
	deal(info)
	return nil
}
