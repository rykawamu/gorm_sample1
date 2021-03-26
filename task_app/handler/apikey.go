package handler

import "os"

const (
	SERVER_API_KEY_ENV string = "MON_API_KEY_SVR"
	CLIENT_API_KEY_ENV string = "MON_API_KEY_CLI"
)

var (
	Server_api_key string
	Client_api_key string
)

// SERVER側のAPIキーを環境変数から取得し、変数へ設定する
func GetServerApiKey() string {
	key := apiKey(SERVER_API_KEY_ENV)
	Server_api_key = key
	return key
}

// CLIENT側のAPIキーを取得し、変数へ設定する。
func GetClientApiKey() string {
	key := apiKey(CLIENT_API_KEY_ENV)
	Client_api_key = key
	return key
}

func apiKey(env_key string) (key string) {
	key = os.Getenv(env_key)
	return key
}

func IsSettingServerApikey() bool {
	if len(Server_api_key) == 0 {
		return false
	}
	return true
}

func AuthApiKey(clientKey string) bool {
	if len(clientKey) == 0 {
		return false
	}
	if clientKey != Server_api_key {
		return false
	}
	return true
}
