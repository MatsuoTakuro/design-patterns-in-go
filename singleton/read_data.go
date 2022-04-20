package singleton

import (
	"bufio"
	"os"
	"path"
	"runtime"
	"strconv"
)

func readData(filename string) (map[string]int, error) {
	relPath := getRelativePath()
	filePath := path.Join(relPath, filename)

	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	result := map[string]int{}

	for scanner.Scan() {
		k := scanner.Text()
		scanner.Scan()
		v, _ := strconv.Atoi(scanner.Text())
		result[k] = v
	}

	return result, nil
}

func getRelativePath() string {

	// Relative on runtime DIR:
	_, b, _, _ := runtime.Caller(0)
	relPath := path.Join(path.Dir(b))
	return relPath
}
