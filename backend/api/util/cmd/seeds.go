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
		// A組織を作成
		orgA := client.Organization.Create().
			SetName("A組織").
			SaveX(ctx)

		// B組織を作成
		orgB := client.Organization.Create().
			SetName("B組織").
			SaveX(ctx)

		// Adminユーザーを作成し、A組織に割り当て
		admin := client.User.Create().
			SetUID("admin").
			SetName("管理者").
			SetEmail("admin@hoge.jp").
			SetPassword("$2a$10$HOn8WTUwZFj6CtT0rOktluNjyLjd1kennkRZWOmKn7TpBXmY7J8Qq").
			SetRoleType("admin").
			SetStatusType("active").
			SetOauthType("local").
			SetOrganization(orgA).
			SaveX(ctx)

		// プロファイル1を作成し、Adminに割り当て
		profile1 := client.Photo.Create().
			SetName("profile1").
			SetURL("http://example.com/profile1.jpg").
			SetUser(admin).
			SaveX(ctx)

		// プロファイル2を作成（割り当てなし）
		profile2 := client.Photo.Create().
			SetName("profile2").
			SetURL("http://example.com/profile2.jpg").
			SaveX(ctx)

		// エラーハンドリング
		if admin == nil || profile1 == nil || profile2 == nil || orgA == nil || orgB == nil {
			log.Fatalf("failed creating seed data")
		}

		log.Println("Admin user created with profile1 and assigned to A組織")
	} else {
		log.Println("Admin user already exists")
	}
}
