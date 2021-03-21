# gorm_sample1

gormの動作確認用。

- gormを利用して、SQLiteへ接続。
- テーブル「users」「tasks」をMigrateで作成。（Has Manyの関係）
- レコードの削除確認（Soft DeleteとHard Delete）
- レコードの作成（Has Manyの関係のデータについて）
- レコードの読込（最初の一件、条件にマッチした最初の一件、全件）

## 各種バージョン

- OS: macOS Mojave 10.14.6
- Go: go version go1.16.2 darwin/amd64
- GORM: v1.9.16
- sqlite3ドライバ: v1.14.6

## 準備

- 事前に動作確認環境へSQLiteをインストールしておくこと

## 動作確認

1. `task_app`ディレクトリへ移動
2. `go run gormsample.go`を実行する
3. 出力結果を確認する

## 備忘

vscodeを利用しているならば、プラグインの「SQLITE EXPLORER」がインストールしてあると便利。
インストール後に、コマンドパレットで`sql`と実行すると、エクスプローラに表示される。
