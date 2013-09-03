package dbitstream

import (
	"bytes"
	"io"
	"strings"
	"testing"
)

func TestBitStreamEOF(t *testing.T) {

	br := NewReader(strings.NewReader("0"))

	b, err := br.ReadByte()
	if b != '0' {
		t.Error("ReadByte didn't return first byte")
	}

	b, err = br.ReadByte()
	if err != io.EOF {
		t.Error("ReadByte on empty string didn't return EOF")
	}

	// 0 = 0b00110000
	br = NewReader(strings.NewReader("0"))

	buf := bytes.NewBuffer(nil)
	bw := NewWriter(buf)

	for i := 0; i < 4; i++ {
		bit, err := br.ReadBit()
		if err == io.EOF {
			break
		}
		if err != nil {
			t.Error("GetBit returned error err=", err.Error())
			return
		}
		bw.WriteBit(bit)
	}

	bw.Flush(One)

	err = bw.WriteByte(0xAA)
	if err != nil {
		t.Error("unable to WriteByte")
	}

	c := buf.Bytes()

	if len(c) != 2 || c[1] != 0xAA || c[0] != 0x3f {
		t.Error("bad return from 4 read bytes")
	}

	_, err = NewReader(strings.NewReader("")).ReadBit()
	if err != io.EOF {
		t.Error("ReadBit on empty string didn't return EOF")
	}

}

func TestBitStream(t *testing.T) {

	buf := bytes.NewBuffer(nil)
	br := NewReader(strings.NewReader("hello"))
	bw := NewWriter(buf)

	for {
		bit, err := br.ReadBit()
		if err == io.EOF {
			break
		}
		if err != nil {
			t.Error("GetBit returned error err=", err.Error())
			return
		}
		bw.WriteBit(bit)
	}

	s := buf.String()

	if s != "hello" {
		t.Error("expected 'hello', got=", []byte(s))
	}
}

func TestByteStream(t *testing.T) {

	buf := bytes.NewBuffer(nil)
	br := NewReader(strings.NewReader("hello"))
	bw := NewWriter(buf)

	for i := 0; i < 3; i++ {
		bit, err := br.ReadBit()
		if err == io.EOF {
			break
		}
		if err != nil {
			t.Error("GetBit returned error err=", err.Error())
			return
		}
		bw.WriteBit(bit)
	}

	for i := 0; i < 4; i++ {
		byt, err := br.ReadByte()
		if err == io.EOF {
			break
		}
		if err != nil {
			t.Error("GetByte returned error err=", err.Error())
			return
		}
		bw.WriteByte(byt)
	}

	for i := 0; i < 5; i++ {
		bit, err := br.ReadBit()
		if err == io.EOF {
			break
		}
		if err != nil {
			t.Error("GetBit returned error err=", err.Error())
			return
		}
		bw.WriteBit(bit)
	}

	s := buf.String()

	if s != "hello" {
		t.Error("expected 'hello', got=", []byte(s))
	}
}