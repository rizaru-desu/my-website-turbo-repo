package system

import (
	"strings"
	"testing"
	"time"
)

func TestHealthServiceSnapshotIncludesFormattedMetrics(t *testing.T) {
	service := NewHealthService(
		time.Now().Add(-(5*24*time.Hour + 12*time.Hour + 30*time.Minute)),
		".",
		stubMetricsProvider{
			cpuUsage: 12.5,
			memory: MemoryMetrics{
				UsedBytes:  64 * 1024 * 1024,
				TotalBytes: 1024 * 1024 * 1024,
			},
			storage: StorageMetrics{
				UsedBytes:  (uint64(21) * 1024 * 1024 * 1024) / 10,
				FreeBytes:  (uint64(128) * 1024 * 1024 * 1024) / 10,
				TotalBytes: (uint64(149) * 1024 * 1024 * 1024) / 10,
			},
		},
	)

	snapshot := service.Snapshot()

	if snapshot.Status != "operational" {
		t.Fatalf("expected status operational, got %q", snapshot.Status)
	}

	if !strings.HasSuffix(snapshot.SystemHealth.CPUUsage, "%") {
		t.Fatalf("expected CPU usage to be formatted as percentage, got %q", snapshot.SystemHealth.CPUUsage)
	}

	if snapshot.SystemHealth.Memory.Used == "" || snapshot.SystemHealth.Memory.Total == "" {
		t.Fatalf("expected memory usage fields to be populated, got %+v", snapshot.SystemHealth.Memory)
	}

	if !strings.HasSuffix(snapshot.SystemHealth.Memory.Percentage, "%") {
		t.Fatalf("expected memory percentage to be formatted, got %q", snapshot.SystemHealth.Memory.Percentage)
	}

	if snapshot.SystemHealth.Storage.Used == "" || snapshot.SystemHealth.Storage.Free == "" {
		t.Fatalf("expected storage usage fields to be populated, got %+v", snapshot.SystemHealth.Storage)
	}

	if !strings.HasSuffix(snapshot.SystemHealth.Storage.Percentage, "%") {
		t.Fatalf("expected storage percentage to be formatted, got %q", snapshot.SystemHealth.Storage.Percentage)
	}

	if snapshot.SystemHealth.Uptime != "5d 12h 30m" {
		t.Fatalf("expected uptime to be formatted, got %q", snapshot.SystemHealth.Uptime)
	}

	if snapshot.SystemHealth.Memory.Used != "64MB" {
		t.Fatalf("expected memory used to be formatted, got %q", snapshot.SystemHealth.Memory.Used)
	}

	if snapshot.SystemHealth.Storage.Used != "2.1GB" {
		t.Fatalf("expected storage used to be formatted, got %q", snapshot.SystemHealth.Storage.Used)
	}

	if snapshot.SystemHealth.Storage.Total != "14.9GB" {
		t.Fatalf("expected storage total to be formatted, got %q", snapshot.SystemHealth.Storage.Total)
	}
}

func TestHealthServiceSnapshotMarksWarningWhenResourcesAreHigh(t *testing.T) {
	service := NewHealthService(
		time.Now().Add(-time.Hour),
		"/data",
		stubMetricsProvider{
			cpuUsage: 91,
			memory: MemoryMetrics{
				UsedBytes:  9 * 1024 * 1024 * 1024,
				TotalBytes: 10 * 1024 * 1024 * 1024,
			},
			storage: StorageMetrics{
				UsedBytes:  90 * 1024 * 1024 * 1024,
				FreeBytes:  10 * 1024 * 1024 * 1024,
				TotalBytes: 100 * 1024 * 1024 * 1024,
			},
		},
	)

	snapshot := service.Snapshot()

	if snapshot.Status != "degraded" {
		t.Fatalf("expected status degraded, got %q", snapshot.Status)
	}

	if len(snapshot.Alerts) == 0 {
		t.Fatalf("expected alerts to be populated")
	}
}

func TestHealthServiceSnapshotMarksCriticalWhenStorageIsAlmostFull(t *testing.T) {
	service := NewHealthService(
		time.Now().Add(-time.Hour),
		"/data",
		stubMetricsProvider{
			cpuUsage: 20,
			memory: MemoryMetrics{
				UsedBytes:  2 * 1024 * 1024 * 1024,
				TotalBytes: 8 * 1024 * 1024 * 1024,
			},
			storage: StorageMetrics{
				UsedBytes:  63 * 1024 * 1024 * 1024,
				FreeBytes:  256 * 1024 * 1024,
				TotalBytes: (63 * 1024 * 1024 * 1024) + (256 * 1024 * 1024),
			},
		},
	)

	snapshot := service.Snapshot()

	if snapshot.Status != "critical" {
		t.Fatalf("expected status critical, got %q", snapshot.Status)
	}

	if len(snapshot.Alerts) == 0 || snapshot.Alerts[0].Resource != "storage" {
		t.Fatalf("expected storage alert, got %+v", snapshot.Alerts)
	}
}

func TestFormatBytesUsesReadableUnits(t *testing.T) {
	if got := formatBytes(64 * 1024 * 1024); got != "64MB" {
		t.Fatalf("expected 64MB, got %q", got)
	}

	if got := formatBytes((uint64(21) * 1024 * 1024 * 1024) / 10); got != "2.1GB" {
		t.Fatalf("expected 2.1GB, got %q", got)
	}
}

type stubMetricsProvider struct {
	cpuUsage float64
	memory   MemoryMetrics
	storage  StorageMetrics
}

func (s stubMetricsProvider) CPUUsage() (float64, error) {
	return s.cpuUsage, nil
}

func (s stubMetricsProvider) MemoryUsage() (MemoryMetrics, error) {
	return s.memory, nil
}

func (s stubMetricsProvider) StorageUsage(path string) (StorageMetrics, error) {
	return s.storage, nil
}
