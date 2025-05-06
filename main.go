package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"

	"github.com/pkg/errors"
)

const (
	MAX_BALL_COUNT = 1000000000
	MIN_BALL_COUNT = 0

	MAX_N = 100
	MIN_N = 1

	SUCCESS = "yes"
	FAIL    = "no"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)

	err := Solution(reader, writer)
	if err != nil {
		log.Fatal(err)
	}
}

func Solution(reader *bufio.Reader, writer *bufio.Writer) error {
	nStr, err := reader.ReadString('\n')
	if err != nil {
		if err == io.EOF && strings.TrimSpace(nStr) == "" {
			errMsg := errors.Wrap(err, "Empty input")
			return errMsg
		}
		errMsg := errors.Wrap(err, "Error reading n")
		return errMsg
	}

	nStr = strings.TrimSpace(nStr)
	n, err := strconv.Atoi(nStr)
	if err != nil {
		errMsg := errors.Wrap(err, "Error converting n to integer")
		return errMsg
	}

	if n < MIN_N || n > MAX_N {
		errMsg := errors.New(fmt.Sprintf("n must be between %d and %d, got %d ", MIN_N, MAX_N, n))
		return errMsg
	}

	ballsInContainer := make([]int, n)
	ballsOfColor := make([]int, n)

	for i := 0; i < n; i++ {
		row, err := reader.ReadString('\n')
		if err != nil {
			errMsg := errors.Wrapf(err, "Error reading container %d data ", i)
			return errMsg
		}
		row = strings.TrimSpace(row)
		rowParts := strings.Fields(row)

		if len(rowParts) != n {
			errMsg := errors.New(fmt.Sprintf("Error happened: Container %d description has %d numbers, expected %d",
				i,
				len(rowParts),
				n,
			))
			return errMsg
		}

		currentContainerSum := 0
		for j := 0; j < n; j++ {
			count, err := strconv.Atoi(rowParts[j])
			if err != nil {
				errMsg := errors.Wrapf(err, "Error converting ball count for container %d, color %d", i, j)
				return errMsg
			}

			if count < MIN_BALL_COUNT || count > MAX_BALL_COUNT {
				errMsg := errors.New(
					fmt.Sprintf("Error checking ball count for container %d, color %d must be between %d and %d, got %d\n",
						i,
						j,
						MIN_BALL_COUNT,
						MAX_BALL_COUNT,
						count),
				)
				return errMsg
			}

			ballsOfColor[j] += count
			currentContainerSum += count
		}
		ballsInContainer[i] = currentContainerSum
	}

	sort.Ints(ballsInContainer)
	sort.Ints(ballsOfColor)

	possible := true
	for i := 0; i < n; i++ {
		if ballsInContainer[i] != ballsOfColor[i] {
			possible = false
			break
		}
	}

	if possible {
		fmt.Fprintln(writer, SUCCESS)
	} else {
		fmt.Fprintln(writer, FAIL)
	}

	writer.Flush()
	return nil
}
