package config

import "time"

type SchemeType string

const (
    SchemeLocal  SchemeType = "local"
    SchemeRemote SchemeType = "remote"
)

type Scheme struct {
    ID        string     `json:"id"`
    Name      string     `json:"name"`
    Type      SchemeType `json:"type"`
    Content   string     `json:"content,omitempty"`   // 本地方案内容
    URL       string     `json:"url,omitempty"`       // 远程 URL（Sprint 3 使用）
    Enabled   bool       `json:"enabled"`
    CreatedAt time.Time  `json:"createdAt"`
    UpdatedAt time.Time  `json:"updatedAt"`
}

type Config struct {
    Schemes      []Scheme `json:"schemes"`
    ActiveScheme string   `json:"activeScheme"` // 当前生效的方案 ID
}