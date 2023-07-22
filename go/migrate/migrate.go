// migrateを実行する際は、プログラムのエントリーポイントに配置したい。なのでpackage mainにしてmainに所属させる
// migrateさせる際は
package main

import (
	"fmt"
	"go-rest-api/db"
	"go-rest-api/model"
)

func main() {
	dbConn := db.NewDB()
	defer fmt.Println("Successfully Migrated")
	defer db.CloseDB(dbConn)
	// フィールドの値をゼロ値にして、テーブルを作成する
	dbConn.AutoMigrate(&model.User{}, &model.Task{}, &model.Tweet{})
}
