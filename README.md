Google Cloud Pub/Sub over HTTP Gateway server.

[![Docker Automated build](https://img.shields.io/docker/automated/pyyoshi/pubsub-gateway-server.svg)](https://hub.docker.com/r/pyyoshi/pubsub-gateway-server/) [![Docker Build Status](https://img.shields.io/docker/build/pyyoshi/pubsub-gateway-server.svg)](https://hub.docker.com/r/pyyoshi/pubsub-gateway-server/)

# gateway-server

## 環境変数

### (Optional) GATEWAY_SERVER_DEBUG

``cmd/gateway``のデバッグ機能を有効にするかどうか

- ``true``
- ``false`` <= デフォルト値

### (Optional) GATEWAY_SERVER_BIND_ADDRESS

``cmd/gateway``のHTTPサーバのBindAddress

- ``0.0.0.0:8089`` <= デフォルト値

### (Required) GATEWAY_SERVER_GOOGLE_SERVICE_ACCOUNT_BASE64

Google Service AccountファイルをBase64でエンコードしたもの
Google Cloud Pub/Subを利用可能なサービスアカウントを設定してください

### (Required) GATEWAY_SERVER_GOOGLE_PROJECT_ID

Google Cloud Pub/Subを利用するためのGoogle Cloud PlatformのProject IDを設定してください
``GATEWAY_SERVER_GOOGLE_SERVICE_ACCOUNT_BASE64``で設定したサービスアカウントと同じProject IDを設定してください

### (Optional) PUBSUB_EMULATOR_HOST

``Cloud Pub/Sub emulator``を利用する場合は設定してください
docker-composeで実行している``pubsub_emulator``サービスの``Host``を指定してください
``$ gcloud beta emulators pubsub env-init`` と同じ値になります
この値を設定すると強制的にエミュレーターへ接続を試みます

## エンドポイント一覧

``gen/http/openapi.yaml``にAPIスキーマが保存されているので使い方はそれを参照

## 注意事項

``gateway-server``はプライベートネットワーク内でのみ利用してください
パブリックネットワークを考慮した作りにはなっていません
パブリックネットワークで利用する場合はプロキシなどで前段に認証機構を備えるようにしてください

# testing-subscriber

## 環境変数

### (Required) PUBSUB_EMULATOR_HOST

docker-composeで実行している``pubsub_emulator``サービスの``Host``を指定してください
``$ gcloud beta emulators pubsub env-init`` と同じ値になります

### (Required) PUBSUB_TOPIC

``gateway-server``でパブリッシュするときの``Topic``と同じものを指定してください
指定した``Topic``が未作成の場合は自動で作成されます

### (Required) PUBSUB_SUBSCRIPTION

``PUBSUB_TOPIC``と関連づいた``Subscription``を指定してください
指定した``Subscription``が未作成の場合は自動で作成されます

### (Required) GATEWAY_SERVER_GOOGLE_PROJECT_ID

Google Cloud Pub/Subを利用するためのGoogle Cloud PlatformのProject IDを設定してください
``GATEWAY_SERVER_GOOGLE_SERVICE_ACCOUNT_BASE64``で設定したサービスアカウントと同じProject IDを設定してください

# 開発

``realize``でファイル監視を行い、更新がある都度リビルドし再実行する仕組みになっているため、``docker-compose up``を行えば、そのままコードを変更し確認することができます

## 構成

### cmd/gateway

``http``プロトコルで``Cloud Pub/Sub``を使えるようにするためのゲートウェイ

``gen/http/openapi.yaml``にAPIスキーマが保存されているので使い方はそれを参照

### cmd/gateway-cli

``gateway-server``のAPIクライアントCLI

### cmd/testing_subscriber

``Cloud Pub/Subエミュレータ``に接続し、``gateway-server``でパブリッシュしたメッセージを表示させます
デバッグ用途のみで利用してください

## サーバ起動

```bash
$ docker-compose down -v && docker-compose build && docker-compose up
```

# TODOs

``TODOs.md``を参照
