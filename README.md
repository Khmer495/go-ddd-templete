# go-ddd-templete

このレポジトリはドメイン駆動開発における go のパッケージ構成のサンプルです。
簡易的な RESTful API サーバーを題材としています。

## 動作環境

- Go v1.17.5
- MySQL v8.0.23
- Docker
- docker-compose

## インフラの開発環境の起動

1. `.env.sample`を複製し、`.env`にリネームする。必要があれば環境変数を設定します。
2. `make run-infra-local`を実行します。

## ディレクトリ構成

```text
.
├── Makefile
├── app
│   └── ...
├── docs
│   └── api
│        └── openapi.yaml
├── local.infra.docker-compose.yaml
└── mysql
     └── ...
```

- app

  go の root ディレクトリです。詳しくは[app/README.md](app/README.md)に記載しています。

- docs/api/openapi.yaml

  題材とする RESTful API の定義ファイルです。
