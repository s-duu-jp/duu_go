package mysql

import (
	"database/sql"
	"fmt"
	"log"
	"sync"

	"api/config/env"
	"api/ent"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/mattn/go-sqlite3"
)

var (
	Client     *ent.Client
	Dsn        string
	dbOnce     sync.Once
	ensureOnce sync.Once
	connectErr error
	ensureErr  error
)

// GetDatabaseClient はシングルトンとしてデータベース接続を提供します。
func GetDatabaseClient() (*ent.Client, string, error) {
	dbOnce.Do(func() {
		connectErr = connectDatabase()
	})
	return Client, Dsn, connectErr
}

// connectDatabase は環境変数を読み込み、データベースが存在しない場合に作成し、Entを使用してデータベースに接続します。
func connectDatabase() error {
	// 環境変数の読み込み
	cfg, err := env.GetConfig()
	if err != nil {
		log.Fatal("Failed to load config: ", err)
		return err
	}

	dbUser := cfg["MYSQL_USER"]
	dbPassword := cfg["MYSQL_PASSWORD"]
	dbName := cfg["MYSQL_NAME"]
	dbHost := cfg["MYSQL_HOST"]
	dbPort := cfg["MYSQL_PORT"]

	// 本番環境でない場合、データベースの存在確認と作成を行う
	if cfg["ENV"] != "production" {
		ensureOnce.Do(func() {
			ensureErr = ensureDatabaseExists(dbUser, dbPassword, dbHost, dbPort, dbName)
		})
		if ensureErr != nil {
			return ensureErr
		}
	}

	// Entを使用してデータベースに接続
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=True", dbUser, dbPassword, dbHost, dbPort, dbName)
	client, err := ent.Open("mysql", dsn)
	if err != nil {
		log.Fatalf("Failed opening connection to mysql: %v", err)
		return err
	}

	Dsn = fmt.Sprintf("mysql://%s:%s@%s:%s/%s", dbUser, dbPassword, dbHost, dbPort, dbName)
	Client = client
	return nil
}

// ensureDatabaseExists はデータベースが存在しない場合に作成します。
func ensureDatabaseExists(user, password, host, port, dbName string) error {
	// データベース名が空でないことを確認
	if dbName == "" {
		log.Fatal("Database name is empty")
		return fmt.Errorf("database name is empty")
	}

	// データベース名を含まないDSN
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/", user, password, host, port)

	// データベースサーバーに接続
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal("Failed to connect to database server:", err)
		return err
	}
	defer db.Close()

	// データベースが存在しない場合に作成
	createDBSQL := fmt.Sprintf("CREATE DATABASE IF NOT EXISTS `%s` CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci;", dbName)
	_, err = db.Exec(createDBSQL)
	if err != nil {
		log.Fatal("Failed to create database:", err)
		return err
	}
	return nil
}
