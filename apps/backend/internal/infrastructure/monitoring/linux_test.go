package monitoring

import (
	"os"
	"path/filepath"
	"testing"
)

func TestReadCPUStatParsesProcFormat(t *testing.T) {
	path := writeTempFile(t, "cpu  100 20 30 400 10 0 0 0 0 0\ncpu0 50 10 15 200 5 0 0 0 0 0\n")

	total, idle, err := readCPUStat(path)
	if err != nil {
		t.Fatalf("expected cpu stat to parse, got error: %v", err)
	}

	if total != 560 {
		t.Fatalf("expected total 560, got %d", total)
	}

	if idle != 410 {
		t.Fatalf("expected idle 410, got %d", idle)
	}
}

func TestReadMemInfoParsesLinuxMeminfo(t *testing.T) {
	path := writeTempFile(t, "MemTotal:       1048576 kB\nMemAvailable:    524288 kB\n")

	total, available, err := readMemInfo(path)
	if err != nil {
		t.Fatalf("expected meminfo to parse, got error: %v", err)
	}

	if total != 1073741824 {
		t.Fatalf("expected total bytes, got %d", total)
	}

	if available != 536870912 {
		t.Fatalf("expected available bytes, got %d", available)
	}
}

func writeTempFile(t *testing.T, content string) string {
	t.Helper()

	dir := t.TempDir()
	path := filepath.Join(dir, "sample.txt")
	if err := os.WriteFile(path, []byte(content), 0o600); err != nil {
		t.Fatalf("expected temp file to be written, got error: %v", err)
	}

	return path
}
