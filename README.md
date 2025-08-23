# emopathy-api

Emopathy.app の API サーバー

## 概要

Emopathy.app のバックエンド API Server であり、Golang/Echo で構築、PostgreSQL を Database として使用

## 技術スタック

- **言語**: Go 1.25
- **フレームワーク**: Echo v4
- **データベース**: PostgreSQL 17

## ローカルサーバーの立ち上げ

#### 全サービスを起動（推奨）

```bash
docker-compose up -d
```

このコマンドで以下が実行される：

- PostgreSQL データベースの起動
- データベースマイグレーションの実行
- API サーバーの起動

#### 個別に起動する場合

```bash
# PostgreSQLのみ起動
docker-compose up -d postgres

# マイグレーション実行
docker-compose up migrate

# APIサーバー起動
docker-compose up -d app
```

### マイグレーション

#### 新しいマイグレーションの追加

1. `db/migrations/`ディレクトリに新しいファイルを作成：


```bash
migrate create -ext sql -dir db/migrations -seq create_users_table
```

-ext	マイグレーションファイルの拡張子（今回は SQL）
-dir	マイグレーションファイルを作成する場所（指定したディレクトリが存在しなければ新規作成される）
-seq	マイグレーションファイルの名前

   - `000001_create_users_table.up.sql`
   - `000001_create_users_table.down.sql`

2. マイグレーションを実行：
   ```bash
   docker-compose build migrate
   docker-compose up migrate
   ```
