package tests

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"gohosts/backend/app"
	"gohosts/backend/config"
	"gohosts/backend/db"
	"gohosts/backend/hosts"
)

func TestApp_GetSystemHosts(t *testing.T) {
	// 重置单例以确保测试隔离
	db.Reset()

	tmpDir := t.TempDir()
	testPath := filepath.Join(tmpDir, "hosts")
	err := os.WriteFile(testPath, []byte("127.0.0.1 localhost\n::1 localhost\n"), 0644)
	require.NoError(t, err)

	dbPath := filepath.Join(tmpDir, "config.db")
	configSvc, err := config.NewService(dbPath)
	require.NoError(t, err)
	defer db.Close()

	hostsSvc := hosts.NewService(testPath)
	application := app.NewApp(
		app.WithHostsService(hostsSvc),
		app.WithConfigService(configSvc),
	)

	content, err := application.GetSystemHosts()
	require.NoError(t, err)
	assert.NotEmpty(t, content)
	assert.Contains(t, content, "localhost")
}

func TestApp_LoadAndSaveConfig(t *testing.T) {
	// 重置单例以确保测试隔离
	db.Reset()

	tmpDir := t.TempDir()
	testPath := filepath.Join(tmpDir, "hosts")
	err := os.WriteFile(testPath, []byte("127.0.0.1 localhost\n"), 0644)
	require.NoError(t, err)

	dbPath := filepath.Join(tmpDir, "config.db")
	configSvc, err := config.NewService(dbPath)
	require.NoError(t, err)
	defer db.Close()

	hostsSvc := hosts.NewService(testPath)
	application := app.NewApp(
		app.WithHostsService(hostsSvc),
		app.WithConfigService(configSvc),
	)

	// 测试 1: 初始加载应返回空配置
	cfg, err := application.LoadConfig()
	require.NoError(t, err)
	require.NotNil(t, cfg)
	assert.Empty(t, cfg.Schemes)
	assert.Empty(t, cfg.ActiveScheme)

	// 测试 2: 创建配置并保存
	cfg.Schemes = []config.Scheme{
		{
			ID:      "scheme-1",
			Name:    "开发环境",
			Type:    config.SchemeLocal,
			Content: "127.0.0.1 dev.local",
			Enabled: true,
		},
		{
			ID:      "scheme-2",
			Name:    "测试环境",
			Type:    config.SchemeLocal,
			Content: "127.0.0.1 test.local",
			Enabled: false,
		},
	}
	cfg.ActiveScheme = "scheme-1"

	err = application.SaveConfig(cfg)
	require.NoError(t, err)

	// 测试 3: 重新加载应包含所有数据
	loadedCfg, err := application.LoadConfig()
	require.NoError(t, err)
	require.NotNil(t, loadedCfg)
	assert.Len(t, loadedCfg.Schemes, 2)
	assert.Equal(t, "开发环境", loadedCfg.Schemes[0].Name)
	assert.Equal(t, "测试环境", loadedCfg.Schemes[1].Name)
	assert.Equal(t, "scheme-1", loadedCfg.ActiveScheme)

	// 测试 4: 更新配置
	loadedCfg.Schemes[0].Name = "更新后的开发环境"
	loadedCfg.ActiveScheme = "scheme-2"

	err = application.SaveConfig(loadedCfg)
	require.NoError(t, err)

	// 验证更新
	updatedCfg, err := application.LoadConfig()
	require.NoError(t, err)
	require.NotNil(t, updatedCfg)
	assert.Equal(t, "更新后的开发环境", updatedCfg.Schemes[0].Name)
	assert.Equal(t, "scheme-2", updatedCfg.ActiveScheme)
}
