package main

import (
	"flag"

	mymodel "taskapp/model" //modelパッケージ名を省略
)

func main() {

	// DB作成用可否のコマンドライン引数設定
	dbinit := flag.Bool("dbinit", false, "DBを初期化します")
	flag.Parse()

	mydb := &mymodel.MyModel{}
	mydb.Connect_database()
	// コマンドライン引数に「dbinit」が指定されている場合のみ、
	// DBのmigratte & 確認用の初期データの投入を行う
	if *dbinit == true {
		mydb.Init()
		mydb.InsertInitData()
		mydb.InitDataPrint() //デバッグプリント
	}

	// start Echo server
	e := newRouter()
	// 実行
	e.Logger.Fatal(e.Start(":1323"))

}
