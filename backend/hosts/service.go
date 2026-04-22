package hosts

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"time"
)

// Service 提供 hosts 文件的读写和备份功能。
type Service struct {
	hostsPath string
	backupDir string
	logger    *log.Logger
}

// Option 是 Service 的可选配置函数。
type Option func(*Service)

// WithBackupDir 设置备份目录路径。
func WithBackupDir(dir string) Option {
	return func(s *Service) {
		s.backupDir = dir
	}
}

// WithLogger 设置自定义日志记录器。
func WithLogger(logger *log.Logger) Option {
	return func(s *Service) {
		s.logger = logger
	}
}

// NewService 创建一个新的 HostsService。
// hostsPath 指定系统 hosts 文件路径；opts 可传入 WithBackupDir、WithLogger 等选项。
// 默认备份目录为用户主目录下的 .gohosts/backups。
func NewService(hostsPath string, opts ...Option) *Service {
	s := &Service{
		hostsPath: hostsPath,
		backupDir: filepath.Join(os.Getenv("HOME"), ".gohosts", "backups"),
		logger:    log.New(os.Stderr, "[hosts] ", log.LstdFlags),
	}
	for _, opt := range opts {
		opt(s)
	}
	return s
}

// logf 内部日志输出辅助方法。
func (s *Service) logf(format string, v ...interface{}) {
	if s.logger != nil {
		s.logger.Printf(format, v...)
	}
}

// Read 读取系统 hosts 文件内容并返回字符串。
// 如果文件不存在或读取失败，返回相应的错误信息。
func (s *Service) Read() (string, error) {
	s.logf("reading hosts file: %s", s.hostsPath)
	content, err := os.ReadFile(s.hostsPath)
	if err != nil {
		return "", fmt.Errorf("read hosts file %s: %w", s.hostsPath, err)
	}
	s.logf("successfully read hosts file, size: %d bytes", len(content))
	return string(content), nil
}

// Write 将指定内容写入 hosts 文件。
// 写入权限默认为 0644（所有者可读写，其他用户只读）。
func (s *Service) Write(content string) error {
	s.logf("writing hosts file: %s", s.hostsPath)
	if err := os.WriteFile(s.hostsPath, []byte(content), 0644); err != nil {
		return fmt.Errorf("write hosts file %s: %w", s.hostsPath, err)
	}
	s.logf("successfully wrote hosts file, size: %d bytes", len(content))
	return nil
}

// Backup 创建当前 hosts 文件的备份，保存到独立的备份目录中。
// 备份文件按时间戳命名（格式：hosts.YYYYMMDD_HHMMSS.backup），以避免文件名冲突。
// 返回创建的备份文件绝对路径。
func (s *Service) Backup() (string, error) {
	s.logf("backing up hosts file from %s to %s", s.hostsPath, s.backupDir)

	if err := os.MkdirAll(s.backupDir, 0755); err != nil {
		return "", fmt.Errorf("create backup dir %s: %w", s.backupDir, err)
	}

	content, err := os.ReadFile(s.hostsPath)
	if err != nil {
		return "", fmt.Errorf("read hosts for backup: %w", err)
	}

	timestamp := time.Now().Format("20060102_150405")
	backupPath := filepath.Join(s.backupDir, fmt.Sprintf("hosts.%s.backup", timestamp))
	if err := os.WriteFile(backupPath, content, 0644); err != nil {
		return "", fmt.Errorf("write backup file %s: %w", backupPath, err)
	}

	s.logf("successfully created backup: %s", backupPath)
	return backupPath, nil
}
