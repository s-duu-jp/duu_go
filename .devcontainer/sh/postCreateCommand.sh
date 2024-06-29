#!/bin/sh

# Backend Setting
cd /workspace/backend/api
go mod tidy

# マイグレーションの実行
max_retries=10
attempt=1

until atlas migrate apply --dir "file://ent/migrate/migrations" --url "mysql://root:@mysql:3306/duu_go_dev"
do
  if [ $attempt -ge $max_retries ]
  then
    echo "Migration failed after $attempt attempts."
    exit 1
  fi

  echo "Migration failed, retrying in 1 second... ($attempt/$max_retries)"
  attempt=$((attempt + 1))
  sleep 3
done

echo "Migration succeeded."

# seedsの実行
go run ./util/cmd/seeds.go
