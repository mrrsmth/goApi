package utils_test

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
	"go_api_template/pkg/utils"
)

func TestLogger_Success(t *testing.T) {
	var buf bytes.Buffer
	logger := utils.NewLogger()
	logger.SuccessLogger.SetOutput(&buf) // set output to buffer

	logger.Success("test message")

	output := buf.String()
	assert.Contains(t, output, "SUCCESS")
	assert.Contains(t, output, "test message")
}

func TestLogger_Error(t *testing.T) {
	var buf bytes.Buffer
	logger := utils.NewLogger()
	logger.ErrorLogger.SetOutput(&buf) // set output to buffer

	logger.Error("test message")

	output := buf.String()
	assert.Contains(t, output, "ERROR")
	assert.Contains(t, output, "test message")
}

func TestLogger_Info(t *testing.T) {
	var buf bytes.Buffer
	logger := utils.NewLogger()
	logger.InfoLogger.SetOutput(&buf) // set output to buffer

	logger.Info("test message")

	output := buf.String()
	assert.Contains(t, output, "INFO")
	assert.Contains(t, output, "test message")
}

func TestLogger_Warn(t *testing.T) {
	var buf bytes.Buffer
	logger := utils.NewLogger()
	logger.WarnLogger.SetOutput(&buf) // set output to buffer

	logger.Warn("test message")

	output := buf.String()
	assert.Contains(t, output, "WARN")
	assert.Contains(t, output, "test message")
}

func TestLogger_Successf(t *testing.T) {
	var buf bytes.Buffer
	logger := utils.NewLogger()
	logger.SuccessLogger.SetOutput(&buf) // set output to buffer

	logger.Successf("test message %s", "formatted")

	output := buf.String()
	assert.Contains(t, output, "SUCCESS")
	assert.Contains(t, output, "test message formatted")
}

func TestLogger_Errorf(t *testing.T) {
	var buf bytes.Buffer
	logger := utils.NewLogger()
	logger.ErrorLogger.SetOutput(&buf) // set output to buffer

	logger.Errorf("test message %s", "formatted")

	output := buf.String()
	assert.Contains(t, output, "ERROR")
	assert.Contains(t, output, "test message formatted")
}

func TestLogger_Infof(t *testing.T) {
	var buf bytes.Buffer
	logger := utils.NewLogger()
	logger.InfoLogger.SetOutput(&buf) // set output to buffer

	logger.Infof("test message %s", "formatted")

	output := buf.String()
	assert.Contains(t, output, "INFO")
	assert.Contains(t, output, "test message formatted")
}

func TestLogger_Warnf(t *testing.T) {
	var buf bytes.Buffer
	logger := utils.NewLogger()
	logger.WarnLogger.SetOutput(&buf) // set output to buffer

	logger.Warnf("test message %s", "formatted")

	output := buf.String()
	assert.Contains(t, output, "WARN")
	assert.Contains(t, output, "test message formatted")
}

func TestLogger_NewLogger(t *testing.T) {
	logger := utils.NewLogger()
	assert.NotNil(t, logger)
}
