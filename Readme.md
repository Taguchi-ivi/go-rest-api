
# Goの勉強用レポジトリ

## 環境
go => go version go1.20.5 darwin/arm64


## 学んだこと

- 外部パッケージを読み込んだ時は必ず下記を実行する
```bash
go mod tidy
```

- migrateする際は下記を実行(格納先や環境によって変わる)
```Go
// envを読み込むより先に処理が走ってしまうので、環境変数を指定して送ること
GO_ENV=dev go run migrate/migrate.go
```

- 起動
```
GO_ENV=dev go run main.go
```

## 下記を参照
https://www.udemy.com/course/echo-go-react-restapi/