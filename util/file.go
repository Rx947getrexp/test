package util

import (
	"bufio"
	"go-speed/global"
	"os"
	"strings"
)

func ReadFile(fileName string) ([]string, error) {
	file, err := os.Open(fileName)
	if err != nil {
		global.Logger.Err(err).Msgf("os.Open(%s) failed, err: %s", fileName, err.Error())
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		line = strings.TrimSpace(line) // 去除空格和换行
		if line != "" {
			lines = append(lines, line)
		}
	}
	if err = scanner.Err(); err != nil {
		global.Logger.Err(err).Msgf("read (%s) failed, scanner.Err() is not nil, err: %s", fileName, err.Error())
		return nil, err
	}
	return lines, nil
}

func MustReadFile(fileName string) []string {
	lines, err := ReadFile(fileName)
	if err != nil {
		global.Logger.Fatal().Msgf("ReadFile(%s) failed, err: %s", fileName, err.Error())
	}
	return lines
}
