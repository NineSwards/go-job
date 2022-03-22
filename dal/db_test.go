package dal

import (
	"testing"
)

func Test_Select(t *testing.T) {
	//使用工具获取数据库连接
	db := InitDB()
	//准备SQL语句
	sql := "select * from xxl_job_lock"
	//对SQL语句进行预处理
	stmt, err := db.Prepare(sql)
	if err != nil {
		panic(err)
	}
	rows, err := stmt.Query()
	if err != nil {
		//SQL执行失败，直接panic
		panic(err)
	}
	for rows.Next() {
		var lock_name string
		err := rows.Scan(&lock_name)
		if err != nil {
			//读取结果集失败
			panic(err)
		}
		t.Logf("lock_name=%s", lock_name)
	}
}
