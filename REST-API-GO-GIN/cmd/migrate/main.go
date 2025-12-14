package main

import (
	"database/sql" // Go 標準庫的資料庫抽象層，類似 Java 的 JDBC
	"log"
	"os"

	"github.com/golang-migrate/migrate/v4"                          // 主遷移套件，相當於 Flyway/Liquibase 的核心
	Sqlite3 "github.com/golang-migrate/migrate/v4/database/sqlite3" // SQLite 資料庫驅動（golang-migrate 專用），用別名 Sqlite3 是因為原套件名小寫（私有）
	"github.com/golang-migrate/migrate/v4/source/file"              // 遷移檔案的來源驅動：從本地檔案夾讀取 .up.sql 和 .down.sql
	_ "modernc.org/sqlite"                                          // SQLite 底層驅動，前面加 _ 是為了觸發它的 init() 註冊自己（前面已詳細解釋）
)

func main() {
	// 檢查命令列參數是否提供遷移方向（UP 或 DOWN）
	if len(os.Args) < 2 {
		log.Fatal("Pleasr provide a migration direction : 'UP' or 'DOWN' ")
	}

	// 取得第一個參數作為遷移方向（例如執行 go run main.go UP）
	direction := os.Args[1]

	// 開啟 SQLite 資料庫連線（檔案路徑 ./data.db）
	// "sqlite3" 是驅動名稱，因為前面 _ "github.com/mattn/go-sqlite3" 已經註冊過了
	db, err := sql.Open("sqlite", "./data.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close() // 程式結束前關閉連線

	// 建立 golang-migrate 專用的 SQLite 資料庫實例（driver instance）
	// 需要用別名 Sqlite3 呼叫，因為原套件名小寫（Go 存取控制規則）
	instance, err := Sqlite3.WithInstance(db, &Sqlite3.Config{})
	if err != nil {
		log.Fatal(err)
	}

	// 建立遷移檔案的來源實例
	// 這裡指定遷移檔案放在 cmd/migrate/migrations 資料夾（注意你的路徑要正確）
	// 舊版寫法是 (&file.File{}).Open，新的推薦用 iofs 或其他方式，但這行還能用
	fSrc, err := (&file.File{}).Open("cmd/migrate/migrations")
	if err != nil {
		log.Fatal(err)
	}

	// 建立 migrate 核心物件
	// 參數：source driver 名稱 ("file")、source 實例、database driver 名稱 ("sqlite3")、database 實例
	m, err := migrate.NewWithInstance("file", fSrc, "sqlite3", instance)
	if err != nil {
		log.Fatal(err)
	}

	// 根據方向執行遷移
	switch direction {
	case "UP":
		// 執行所有尚未執行的 .up.sql 檔案
		// 如果沒有變更會回傳 ErrNoChange，我們忽略它
		if err := m.Up(); err != nil && err != migrate.ErrNoChange {
			log.Fatal(err)
		}
	case "DOWN":
		// 回滾上一個遷移（執行最新的 .down.sql）
		if err := m.Down(); err != nil && err != migrate.ErrNoChange {
			log.Fatal(err)
		}
	default:
		log.Fatal("Unknown migration direction. Use 'UP' or 'DOWN'.")
	}

	// 程式正常結束（如果有錯誤會在上面 log.Fatal 直接結束）
	log.Println("Migration completed successfully!")
}
