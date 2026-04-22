package config

type SchemeType string

const (
	SchemeLocal  SchemeType = "local"
	SchemeRemote SchemeType = "remote"
)

type Scheme struct {
	ID        string     `json:"id"`
	Name      string     `json:"name"`
	Type      SchemeType `json:"type"`
	Content   string     `json:"content,omitempty"` // 本地方案内容
	URL       string     `json:"url,omitempty"`     // 远程 URL（Sprint 3 使用）
	Enabled   bool       `json:"enabled"`
	CreatedAt string     `json:"createdAt"` // ISO 8601 format string
	UpdatedAt string     `json:"updatedAt"` // ISO 8601 format string
}

type Config struct {
	Schemes      []Scheme `json:"schemes"`
	ActiveScheme string   `json:"activeScheme"` // 当前生效的方案 ID
}
