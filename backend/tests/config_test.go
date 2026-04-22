package tests

import (
	"path/filepath"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"gohosts/backend/config"
	"gohosts/backend/db"
)

func TestConfigService_SQLiteCRUD(t *testing.T) {
	// 每个测试使用独立的数据库路径，测试前重置单例
	db.Reset()
	tmpDir := t.TempDir()
	dbPath := filepath.Join(tmpDir, "config.db")

	svc, err := config.NewService(dbPath)
	require.NoError(t, err)
	defer db.Close()

	// 初始应无任何方案
	schemes, err := svc.GetSchemes()
	require.NoError(t, err)
	assert.Empty(t, schemes)

	// 创建方案
	scheme := config.Scheme{
		ID:        "scheme-1",
		Name:      "开发环境",
		Type:      config.SchemeLocal,
		Content:   "127.0.0.1 dev.local",
		Enabled:   true,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	err = svc.SaveScheme(scheme)
	require.NoError(t, err)

	// 查询所有方案
	schemes, err = svc.GetSchemes()
	require.NoError(t, err)
	assert.Len(t, schemes, 1)
	assert.Equal(t, "开发环境", schemes[0].Name)
	assert.True(t, schemes[0].Enabled)

	// 查询单个方案
	got, err := svc.GetScheme("scheme-1")
	require.NoError(t, err)
	require.NotNil(t, got)
	assert.Equal(t, "开发环境", got.Name)

	// 查询不存在的方案
	notFound, err := svc.GetScheme("nonexistent")
	require.NoError(t, err)
	assert.Nil(t, notFound)

	// 更新方案
	scheme.Name = "更新后的开发环境"
	scheme.Enabled = false
	err = svc.SaveScheme(scheme)
	require.NoError(t, err)

	updated, err := svc.GetScheme("scheme-1")
	require.NoError(t, err)
	require.NotNil(t, updated)
	assert.Equal(t, "更新后的开发环境", updated.Name)
	assert.False(t, updated.Enabled)

	// 删除方案
	err = svc.DeleteScheme("scheme-1")
	require.NoError(t, err)

	schemes, err = svc.GetSchemes()
	require.NoError(t, err)
	assert.Empty(t, schemes)
}

func TestConfigService_LoadAndSave(t *testing.T) {
	db.Reset()
	tmpDir := t.TempDir()
	dbPath := filepath.Join(tmpDir, "config.db")

	svc, err := config.NewService(dbPath)
	require.NoError(t, err)
	defer db.Close()

	// 初始加载应返回空配置
	cfg, err := svc.Load()
	require.NoError(t, err)
	assert.Empty(t, cfg.Schemes)
	assert.Empty(t, cfg.ActiveScheme)

	// 添加多个方案并设置当前生效方案
	scheme1 := config.Scheme{
		ID:        "1",
		Name:      "开发环境",
		Type:      config.SchemeLocal,
		Content:   "127.0.0.1 dev.local",
		Enabled:   true,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	scheme2 := config.Scheme{
		ID:        "2",
		Name:      "测试环境",
		Type:      config.SchemeLocal,
		Content:   "127.0.0.1 test.local",
		Enabled:   false,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	cfg.Schemes = []config.Scheme{scheme1, scheme2}
	cfg.ActiveScheme = "1"
	err = svc.Save(cfg)
	require.NoError(t, err)

	// 重新加载应包含所有方案和生效设置
	cfg2, err := svc.Load()
	require.NoError(t, err)
	assert.Len(t, cfg2.Schemes, 2)
	assert.Equal(t, "开发环境", cfg2.Schemes[0].Name)
	assert.Equal(t, "测试环境", cfg2.Schemes[1].Name)
	assert.Equal(t, "1", cfg2.ActiveScheme)
}

func TestConfigService_Setting(t *testing.T) {
	db.Reset()
	tmpDir := t.TempDir()
	dbPath := filepath.Join(tmpDir, "config.db")

	svc, err := config.NewService(dbPath)
	require.NoError(t, err)
	defer db.Close()

	// 获取不存在的设置项
	val, err := svc.GetSetting("activeScheme")
	require.NoError(t, err)
	assert.Empty(t, val)

	// 设置值
	err = svc.SetSetting("activeScheme", "scheme-1")
	require.NoError(t, err)

	// 读取值
	val, err = svc.GetSetting("activeScheme")
	require.NoError(t, err)
	assert.Equal(t, "scheme-1", val)

	// 更新值
	err = svc.SetSetting("activeScheme", "scheme-2")
	require.NoError(t, err)

	val, err = svc.GetSetting("activeScheme")
	require.NoError(t, err)
	assert.Equal(t, "scheme-2", val)
}
