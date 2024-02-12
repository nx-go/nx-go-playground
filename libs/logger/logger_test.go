package logger

import (
	"bytes"
	"log"
	"os"
	"strings"
	"testing"
)

func TestInfo(t *testing.T) {
	var buf bytes.Buffer
	log.SetOutput(&buf)
	defer func() {
		log.SetOutput(os.Stderr)
	}()
	Info("informative message")
	if !strings.Contains(buf.String(), "[INFO] informative message") {
		t.Errorf("Expected informative message, got %s", buf.String())
	}
}

func TestWarn(t *testing.T) {
	var buf bytes.Buffer
	log.SetOutput(&buf)
	defer func() {
		log.SetOutput(os.Stderr)
	}()
	Warn("warning message")
	if !strings.Contains(buf.String(), "[WARN] warning message") {
		t.Errorf("Expected warning message, got %s", buf.String())
	}
}

func TestDebug(t *testing.T) {
	var buf bytes.Buffer
	log.SetOutput(&buf)
	defer func() {
		log.SetOutput(os.Stderr)
	}()
	Debug("debug message")
	if !strings.Contains(buf.String(), "[DEBUG] debug message") {
		t.Errorf("Expected debug message, got %s", buf.String())
	}
}

func TestError(t *testing.T) {
	var buf bytes.Buffer
	log.SetOutput(&buf)
	defer func() {
		log.SetOutput(os.Stderr)
	}()
	Error("error message")
	if !strings.Contains(buf.String(), "[ERROR] error message") {
		t.Errorf("Expected error message, got %s", buf.String())
	}
}
