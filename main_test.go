package main

import (
	"bytes"
	"log"
	"strconv"
	"testing"
)

type Logger struct{}

var logContent bytes.Buffer

func (logger Logger) Write(p []byte) (n int, err error) {
	return logContent.Write(p)
}

func TestLogf(t *testing.T) {
	var logWriter Logger
	logger = log.New(logWriter, "", 0)

	verboseLogging = false
	logf("Testing")
	if logContent.String() != "" {
		t.Errorf("logf() should not write any content when verbose logging is disabled, but it wrote content")
	}

	verboseLogging = true
	logf("Testing")
	if logContent.String() != "Testing\n" {
		t.Errorf("logf() should write \"Testing\" when verbose logging is enabled, but it didn't")
	}
}

func TestVerboseLogging(t *testing.T) {
	verboseLogging = true
	setVerboseLogging([]string{})
	if verboseLogging {
		t.Errorf("setVerboseLogging() without a \"-v\" should disable verbose logging, but it didn't")
	}

	verboseLogging = false
	setVerboseLogging([]string{"-v"})
	if !verboseLogging {
		t.Errorf("setVerboseLogging() with a \"-v\" should enable verbose logging, but it didn't")
	}
}

func TestLengthArg(t *testing.T) {
	type lengthTest struct {
		Length         int
		ExpectedLength int
	}

	tests := []lengthTest{
		lengthTest{8192, 4096},
		lengthTest{1024, 1024},
		lengthTest{512, 512},
		lengthTest{128, 128},
		lengthTest{64, 64},
		lengthTest{32, 32},
		lengthTest{16, 16},
		lengthTest{8, 8},
		lengthTest{4, 8},
	}

	for _, test := range tests {
		arg := bytes.NewBufferString("--length=")
		arg.WriteString(strconv.Itoa(test.Length))

		length := getLength([]string{arg.String()})
		if length != test.ExpectedLength {
			t.Errorf("getLength() for %d length should return %d, but received %d", test.Length, test.ExpectedLength, length)
		}
	}
}

func TestSplitArg(t *testing.T) {
	type splitTest struct {
		Length           int
		Split            int
		NoSymbolSplit    int
		RepeatCharacters bool
	}

	tests := []splitTest{
		splitTest{1024, 341, 512, true},
		splitTest{512, 170, 256, true},
		splitTest{256, 85, 128, true},
		splitTest{128, 42, 64, true},
		splitTest{64, 21, 32, true},
		splitTest{32, 10, 16, true},
		splitTest{26, 8, 10, false},
		splitTest{16, 5, 8, false},
		splitTest{8, 2, 4, false},
	}

	for _, test := range tests {
		var split [2]int
		var allowRepeatCharacters bool

		split, allowRepeatCharacters = getSplit(test.Length, []string{})

		if split != [2]int{test.Split, test.Split} {
			t.Errorf("Split expected to be [%d, %d] for getSplit(%d) with symbols, but received %+v", test.Split, test.Split, test.Length, split)
		}

		if allowRepeatCharacters != test.RepeatCharacters {
			t.Errorf("Repeat characters expected to be %T for getSplit(%d) with symbols, but received %T", test.RepeatCharacters, test.Length, allowRepeatCharacters)
		}

		split, allowRepeatCharacters = getSplit(test.Length, []string{"--no-symbols"})

		if split != [2]int{test.NoSymbolSplit, 0} {
			t.Errorf("Split expected to be [%d, 0] for getSplit(%d) without symbols, but received %+v", test.NoSymbolSplit, test.Length, split)
		}

		if allowRepeatCharacters != test.RepeatCharacters {
			t.Errorf("Repeat characters expected to be %T for getSplit(%d) without symbols, but received %T", test.RepeatCharacters, test.Length, allowRepeatCharacters)
		}

	}
}
