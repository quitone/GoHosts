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
