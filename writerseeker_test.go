package writerseeker

import (
	"io"
	"testing"
)

func TestWrite(t *testing.T) {
	writerSeeker := &WriterSeeker{}
	var ws io.WriteSeeker = writerSeeker

	ws.Write([]byte("hello"))
	if string(writerSeeker.buf) != "hello" {
		t.Fail()
	}

	ws.Write([]byte(" world"))
	if string(writerSeeker.buf) != "hello world" {
		t.Fail()
	}

}

func TestSeek(t *testing.T) {
	writerSeeker := &WriterSeeker{}
	var ws io.WriteSeeker = writerSeeker

	ws.Write([]byte("hello"))
	if string(writerSeeker.buf) != "hello" {
		t.Fail()
	}

	ws.Write([]byte(" world"))
	if string(writerSeeker.buf) != "hello world" {
		t.Fail()
	}

	ws.Seek(-2, io.SeekEnd)
	ws.Write([]byte("k!"))
	if string(writerSeeker.buf) != "hello work!" {
		t.Fail()
	}

	ws.Seek(6, io.SeekStart)
	ws.Write([]byte("gopher"))
	if string(writerSeeker.buf) != "hello gopher" {
		t.Fail()
	}
}
