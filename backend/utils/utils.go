package utils

import (
	"context"
	"github.com/wailsapp/wails/v2/pkg/runtime"
)

func GetEnv(ctx context.Context) runtime.EnvironmentInfo {
	return runtime.Environment(ctx)
}
