package system

import (
	"fmt"
	"strings"
	"time"
)

type HealthReporter interface {
	Snapshot() HealthSnapshot
}

type MetricsProvider interface {
	CPUUsage() (float64, error)
	MemoryUsage() (MemoryMetrics, error)
	StorageUsage(path string) (StorageMetrics, error)
}

type HealthSnapshot struct {
	Status       string        `json:"status"`
	SystemHealth SystemHealth  `json:"system_health"`
	Alerts       []HealthAlert `json:"alerts,omitempty"`
}

type SystemHealth struct {
	CPUUsage string        `json:"cpu_usage"`
	Memory   ResourceUsage `json:"memory"`
	Storage  StorageUsage  `json:"storage"`
	Uptime   string        `json:"uptime"`
}

type HealthAlert struct {
	Level    string `json:"level"`
	Resource string `json:"resource"`
	Message  string `json:"message"`
}

type ResourceUsage struct {
	Used       string `json:"used"`
	Total      string `json:"total"`
	Percentage string `json:"percentage"`
}

type StorageUsage struct {
	Used       string `json:"used"`
	Free       string `json:"free"`
	Total      string `json:"total"`
	Percentage string `json:"percentage"`
}

type MemoryMetrics struct {
	UsedBytes  uint64
	TotalBytes uint64
}

type StorageMetrics struct {
	UsedBytes  uint64
	FreeBytes  uint64
	TotalBytes uint64
}

type HealthService struct {
	startedAt       time.Time
	storagePath     string
	metricsProvider MetricsProvider
	thresholds      HealthThresholds
}

func NewHealthService(startedAt time.Time, storagePath string, metricsProvider MetricsProvider) *HealthService {
	return NewHealthServiceWithThresholds(startedAt, storagePath, metricsProvider, DefaultHealthThresholds())
}

func NewHealthServiceWithThresholds(startedAt time.Time, storagePath string, metricsProvider MetricsProvider, thresholds HealthThresholds) *HealthService {
	if storagePath == "" {
		storagePath = "."
	}

	return &HealthService{
		startedAt:       startedAt,
		storagePath:     storagePath,
		metricsProvider: metricsProvider,
		thresholds:      thresholds,
	}
}

func (s *HealthService) Snapshot() HealthSnapshot {
	cpuUsage, err := s.metricsProvider.CPUUsage()
	if err != nil {
		cpuUsage = 0
	}

	memoryMetrics, err := s.metricsProvider.MemoryUsage()
	if err != nil {
		memoryMetrics = MemoryMetrics{}
	}

	storageMetrics, err := s.metricsProvider.StorageUsage(s.storagePath)
	if err != nil {
		storageMetrics = StorageMetrics{}
	}

	alerts := s.evaluateAlerts(cpuUsage, memoryMetrics, storageMetrics)

	return HealthSnapshot{
		Status: statusFromAlerts(alerts),
		SystemHealth: SystemHealth{
			CPUUsage: formatPercentage(cpuUsage),
			Memory:   formatMemoryUsage(memoryMetrics),
			Storage:  formatStorageUsage(storageMetrics),
			Uptime:   formatUptime(time.Since(s.startedAt)),
		},
		Alerts: alerts,
	}
}

func formatMemoryUsage(metrics MemoryMetrics) ResourceUsage {
	return ResourceUsage{
		Used:       formatBytes(metrics.UsedBytes),
		Total:      formatBytes(metrics.TotalBytes),
		Percentage: formatPercentage(safePercentage(float64(metrics.UsedBytes), float64(metrics.TotalBytes))),
	}
}

func formatStorageUsage(metrics StorageMetrics) StorageUsage {
	return StorageUsage{
		Used:       formatBytes(metrics.UsedBytes),
		Free:       formatBytes(metrics.FreeBytes),
		Total:      formatBytes(metrics.TotalBytes),
		Percentage: formatPercentage(safePercentage(float64(metrics.UsedBytes), float64(metrics.TotalBytes))),
	}
}

func safePercentage(value float64, total float64) float64 {
	if total <= 0 {
		return 0
	}

	return (value / total) * 100
}

func formatBytes(value uint64) string {
	units := []string{"B", "KB", "MB", "GB", "TB"}
	formatted := float64(value)
	unitIndex := 0

	for formatted >= 1024 && unitIndex < len(units)-1 {
		formatted /= 1024
		unitIndex++
	}

	return trimTrailingZeros(formatted) + units[unitIndex]
}

func formatPercentage(value float64) string {
	return trimTrailingZeros(value) + "%"
}

func formatUptime(duration time.Duration) string {
	if duration < 0 {
		duration = 0
	}

	totalMinutes := int(duration.Minutes())
	days := totalMinutes / (24 * 60)
	hours := (totalMinutes % (24 * 60)) / 60
	minutes := totalMinutes % 60

	parts := make([]string, 0, 3)
	if days > 0 {
		parts = append(parts, fmt.Sprintf("%dd", days))
	}
	if hours > 0 || days > 0 {
		parts = append(parts, fmt.Sprintf("%dh", hours))
	}
	parts = append(parts, fmt.Sprintf("%dm", minutes))

	return strings.Join(parts, " ")
}

func trimTrailingZeros(value float64) string {
	text := fmt.Sprintf("%.1f", value)
	return strings.TrimSuffix(strings.TrimSuffix(text, "0"), ".")
}
