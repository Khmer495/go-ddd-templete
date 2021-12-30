# Go

## lint

golangci-lint を使用します。

```sh
make lint
```

## テスト

```sh
make test
```

## APi サーバー

docker compose を使用します。ホットリロードに対応しています。  
事前にミドルウェアを起動してください。

- サーバーの起動：`make run-api`
- サーバーの停止：`make run-api-down`
- イメージの再ビルド（キャッシュを無視する）：`make run-api-build`
