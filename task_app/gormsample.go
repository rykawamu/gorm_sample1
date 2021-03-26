package main

import (
	"flag"
	"fmt"

	myhandler "taskapp/handler"
	mymodel "taskapp/model" //modelパッケージ名を省略
)

func main() {

	// API用の環境変数が設定されているかの確認
	myhandler.GetServerApiKey()
	myhandler.GetClientApiKey()
	if b := myhandler.IsSettingServerApikey(); b != true {
		fmt.Printf("環境変数「%v」にAPIキーを設定してください。\n",
			myhandler.SERVER_API_KEY_ENV)
		return
	}
	if auth_apikey := myhandler.AuthApiKey(myhandler.Client_api_key); auth_apikey != true {
		fmt.Printf("サーバとクライアントのAPIキーの設定値が異なります。「%v」「%v」。\n",
			myhandler.Server_api_key, myhandler.Client_api_key)
		return
	}

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
