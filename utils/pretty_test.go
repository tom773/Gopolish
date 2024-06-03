package pretty

import (
	"bufio"
	"os"
	"testing"
)

func Test1(t *testing.T) {

	msg, err := os.OpenFile("../log1.txt", os.O_RDONLY, 0)
	defer msg.Close()
	if err != nil {
		t.Fatalf("error: %v", err)
	}

	inbuf := bufio.NewReader(msg)
	err = Pretty(inbuf, os.Stdout)
	if err != nil {
		t.Fatalf("error: %v", err)
	}
}

func Test2(t *testing.T) {

	msg, err := os.OpenFile("../log2.txt", os.O_RDONLY, 0)
	defer msg.Close()
	if err != nil {
		t.Fatalf("error: %v", err)
	}

	inbuf := bufio.NewReader(msg)
	err = Pretty(inbuf, os.Stdout)
	if err != nil {
		t.Fatalf("error: %v", err)
	}
}

func Test3(t *testing.T) {

	msg, err := os.OpenFile("../log3.txt", os.O_RDONLY, 0)
	defer msg.Close()
	if err != nil {
		t.Fatalf("error: %v", err)
	}

	inbuf := bufio.NewReader(msg)
	err = Pretty(inbuf, os.Stdout)
	if err != nil {
		t.Fatalf("error: %v", err)
	}
}
