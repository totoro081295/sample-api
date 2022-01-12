# Sample-api

JWT認証を使ったサンプルAPI

## 必須
- postgreSQL
- [direnv](https://github.com/direnv/direnv)

## Usage

```bash
git clone https://github.com/totoro081295/sample-api
cd sample-api
cp .envrc.example .envrc
## .envrcのDATABASE_URLのDB_USER_NAMEを自身のpostgresのuser名に変更する

## db作成
createdb sample-api
## シードデータ挿入
cd database/seeds
go build
./seeds

## 直下に戻る
cd ../../
## APIを立ち上げる
go build
./sample-api
```

```bash
## ログイン
curl -X POST -H "Content-Type: application/json" -d '{"email":"test@example.com", "password":"password"}' localhost:8081/oauth/login

## アカウント取得
curl -H "Authorization: Bearer {{ログインAPIで取得したaccessToken}}" localhost:8081/api/accounts
```

### ログインAPIのレスポンス
```json
{
    "status":{
        "result":"success",
        "message":"login is success."
    },
    "accessToken":"アクセストークン"
}
```

### アカウント取得APIのレスポンス
```json
{
    "id":"アカウントのID(UUID V4)",
    "email":"メールアドレス"
}
```