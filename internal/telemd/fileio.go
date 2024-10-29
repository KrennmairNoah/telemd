package telemd

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// visitLines processes the given file line by line and applies a visitor
// function to each line. If the visitor returns false, then the method returns
// and the file is closed.
func visitLines(path string, visitor func(string) bool) error {
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		cont := visitor(scanner.Text())
		if !cont {
			break
		}
	}

	return scanner.Err()
}

func parseInt64(str string) (int64, error) {
	return strconv.ParseInt(str, 10, 64)
}

func parseInt64Array(arr []string) ([]int64, error) {
	ints := make([]int64, len(arr))
	var err error = nil

	for i := 0; i < len(arr); i++ {
		ints[i], err = parseInt64(arr[i])
		if err != nil {
			return ints, err
		}
	}

	return ints, err
}

func readSpecificLine(path string, lineNumber int) (string, error) {
	file, err := os.Open(path)
	if err != nil {
		return "", err
	}
	defer func() {
		_ = file.Close()
	}()

	scanner := bufio.NewScanner(file)

	currentLine := 0
	for scanner.Scan() {
		if currentLine == lineNumber {
			return scanner.Text(), nil
		}
		currentLine++
	}

	if err := scanner.Err(); err != nil {
		return "", err
	}

	return "", fmt.Errorf("line %d not found in file", lineNumber)
}

func readLineAndParseInt(path string, lineNumber int) (int64, error) {
	line, err := readSpecificLine(path, lineNumber)
	if err != nil {
		return -1, err
	}
	return strconv.ParseInt(line, 10, 64)
}

func fileDirExists(filename string) bool {
	_, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return true
}
