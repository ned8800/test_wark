package main

import (
	"bufio"
	"bytes"
	"io"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolutionSuccess(t *testing.T) {
	input := "2\n1 1\n1 1\n"
	expected := SUCCESS + "\n"

	r := bufio.NewReader(strings.NewReader(input))
	var out bytes.Buffer
	w := bufio.NewWriter(&out)

	err := Solution(r, w)
	assert.NoError(t, err)

	result := out.String()
	assert.Equal(t, expected, result)
}

func TestSolutionFail(t *testing.T) {
	input := "3\n10 20 30\n1 1 1\n 0 0 1\n"
	expected := FAIL + "\n"

	r := bufio.NewReader(strings.NewReader(input))
	var out bytes.Buffer
	w := bufio.NewWriter(&out)

	err := Solution(r, w)
	assert.NoError(t, err)

	result := out.String()
	assert.Equal(t, expected, result)
}

func TestSolutionInvalidN(t *testing.T) {
	input := "abc\n"
	r := bufio.NewReader(strings.NewReader(input))
	var out bytes.Buffer
	w := bufio.NewWriter(&out)

	err := Solution(r, w)
	assert.Error(t, err)
}

func TestSolutionOutOfRangeN(t *testing.T) {
	input := "0\n"
	r := bufio.NewReader(strings.NewReader(input))
	var out bytes.Buffer
	w := bufio.NewWriter(&out)

	err := Solution(r, w)
	assert.Error(t, err)

	assert.Contains(t, err.Error(), "n must be between")
}

func TestSolutionInvalidRowLength(t *testing.T) {
	input := "2\n1 1\n1\n"
	r := bufio.NewReader(strings.NewReader(input))
	var out bytes.Buffer
	w := bufio.NewWriter(&out)

	err := Solution(r, w)
	assert.Error(t, err)

	assert.Contains(t, err.Error(), "description has")
}

func TestSolutionInvalidBallCount(t *testing.T) {
	input := "2\n1 1\n1 x\n"
	r := bufio.NewReader(strings.NewReader(input))
	var out bytes.Buffer
	w := bufio.NewWriter(&out)

	err := Solution(r, w)
	assert.Error(t, err)

	assert.Contains(t, err.Error(), "Error converting ball count")
}

func TestSolutionBallCountOutOfRange(t *testing.T) {
	input := "2\n-1 1\n1 1\n"
	r := bufio.NewReader(strings.NewReader(input))
	var out bytes.Buffer
	w := bufio.NewWriter(&out)

	err := Solution(r, w)
	assert.Error(t, err)

	assert.Contains(t, err.Error(), "must be between")
}

func TestSolutionEmptyInput(t *testing.T) {
	input := ""
	r := bufio.NewReader(strings.NewReader(input))
	var out bytes.Buffer
	w := bufio.NewWriter(&out)

	err := Solution(r, w)
	assert.Error(t, err)

	assert.Contains(t, err.Error(), "Empty input")
}

func TestSolutionBrokenInput(t *testing.T) {
	pr, pw := io.Pipe()

	reader := bufio.NewReader(pr)
	var out bytes.Buffer
	writer := bufio.NewWriter(&out)

	pw.Close()
	pr.Close()

	err := Solution(reader, writer)
	assert.Error(t, err)

	assert.Contains(t, err.Error(), "Error reading")
}
