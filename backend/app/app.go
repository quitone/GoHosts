package app

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"runtime"

	"gohosts/backend/config"
	"gohosts/backend/hosts"
)

type App struct {
	ctx           context.Context
	hostsService  *hosts.Service
	configService *config.Service
}

// AppOption 是 App 的可选配置函数。
type AppOption func(*App)

// WithHostsService 注入自定义的 hosts.Service（通常用于测试）。
func WithHostsService(svc *hosts.Service) AppOption {
	return func(a *App) {
		a.hostsService = svc
	}
}

// WithConfigService 注入自定义的 config.Service（通常用于测试）。
func WithConfigService(svc *config.Service) AppOption {
	return func(a *App) {
		a.configService = svc
	}
}

// GetSystemHostsPath 返回当前操作系统的 hosts 文件路径。
func GetSystemHostsPath() string {
	switch runtime.GOOS {
	case "windows":
		return "C:\\Windows\\System32\\drivers\\etc\\hosts"
	case "linux", "darwin", "freebsd", "netbsd", "openbsd":
		return "/etc/hosts"
	default:
		// 对于未知操作系统，返回常见的 Linux 路径
		return "/etc/hosts"
	}
}

// NewApp 创建一个新的 App 实例，初始化 hosts 和配置服务。
// 默认数据库路径为用户主目录下的 ~/.gohosts/config.db。
// 可通过 WithHostsService、WithConfigService 等选项注入自定义服务（常用于测试）。
func NewApp(opts ...AppOption) *App {
	app := &App{}

	for _, opt := range opts {
		opt(app)
	}

	// 如果未注入 hostsService，创建默认的
	if app.hostsService == nil {
		app.hostsService = hosts.NewService(GetSystemHostsPath())
	}

	// 如果未注入 configService，创建默认的
	if app.configService == nil {
		home, _ := os.UserHomeDir()
		dbPath := filepath.Join(home, ".gohosts", "config.db")
		_ = os.MkdirAll(filepath.Dir(dbPath), 0755)
		configSvc, err := config.NewService(dbPath)
		if err != nil {
			panic(fmt.Sprintf("failed to create config service: %v", err))
		}
		app.configService = configSvc
	}

	return app
}

// GetSystemHosts 返回系统当前 Hosts 内容。
func (a *App) GetSystemHosts() (string, error) {
	return a.hostsService.Read()
}

// GetSchemes 返回所有配置方案。
func (a *App) GetSchemes() ([]config.Scheme, error) {
	return a.configService.GetSchemes()
}

// SaveScheme 保存方案（创建或更新）。
func (a *App) SaveScheme(scheme config.Scheme) error {
	return a.configService.SaveScheme(scheme)
}

// DeleteScheme 删除方案。
func (a *App) DeleteScheme(id string) error {
	return a.configService.DeleteScheme(id)
}
