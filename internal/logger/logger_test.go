package logger

import (
	"log"
	"strings"
	"testing"
)

func TestInfo(t *testing.T) {
	const expected = `[INFO][logger] hello, 1`

	defaultLogger.SetFlags(log.Lmsgprefix)

	w := new(strings.Builder)
	SetOutput(w)

	Info("hello, %d", 1)
	SetLevel(LWarning)
	Info("hello, %d", 2)

	if got := strings.TrimSpace(w.String()); got != expected {
		t.Errorf("expected %q, got %q", expected, got)
	}
}
