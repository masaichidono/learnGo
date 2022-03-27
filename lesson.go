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
		//包装一下，记录下查询的参数,然后返回
		return fmt.Errorf("data not found, sql: %s", query)
	}
	//处理业务逻辑
	deal(info)
	return nil
}
