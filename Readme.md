
# GoとReactの勉強用レポジトリ

## 環境
- Go => go version go1.20.5 darwin/arm64
- React => 18.2.15
  - tailwindcss使う(https://tailwindcss.com/docs/guides/create-react-app)

- **※デプロイは対応していない**
## 学んだこと

- Go 外部パッケージを読み込んだ時は必ず下記を実行する
```bash
go mod tidy
```

- migrateする際は下記を実行(格納先や環境によって変わる)
```Go
// envを読み込むより先に処理が走ってしまうので、環境変数を指定して送ること
GO_ENV=dev go run migrate/migrate.go
```

- Goの起動
```
GO_ENV=dev go run main.go
```

## React 何かあれば記載

- prettireの設定
  - シングルコートがtrue
  - セミコロンがfalse


### 下記packageをインストール
```
npm i @tanstack/react-query@4.28.0
npm i @tanstack/react-query-devtools@4.28.0
npm i zustand@4.3.6
npm i @heroicons/react@2.0.16
npm i react-router-dom@6.10.0 axios@1.3.4
```

## 下記を参照
https://www.udemy.com/course/echo-go-react-restapi/