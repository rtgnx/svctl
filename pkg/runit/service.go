package runit

import (
	"io"
	"os"
	"path"
	"path/filepath"
	"syscall"
	"time"
)

const (
	taiOffset = 4611686018427387914
	statusLen = 20

	posTimeStart = 0
	posTimeEnd   = 7
	posPidStart  = 12
	posPidEnd    = 15

	posWant  = 17
	posState = 19
)

const (
	superviseDir = "supervise"
	statusPath   = "supervise/status"
)

func ListServiceNames(root string) ([]string, error) {
	names := []string{}
	dirEntries, err := os.ReadDir(root)

	if err != nil {
		return names, err
	}

	for _, e := range dirEntries {
		names = append(names, path.Base(e.Name()))
	}
	return names, nil
}

func ReadService(root, name string) (Service, error) {
	statusFilePath := filepath.Join(root, name, statusPath)
	fd, err := os.Open(statusFilePath)
	if err != nil {
		return Service{}, err
	}

	defer fd.Close()

	b, err := io.ReadAll(fd)
	if err != nil {
		return Service{}, err
	}

	return parseStatus(name, b)
}

func parseStatus(name string, data []byte) (Service, error) {
	svc := Service{}

	var pid int
	pid = int(data[posPidEnd])
	for i := posPidEnd - 1; i >= posPidStart; i-- {
		pid <<= 8
		pid += int(data[i])
	}

	svc.Pid = pid

	tai := int64(data[posTimeStart])
	for i := posTimeStart + 1; i <= posTimeEnd; i++ {
		tai <<= 8
		tai += int64(data[i])
	}
	state := data[posState] // 0: down, 1: run, 2: finish
	svc.Timestamp = time.Unix(tai-taiOffset, 0)
	svc.State = int(state)
	tv := &syscall.Timeval{}
	if err := syscall.Gettimeofday(tv); err != nil {
		return svc, err
	}

	svc.Duration = time.Duration(int(int64(tv.Sec) - (tai - taiOffset)))

	switch data[posWant] {
	case 'u':
		svc.Want = StateUp

	case 'd':
		svc.Want = StateDown
	}

	svc.Name = name

	return svc, nil
}
