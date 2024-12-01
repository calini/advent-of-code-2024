package main

import (
	"bufio"
	"log/slog"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Handler struct {
	Level slog.Level
}

// Add a simple DEBUG ENV var toggle for debug logs
func init() {
	debug, ok := os.LookupEnv("DEBUG")
	if ok && debug != "0" && debug != "false" && debug != "FALSE" {
		// TODO change just the Default logger Level once Go 1.22 drops
		slog.SetDefault(slog.New(slog.NewTextHandler(os.Stderr, &slog.HandlerOptions{
			Level: slog.LevelDebug,
		})))
	}
}

func main() {
	if len(os.Args) != 2 {
		slog.Error("Usage: main input_file")
		os.Exit(1)
	}

	// file opening
	file, err := os.Open(os.Args[1])
	defer file.Close()

	if err != nil {
		slog.Error("Error opening file: %s", err)
		os.Exit(2)
	}

	// file reading
	sc := bufio.NewScanner(file)

	// business logic

	var leftNums []int
	var rightNums []int
	for sc.Scan() {
		line := sc.Text()
		// slog.Info("read", "line", line)

		nums := strings.Split(line, "   ")
		leftNum, _ := strconv.Atoi(nums[0])
		rightNum, _ := strconv.Atoi(nums[1])
		leftNums = append(leftNums, leftNum)
		rightNums = append(rightNums, rightNum)
	}

	sort.Ints(leftNums)
	sort.Ints(rightNums)

	diffSum := 0
	for i := 0; i < len(leftNums); i++ {
		if leftNums[i] != rightNums[i] {
			diff := leftNums[i] - rightNums[i]
			if diff < 0 {
				diff = -diff
			}
			diffSum += diff
		}
	}

	// result
	slog.Info("Result", "diff_sum", diffSum)
}
