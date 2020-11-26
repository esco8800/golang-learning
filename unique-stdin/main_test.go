package main

import (
	"bufio"
	"bytes"
	"strings"
	"testing"
)

var testOk = `1
2
3
4
4
5`

var testResult = `1
2
3
4
5
`

func TestOk(t *testing.T) {

	in := bufio.NewReader(strings.NewReader(testOk))
	out := new(bytes.Buffer)
	err := uniq(in, out)

	if err != nil {
		t.Errorf("TestOk is fail")
	}
	result := out.String()
	if testResult != result {
		t.Errorf("TestOk is fail. Result not match /n %v %v", result, testResult)
	}
}

var testFail = `1
2
1`

func TestForError(t *testing.T)  {
	in := bufio.NewReader(strings.NewReader(testFail))
	out := new(bytes.Buffer)
	err := uniq(in, out)

	if err == nil {
		t.Errorf("TestForError is fail")
	}
}
