//go:build ignore

package main

import (
	"context"
	"log"
	"os"

	"api/ent/migrate"
	mysql "api/services/db"

	atlas "ariga.io/atlas/sql/migrate"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql/schema"
)

func main() {
	// シングルトン化されたデータベースクライアントとDSNを取得
	client, dsn, err := mysql.GetDatabaseClient()
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer client.Close() // プログラム終了時にクライアントを閉じる

	ctx := context.Background()

	// ローカルのマイグレーションディレクトリを作成し、Atlasのマイグレーションファイルフォーマットを理解できるようにする
	dir, err := atlas.NewLocalDir("ent/migrate/migrations")
	if err != nil {
		log.Fatalf("failed creating atlas migration directory: %v", err)
	}

	// マイグレーションのオプションを設定
	opts := []schema.MigrateOption{
		schema.WithDir(dir),                          // マイグレーションディレクトリを指定
		schema.WithMigrationMode(schema.ModeInspect), // 現在のスキーマ状態に基づいてマイグレーションを生成
		schema.WithDialect(dialect.MySQL),            // 使用するEntのダイアレクトを指定
		schema.WithFormatter(atlas.DefaultFormatter),
		schema.WithDropColumn(true),
		schema.WithDropIndex(true),
	}

	// コマンドライン引数のチェック
	if len(os.Args) != 2 {
		log.Fatalln("migration name is required. Use: 'go run -mod=mod main.go <name>'")
	}

	// Atlasのサポートを使用してMySQL用のマイグレーションを生成（上記で指定したEntのダイアレクトオプションに注目）
	err = migrate.NamedDiff(ctx, dsn, os.Args[1], opts...)
	if err != nil {
		log.Fatalf("failed generating migration file: %v", err)
	}
}
