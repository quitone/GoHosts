package tests

import (
	"log"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"gohosts/backend/hosts"
)

func TestService_Read(t *testing.T) {
	tmpDir := t.TempDir()
	testPath := filepath.Join(tmpDir, "hosts")
	content := []byte("127.0.0.1 localhost\n::1 localhost\n")
	err := os.WriteFile(testPath, content, 0644)
	require.NoError(t, err)

	svc := hosts.NewService(testPath)
	got, err := svc.Read()

	assert.NoError(t, err)
	assert.Equal(t, string(content), got)
}

func TestService_Read_NotFound(t *testing.T) {
	tmpDir := t.TempDir()
	testPath := filepath.Join(tmpDir, "nonexistent")

	svc := hosts.NewService(testPath)
	_, err := svc.Read()

	assert.Error(t, err)
	assert.True(t, os.IsNotExist(err) || strings.Contains(err.Error(), "no such file"))
}

func TestService_Write(t *testing.T) {
	tmpDir := t.TempDir()
	testPath := filepath.Join(tmpDir, "hosts")

	svc := hosts.NewService(testPath)
	newContent := "192.168.1.1 myapp.local"
	err := svc.Write(newContent)
	require.NoError(t, err)

	got, err := os.ReadFile(testPath)
	require.NoError(t, err)
	assert.Equal(t, newContent, string(got))

	// 验证文件权限（类 Unix 系统）
	info, err := os.Stat(testPath)
	require.NoError(t, err)
	if os.PathSeparator == '/' {
		mode := info.Mode().Perm()
		assert.Equal(t, os.FileMode(0644), mode)
	}
}

func TestService_Backup(t *testing.T) {
	tmpDir := t.TempDir()
	testPath := filepath.Join(tmpDir, "hosts")
	backupDir := filepath.Join(tmpDir, "backups")

	content := []byte("127.0.0.1 localhost")
	err := os.WriteFile(testPath, content, 0644)
	require.NoError(t, err)

	svc := hosts.NewService(testPath, hosts.WithBackupDir(backupDir))
	backupPath, err := svc.Backup()
	require.NoError(t, err)

	assert.FileExists(t, backupPath)
	assert.True(t, strings.HasPrefix(filepath.Base(backupPath), "hosts."))
	assert.True(t, strings.HasSuffix(backupPath, ".backup"))
	assert.Equal(t, backupDir, filepath.Dir(backupPath))

	backupContent, err := os.ReadFile(backupPath)
	require.NoError(t, err)
	assert.Equal(t, content, backupContent)
}

func TestService_Backup_WithLogger(t *testing.T) {
	tmpDir := t.TempDir()
	testPath := filepath.Join(tmpDir, "hosts")
	backupDir := filepath.Join(tmpDir, "backups")

	content := []byte("192.168.1.1 test.local")
	err := os.WriteFile(testPath, content, 0644)
	require.NoError(t, err)

	logger := log.New(os.Stdout, "[test] ", log.LstdFlags)
	svc := hosts.NewService(testPath,
		hosts.WithBackupDir(backupDir),
		hosts.WithLogger(logger),
	)

	backupPath, err := svc.Backup()
	require.NoError(t, err)
	assert.FileExists(t, backupPath)
}
