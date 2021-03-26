package handler

import (
	"os"
	"testing"
)

func TestGetServerApiKey(t *testing.T) {
	const apikey = "boobar"
	os.Setenv(SERVER_API_KEY_ENV, "badapikey")
	apikey1 := GetServerApiKey()
	if apikey1 == apikey {
		t.Fatalf("設定されるAPIキーが誤っている：%v", apikey1)
	}

	os.Setenv(SERVER_API_KEY_ENV, apikey)
	apikey2 := GetServerApiKey()
	if Server_api_key != apikey || apikey2 != apikey {
		t.Fatalf("サーバのAPIキーには[%v]が設定されているが、変数Server_api_key[%v]か戻り値[%v]のどちらかが誤っている。", apikey, Server_api_key, apikey2)
	}

}

func TestGetClientApiKey(t *testing.T) {
	const apikey = "boobar"
	os.Setenv(CLIENT_API_KEY_ENV, "badapikey")
	apikey1 := GetClientApiKey()
	if apikey1 == apikey {
		t.Fatalf("設定されるAPIキーが誤っている：%v", apikey1)
	}

	os.Setenv(CLIENT_API_KEY_ENV, apikey)
	apikey2 := GetClientApiKey()
	if Client_api_key != apikey || apikey2 != apikey {
		t.Fatalf("クライアントのAPIキーには[%v]が設定されているが、変数Client_api_key[%v]か戻り値[%v]のどちらかが誤っている。", apikey, Client_api_key, apikey2)
	}

}

func TestIsSettingServerApikey(t *testing.T) {
	os.Setenv(SERVER_API_KEY_ENV, "")
	GetServerApiKey()

	key1 := IsSettingServerApikey()
	if key1 != false {
		t.Fatalf("サーバ側のAPIキー[%v]の設定に誤りがある。本来[%v] 取得値[%v]",
			SERVER_API_KEY_ENV, true, key1)
	}

	os.Setenv(SERVER_API_KEY_ENV, "tesing_api_key")
	GetServerApiKey()

	key2 := IsSettingServerApikey()
	if key2 != true {
		t.Fatalf("サーバ側のAPIキー[%v]が正しく取得できていない。本来[%v] 取得値[%v]",
			SERVER_API_KEY_ENV, true, key2)
	}
}

func TestAuthApiKey(t *testing.T) {
	const (
		apikey    string = "boobar"
		batapikey string = "batapikey"
	)
	os.Setenv(SERVER_API_KEY_ENV, apikey)
	server_api_key := GetServerApiKey()

	auth1 := AuthApiKey(apikey)
	if auth1 != true {
		t.Fatalf("APIキー[%v[%v]]の認証が正しくできていない。本来[%v] 取得値[%v]",
			SERVER_API_KEY_ENV, server_api_key, true, auth1)
	}

	auth2 := AuthApiKey(batapikey)
	if auth2 != false {
		t.Fatalf("誤ったAPIキー「%v」を渡したのに認証しているのに、正しいキー「%v」との差分なく認証している。本来[%v] 取得値[%v]",
			batapikey, server_api_key, false, auth2)
	}

	os.Setenv(CLIENT_API_KEY_ENV, apikey)
	auth3 := AuthApiKey(Client_api_key)
	if auth3 != true {
		t.Fatalf("正しいクライアントのAPIキー「%v」を渡したのに、正しいサーバ側のキー「%v」との差分なく認証している。本来[%v] 取得値[%v]",
			Client_api_key, Server_api_key, true, auth3)
	}

}
