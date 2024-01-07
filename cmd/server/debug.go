package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"runtime"
	"runtime/debug"
	"strconv"
	"strings"

	"github.com/luisnquin/server-example/internal/log"
)

type MemoryStats struct {
	MemTotal int `json:"mem_total"`
	MemFree  int `json:"mem_free"`
	MemAvail int `json:"mem_avail"`
}

func showDebugInfo() error {
	buildInfo, _ := debug.ReadBuildInfo()

	stats, err := readMemoryStats()
	if err != nil {
		return fmt.Errorf("failed to read memory stats: %w", err)
	}

	cpuModel, err := getCpuModel()
	if err != nil {
		return fmt.Errorf("failed to get cpu model: %w", err)
	}

	log.Trace().Str("go_version", buildInfo.GoVersion).Str("GOOS", runtime.GOOS).Str("GOARCH", runtime.GOARCH).
		Int("cpu_count", runtime.NumCPU()).Str("cpu_model", cpuModel).Any("mem", stats).Int("pid", os.Getpid()).Send()

	return nil
}

func getCpuModel() (string, error) {
	f, err := os.OpenFile("/proc/cpuinfo", os.O_RDONLY, os.ModePerm)
	if err != nil {
		return "", err
	}

	defer f.Close()

	s := bufio.NewScanner(f)

	for s.Scan() {
		t := s.Text()

		if strings.HasPrefix(t, "model name") {
			_, after, _ := strings.Cut(t, ":")

			return strings.TrimSpace(after), nil
		}
	}

	return "", errors.New("cannot be found by name in /proc/cpuinfo")
}

func readMemoryStats() (MemoryStats, error) {
	f, err := os.OpenFile("/proc/meminfo", os.O_RDONLY, os.ModePerm)
	if err != nil {
		return MemoryStats{}, err
	}

	defer f.Close()

	var res MemoryStats

	s := bufio.NewScanner(f)

	c := 0

	for s.Scan() {
		if c == 3 {
			return res, nil
		}

		key, value := readMemoryLine(s.Text())
		switch key {
		case "MemTotal":
			res.MemTotal = value
			c++
		case "MemFree":
			res.MemFree = value
			c++
		case "MemAvailable":
			res.MemAvail = value
			c++
		}
	}
	return res, nil
}

func readMemoryLine(raw string) (key string, value int) {
	text := strings.ReplaceAll(raw[:len(raw)-2], " ", "")
	keyValue := strings.Split(text, ":")

	return keyValue[0], toInt(keyValue[1])
}

func toInt(raw string) int {
	if raw == "" {
		return 0
	}
	res, err := strconv.Atoi(raw)
	if err != nil {
		panic(err)
	}

	return res
}
