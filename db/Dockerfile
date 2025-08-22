FROM alpine:3.20

# golang-migrateをインストール
RUN apk add --no-cache curl bash postgresql-client

# golang-migrateをダウンロードしてインストール
RUN curl -L https://github.com/golang-migrate/migrate/releases/download/v4.16.2/migrate.linux-amd64.tar.gz | tar xvz \
    && mv migrate /usr/local/bin/migrate \
    && chmod +x /usr/local/bin/migrate

WORKDIR /migrations

# マイグレーションファイルをコピー
COPY db/migrations/ .

# マイグレーション実行用のスクリプト
COPY scripts/migrate.sh /migrate.sh
RUN chmod +x /migrate.sh

CMD ["/migrate.sh"]
