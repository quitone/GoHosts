package db

import (
	"database/sql"
	"fmt"
	"sync"

	_ "github.com/mattn/go-sqlite3"
)

var (
	instance *sql.DB
	once     sync.Once
	initErr  error
)

// GetInstance 返回全局唯一的数据库连接实例（单例模式）。
// 首次调用时会初始化数据库连接并创建表结构。
// 该方法是线程安全的，多个 goroutine 并发调用时只会初始化一次。
func GetInstance(dbPath string) (*sql.DB, error) {
	once.Do(func() {
		instance, initErr = initDB(dbPath)
	})
	return instance, initErr
}

// initDB 初始化数据库连接并创建表结构。
func initDB(dbPath string) (*sql.DB, error) {
	db, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		return nil, fmt.Errorf("open sqlite database %s: %w", dbPath, err)
	}

	if err := db.Ping(); err != nil {
		_ = db.Close()
		return nil, fmt.Errorf("ping sqlite database: %w", err)
	}

	if err := initSchema(db); err != nil {
		_ = db.Close()
		return nil, fmt.Errorf("init database schema: %w", err)
	}

	return db, nil
}

// initSchema 初始化数据库表结构。
func initSchema(db *sql.DB) error {
	queries := []string{
		`CREATE TABLE IF NOT EXISTS schemes (
			id TEXT PRIMARY KEY,
			name TEXT NOT NULL,
			type TEXT NOT NULL,
			content TEXT,
			url TEXT,
			enabled INTEGER NOT NULL DEFAULT 0,
			created_at DATETIME NOT NULL,
			updated_at DATETIME NOT NULL
		)`,
		`CREATE TABLE IF NOT EXISTS settings (
			key TEXT PRIMARY KEY,
			value TEXT NOT NULL
		)`,
	}

	for _, q := range queries {
		if _, err := db.Exec(q); err != nil {
			return fmt.Errorf("execute schema query: %w", err)
		}
	}
	return nil
}

// Close 关闭全局数据库连接。
// 应在程序退出时调用此方法释放资源。
// 多次调用是安全的，只有第一次调用会真正关闭连接。
func Close() error {
	if instance != nil {
		err := instance.Close()
		instance = nil
		return err
	}
	return nil
}

// Reset 重置单例状态（仅用于测试）。
// 允许测试代码在每次测试时重新初始化数据库连接。
func Reset() {
	instance = nil
	once = sync.Once{}
	initErr = nil
}
