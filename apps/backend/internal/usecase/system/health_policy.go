package system

type HealthThresholds struct {
	CPUWarningPercentage      float64
	CPUCriticalPercentage     float64
	MemoryWarningPercentage   float64
	MemoryCriticalPercentage  float64
	StorageWarningPercentage  float64
	StorageCriticalPercentage float64
	StorageWarningFreeBytes   uint64
	StorageCriticalFreeBytes  uint64
}

func DefaultHealthThresholds() HealthThresholds {
	return HealthThresholds{
		CPUWarningPercentage:      85,
		CPUCriticalPercentage:     95,
		MemoryWarningPercentage:   85,
		MemoryCriticalPercentage:  95,
		StorageWarningPercentage:  85,
		StorageCriticalPercentage: 95,
		StorageWarningFreeBytes:   2 * 1024 * 1024 * 1024,
		StorageCriticalFreeBytes:  512 * 1024 * 1024,
	}
}

func (s *HealthService) evaluateAlerts(cpuUsage float64, memoryMetrics MemoryMetrics, storageMetrics StorageMetrics) []HealthAlert {
	alerts := make([]HealthAlert, 0, 3)

	if alert, ok := evaluateCPUAlert(cpuUsage, s.thresholds); ok {
		alerts = append(alerts, alert)
	}

	if alert, ok := evaluateMemoryAlert(memoryMetrics, s.thresholds); ok {
		alerts = append(alerts, alert)
	}

	if alert, ok := evaluateStorageAlert(s.storagePath, storageMetrics, s.thresholds); ok {
		alerts = append(alerts, alert)
	}

	return alerts
}

func evaluateCPUAlert(usage float64, thresholds HealthThresholds) (HealthAlert, bool) {
	switch {
	case usage >= thresholds.CPUCriticalPercentage:
		return HealthAlert{
			Level:    "critical",
			Resource: "cpu",
			Message:  "CPU usage is critically high at " + formatPercentage(usage),
		}, true
	case usage >= thresholds.CPUWarningPercentage:
		return HealthAlert{
			Level:    "warning",
			Resource: "cpu",
			Message:  "CPU usage is high at " + formatPercentage(usage),
		}, true
	default:
		return HealthAlert{}, false
	}
}

func evaluateMemoryAlert(metrics MemoryMetrics, thresholds HealthThresholds) (HealthAlert, bool) {
	usage := safePercentage(float64(metrics.UsedBytes), float64(metrics.TotalBytes))

	switch {
	case usage >= thresholds.MemoryCriticalPercentage:
		return HealthAlert{
			Level:    "critical",
			Resource: "memory",
			Message:  "Memory usage is critically high at " + formatPercentage(usage) + " (" + formatBytes(metrics.UsedBytes) + "/" + formatBytes(metrics.TotalBytes) + ")",
		}, true
	case usage >= thresholds.MemoryWarningPercentage:
		return HealthAlert{
			Level:    "warning",
			Resource: "memory",
			Message:  "Memory usage is high at " + formatPercentage(usage) + " (" + formatBytes(metrics.UsedBytes) + "/" + formatBytes(metrics.TotalBytes) + ")",
		}, true
	default:
		return HealthAlert{}, false
	}
}

func evaluateStorageAlert(path string, metrics StorageMetrics, thresholds HealthThresholds) (HealthAlert, bool) {
	usage := safePercentage(float64(metrics.UsedBytes), float64(metrics.TotalBytes))

	switch {
	case metrics.FreeBytes > 0 && metrics.FreeBytes <= thresholds.StorageCriticalFreeBytes:
		return HealthAlert{
			Level:    "critical",
			Resource: "storage",
			Message:  "Storage is critically low at " + path + ": " + formatBytes(metrics.FreeBytes) + " free remaining",
		}, true
	case usage >= thresholds.StorageCriticalPercentage:
		return HealthAlert{
			Level:    "critical",
			Resource: "storage",
			Message:  "Storage usage is critically high at " + formatPercentage(usage) + " on " + path,
		}, true
	case metrics.FreeBytes > 0 && metrics.FreeBytes <= thresholds.StorageWarningFreeBytes:
		return HealthAlert{
			Level:    "warning",
			Resource: "storage",
			Message:  "Storage free space is low at " + path + ": " + formatBytes(metrics.FreeBytes) + " remaining",
		}, true
	case usage >= thresholds.StorageWarningPercentage:
		return HealthAlert{
			Level:    "warning",
			Resource: "storage",
			Message:  "Storage usage is high at " + formatPercentage(usage) + " on " + path,
		}, true
	default:
		return HealthAlert{}, false
	}
}

func statusFromAlerts(alerts []HealthAlert) string {
	status := "operational"

	for _, alert := range alerts {
		if alert.Level == "critical" {
			return "critical"
		}

		if alert.Level == "warning" {
			status = "degraded"
		}
	}

	return status
}
