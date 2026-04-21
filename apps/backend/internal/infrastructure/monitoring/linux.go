package monitoring

import (
	"bufio"
	"errors"
	"math"
	"os"
	"runtime"
	"strconv"
	"strings"
	"sync"
	"syscall"

	systemusecase "api/internal/usecase/system"
)

type LinuxMetricsProvider struct {
	mu        sync.Mutex
	lastTotal uint64
	lastIdle  uint64
	hasSample bool
}

func NewLinuxMetricsProvider() *LinuxMetricsProvider {
	return &LinuxMetricsProvider{}
}

func (p *LinuxMetricsProvider) CPUUsage() (float64, error) {
	total, idle, err := readCPUStat("/proc/stat")
	if err != nil {
		return 0, err
	}

	p.mu.Lock()
	defer p.mu.Unlock()

	if !p.hasSample {
		p.lastTotal = total
		p.lastIdle = idle
		p.hasSample = true
		return fallbackCPUUsage()
	}

	totalDelta := total - p.lastTotal
	idleDelta := idle - p.lastIdle

	p.lastTotal = total
	p.lastIdle = idle

	if totalDelta == 0 {
		return 0, nil
	}

	usage := (1 - float64(idleDelta)/float64(totalDelta)) * 100
	if usage < 0 {
		return 0, nil
	}

	return math.Min(usage, 100), nil
}

func (p *LinuxMetricsProvider) MemoryUsage() (systemusecase.MemoryMetrics, error) {
	total, available, err := readMemInfo("/proc/meminfo")
	if err != nil {
		return systemusecase.MemoryMetrics{}, err
	}

	used := uint64(0)
	if total > available {
		used = total - available
	}

	return systemusecase.MemoryMetrics{
		UsedBytes:  used,
		TotalBytes: total,
	}, nil
}

func (p *LinuxMetricsProvider) StorageUsage(path string) (systemusecase.StorageMetrics, error) {
	var stat syscall.Statfs_t
	if err := syscall.Statfs(path, &stat); err != nil {
		return systemusecase.StorageMetrics{}, err
	}

	total := stat.Blocks * uint64(stat.Bsize)
	free := stat.Bavail * uint64(stat.Bsize)
	used := total - free

	return systemusecase.StorageMetrics{
		UsedBytes:  used,
		FreeBytes:  free,
		TotalBytes: total,
	}, nil
}

func fallbackCPUUsage() (float64, error) {
	data, err := os.ReadFile("/proc/loadavg")
	if err != nil {
		return 0, err
	}

	fields := strings.Fields(string(data))
	if len(fields) == 0 {
		return 0, errors.New("loadavg is empty")
	}

	loadAverage, err := strconv.ParseFloat(fields[0], 64)
	if err != nil {
		return 0, err
	}

	cpuCount := float64(runtime.NumCPU())
	if cpuCount == 0 {
		return 0, errors.New("cpu count is zero")
	}

	usage := (loadAverage / cpuCount) * 100
	if usage < 0 {
		return 0, nil
	}

	return math.Min(usage, 100), nil
}

func readCPUStat(path string) (uint64, uint64, error) {
	file, err := os.Open(path)
	if err != nil {
		return 0, 0, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if !strings.HasPrefix(line, "cpu ") {
			continue
		}

		fields := strings.Fields(line)
		if len(fields) < 5 {
			return 0, 0, errors.New("invalid cpu stat format")
		}

		values := make([]uint64, 0, len(fields)-1)
		for _, field := range fields[1:] {
			value, parseErr := strconv.ParseUint(field, 10, 64)
			if parseErr != nil {
				return 0, 0, parseErr
			}
			values = append(values, value)
		}

		var total uint64
		for _, value := range values {
			total += value
		}

		idle := values[3]
		if len(values) > 4 {
			idle += values[4]
		}

		return total, idle, nil
	}

	if err := scanner.Err(); err != nil {
		return 0, 0, err
	}

	return 0, 0, errors.New("cpu stat line not found")
}

func readMemInfo(path string) (uint64, uint64, error) {
	file, err := os.Open(path)
	if err != nil {
		return 0, 0, err
	}
	defer file.Close()

	var total uint64
	var available uint64

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, "MemTotal:") {
			total, err = parseMemInfoLine(line)
			if err != nil {
				return 0, 0, err
			}
		}

		if strings.HasPrefix(line, "MemAvailable:") {
			available, err = parseMemInfoLine(line)
			if err != nil {
				return 0, 0, err
			}
		}
	}

	if err := scanner.Err(); err != nil {
		return 0, 0, err
	}

	if total == 0 {
		return 0, 0, errors.New("memtotal not found")
	}

	return total, available, nil
}

func parseMemInfoLine(line string) (uint64, error) {
	fields := strings.Fields(line)
	if len(fields) < 2 {
		return 0, errors.New("invalid meminfo line")
	}

	valueKB, err := strconv.ParseUint(fields[1], 10, 64)
	if err != nil {
		return 0, err
	}

	return valueKB * 1024, nil
}
