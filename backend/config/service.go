package config

import (
	"database/sql"
	"fmt"
	"time"

	"gohosts/backend/db"
)

// Service 使用 SQLite 数据库存储 hosts 方案配置数据。
// 使用全局单例数据库连接进行管理。
type Service struct {
	dbPath string
}

// NewService 创建一个新的配置服务，使用指定的 SQLite 数据库路径。
// 数据库连接由全局单例管理，确保整个应用生命周期内只创建一次。
func NewService(dbPath string) (*Service, error) {
	_, err := db.GetInstance(dbPath)
	if err != nil {
		return nil, fmt.Errorf("initialize database %s: %w", dbPath, err)
	}

	return &Service{dbPath: dbPath}, nil
}

// getDB 返回全局单例数据库连接。
func (s *Service) getDB() (*sql.DB, error) {
	return db.GetInstance(s.dbPath)
}

// Load 从数据库加载完整配置（所有方案及当前生效方案 ID）。
func (s *Service) Load() (*Config, error) {
	schemes, err := s.GetSchemes()
	if err != nil {
		return nil, fmt.Errorf("load schemes: %w", err)
	}

	activeScheme, _ := s.GetSetting("activeScheme")

	return &Config{
		Schemes:      schemes,
		ActiveScheme: activeScheme,
	}, nil
}

// Save 将完整配置保存到数据库（使用事务保证一致性）。
func (s *Service) Save(cfg *Config) error {
	database, err := s.getDB()
	if err != nil {
		return err
	}

	tx, err := database.Begin()
	if err != nil {
		return fmt.Errorf("begin transaction: %w", err)
	}
	defer func() { _ = tx.Rollback() }()

	if _, err := tx.Exec("DELETE FROM schemes"); err != nil {
		return fmt.Errorf("clear schemes: %w", err)
	}

	stmt, err := tx.Prepare(`
		INSERT INTO schemes (id, name, type, content, url, enabled, created_at, updated_at)
		VALUES (?, ?, ?, ?, ?, ?, ?, ?)
	`)
	if err != nil {
		return fmt.Errorf("prepare insert statement: %w", err)
	}
	defer func() { _ = stmt.Close() }()

	for _, scheme := range cfg.Schemes {
		enabled := 0
		if scheme.Enabled {
			enabled = 1
		}
		_, err := stmt.Exec(
			scheme.ID, scheme.Name, string(scheme.Type),
			scheme.Content, scheme.URL, enabled,
			scheme.CreatedAt, scheme.UpdatedAt,
		)
		if err != nil {
			return fmt.Errorf("insert scheme %s: %w", scheme.ID, err)
		}
	}

	if cfg.ActiveScheme != "" {
		if _, err := tx.Exec(`
			INSERT INTO settings (key, value) VALUES (?, ?)
			ON CONFLICT(key) DO UPDATE SET value = excluded.value
		`, "activeScheme", cfg.ActiveScheme); err != nil {
			return fmt.Errorf("save active scheme: %w", err)
		}
	}

	if err := tx.Commit(); err != nil {
		return fmt.Errorf("commit transaction: %w", err)
	}
	return nil
}

// GetSchemes 返回数据库中存储的所有方案列表，按创建时间排序。
func (s *Service) GetSchemes() ([]Scheme, error) {
	database, err := s.getDB()
	if err != nil {
		return nil, err
	}

	rows, err := database.Query(`
		SELECT id, name, type, content, url, enabled, created_at, updated_at
		FROM schemes ORDER BY created_at
	`)
	if err != nil {
		return nil, fmt.Errorf("query schemes: %w", err)
	}
	defer func() { _ = rows.Close() }()

	var schemes []Scheme
	for rows.Next() {
		var scheme Scheme
		var enabled int
		err := rows.Scan(
			&scheme.ID, &scheme.Name, &scheme.Type,
			&scheme.Content, &scheme.URL, &enabled,
			&scheme.CreatedAt, &scheme.UpdatedAt,
		)
		if err != nil {
			return nil, fmt.Errorf("scan scheme: %w", err)
		}
		scheme.Enabled = enabled != 0
		schemes = append(schemes, scheme)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("iterate schemes: %w", err)
	}

	return schemes, nil
}

// GetScheme 根据 ID 从数据库查询单个方案。
// 如果方案不存在，返回 nil 和 nil error。
func (s *Service) GetScheme(id string) (*Scheme, error) {
	database, err := s.getDB()
	if err != nil {
		return nil, err
	}

	var scheme Scheme
	var enabled int
	err = database.QueryRow(`
		SELECT id, name, type, content, url, enabled, created_at, updated_at
		FROM schemes WHERE id = ?
	`, id).Scan(
		&scheme.ID, &scheme.Name, &scheme.Type,
		&scheme.Content, &scheme.URL, &enabled,
		&scheme.CreatedAt, &scheme.UpdatedAt,
	)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, fmt.Errorf("query scheme %s: %w", id, err)
	}
	scheme.Enabled = enabled != 0
	return &scheme, nil
}

// SaveScheme 保存单个方案到数据库（不存在则创建，存在则更新）。
func (s *Service) SaveScheme(scheme Scheme) error {
	database, err := s.getDB()
	if err != nil {
		return err
	}

	scheme.UpdatedAt = time.Now().Format(time.RFC3339)
	if scheme.CreatedAt == "" {
		scheme.CreatedAt = scheme.UpdatedAt
	}

	enabled := 0
	if scheme.Enabled {
		enabled = 1
	}

	_, err = database.Exec(`
		INSERT INTO schemes (id, name, type, content, url, enabled, created_at, updated_at)
		VALUES (?, ?, ?, ?, ?, ?, ?, ?)
		ON CONFLICT(id) DO UPDATE SET
			name = excluded.name,
			type = excluded.type,
			content = excluded.content,
			url = excluded.url,
			enabled = excluded.enabled,
			updated_at = excluded.updated_at
	`, scheme.ID, scheme.Name, string(scheme.Type),
		scheme.Content, scheme.URL, enabled,
		scheme.CreatedAt, scheme.UpdatedAt,
	)
	if err != nil {
		return fmt.Errorf("save scheme %s: %w", scheme.ID, err)
	}
	return nil
}

// DeleteScheme 从数据库删除指定 ID 的方案。
func (s *Service) DeleteScheme(id string) error {
	database, err := s.getDB()
	if err != nil {
		return err
	}

	_, err = database.Exec("DELETE FROM schemes WHERE id = ?", id)
	if err != nil {
		return fmt.Errorf("delete scheme %s: %w", id, err)
	}
	return nil
}

// GetSetting 根据 key 查询设置项的值。
// 如果设置项不存在，返回空字符串和 nil error。
func (s *Service) GetSetting(key string) (string, error) {
	database, err := s.getDB()
	if err != nil {
		return "", err
	}

	var value string
	err = database.QueryRow("SELECT value FROM settings WHERE key = ?", key).Scan(&value)
	if err == sql.ErrNoRows {
		return "", nil
	}
	if err != nil {
		return "", fmt.Errorf("get setting %s: %w", key, err)
	}
	return value, nil
}

// SetSetting 保存或更新设置项的值。
func (s *Service) SetSetting(key, value string) error {
	database, err := s.getDB()
	if err != nil {
		return err
	}

	_, err = database.Exec(`
		INSERT INTO settings (key, value) VALUES (?, ?)
		ON CONFLICT(key) DO UPDATE SET value = excluded.value
	`, key, value)
	if err != nil {
		return fmt.Errorf("set setting %s: %w", key, err)
	}
	return nil
}

// Close 关闭全局数据库连接。
// 注意：由于使用单例模式，关闭后所有 Service 实例将无法使用。
func (s *Service) Close() error {
	return db.Close()
}
