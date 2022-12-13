package log

import (
	"github.com/stretchr/testify/assert"
	"path/filepath"
	"testing"
)

func TestInitGlobalLogger(t *testing.T) {
	tempLogFile := filepath.Join(t.TempDir(), "tempLog.log")
	InitGlobalLogger(tempLogFile)
	Log.Info("Initialize log file")
	err := Lum.Rotate()
	if err != nil {
		Log.Fatal(err)
	}
	assert.FileExists(t, tempLogFile)
	CloseGlobalLoggerFile()
}
