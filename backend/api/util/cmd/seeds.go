package main

import (
	"context"
	"log"

	"api/ent"
	"api/ent/migrate"
	"api/ent/user"
	mysql "api/services/db"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	// シングルトン化されたデータベースクライアントを取得
	client, _, err := mysql.GetDatabaseClient()
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer client.Close() // プログラム終了時にクライアントを閉じる

	// データベースのマイグレーションを実行
	if err := client.Schema.Create(context.Background(), migrate.WithGlobalUniqueID(true)); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

	// シードデータを作成
	createSeedData(client)
}

// createSeedData inserts seed data into the database.
func createSeedData(client *ent.Client) {
	ctx := context.Background()

	// ユーザーが既に存在するか確認
	exists, err := client.User.Query().Where(user.UID("admin")).Exist(ctx)
	if err != nil {
		log.Fatalf("failed checking if user exists: %v", err)
	}

	// ユーザーが存在しない場合にのみ作成
	if !exists {
		u := client.User.Create().
			SetUID("admin").
			SetName("管理者").
			SetEmail("admin@hoge.jp").
			SetPassword("$2a$10$HOn8WTUwZFj6CtT0rOktluNjyLjd1kennkRZWOmKn7TpBXmY7J8Qq").
			SetRoleType("admin").
			SetStatusType("active").
			SetOauthType("local").
			SaveX(ctx)

		// エラーハンドリング
		if u == nil {
			log.Fatalf("failed creating user: %v", u)
		}

		log.Println("user created:", u)
	} else {
		log.Println("user already exists")
	}
}
