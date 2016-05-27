package srbuf

import (
    "testing"
)

func TestReadWrite(t *testing.T) {
    rBuf := Create(10)
    rBuf.PutByte(byte('t'))
    rBuf.PutByte(byte('e'))
    rBuf.PutByte(byte('s'))
    rBuf.PutByte(byte('t'))
    if rBuf.AvailByteCnt() != 4 {
        t.Error("Expected 4, got ", rBuf.AvailByteCnt())
    }
    data := rBuf.GetByte()
    if data != byte('t') {
        t.Error("Expected 't', got ", string(data))
    }
    data = rBuf.GetByte()
    if data != byte('e') {
        t.Error("Expected 'e', got ", string(data))
    }
    data = rBuf.GetByte()
    if data != byte('s') {
        t.Error("Expected 's', got ", string(data))
    }
    data = rBuf.GetByte()
    if data != byte('t') {
        t.Error("Expected 't', got ", string(data))
    }
    if rBuf.AvailByteCnt() != 0 {
        t.Error("Expected 0, got ", rBuf.AvailByteCnt())
    }
}

func TestReadWrap(t *testing.T) {
	rBuf := Create(10)
    rBuf.PutByte(byte('1'))
    rBuf.PutByte(byte('2'))
    rBuf.PutByte(byte('3'))
    rBuf.PutByte(byte('4'))
	rBuf.PutByte(byte('5'))
	rBuf.PutByte(byte('6'))
	rBuf.PutByte(byte('7'))
	rBuf.PutByte(byte('8'))
	rBuf.GetBytes()
	rBuf.PutByte(byte('a'))
	rBuf.PutByte(byte('b'))
	rBuf.PutByte(byte('c'))
	rBuf.PutByte(byte('d'))
	if rBuf.AvailByteCnt() != 4 {
        t.Error("Expected 4, got ", rBuf.AvailByteCnt())
    }
	d := rBuf.GetByte()
    if d != byte('a') {
        t.Error("Expected 'a', got ", string(d))
    }
    d = rBuf.GetByte()
    if d != byte('b') {
        t.Error("Expected 'b', got ", string(d))
    }
    d = rBuf.GetByte()
    if d != byte('c') {
        t.Error("Expected 'c', got ", string(d))
    }
    d = rBuf.GetByte()
    if d != byte('d') {
        t.Error("Expected 'd', got ", string(d))
    }
    if rBuf.AvailByteCnt() != 0 {
        t.Error("Expected 0, got ", rBuf.AvailByteCnt())
    }
}

func TestReadBytes(t *testing.T) {
	rBuf := Create(10)
    rBuf.PutByte(byte('t'))
    rBuf.PutByte(byte('e'))
    rBuf.PutByte(byte('s'))
    rBuf.PutByte(byte('t'))
	if rBuf.AvailByteCnt() != 4 {
        t.Error("Expected 4, got ", rBuf.AvailByteCnt())
    }
	data := rBuf.GetBytes()
	if string(data) != "test" {
		t.Error("Expected 'test', got ", string(data))
	}
	if rBuf.AvailByteCnt() != 0 {
        t.Error("Expected 0, got ", rBuf.AvailByteCnt())
    }
}

func TestReadBytesWrap(t *testing.T) {
	rBuf := Create(10)
    rBuf.PutByte(byte('1'))
    rBuf.PutByte(byte('2'))
    rBuf.PutByte(byte('3'))
    rBuf.PutByte(byte('4'))
	rBuf.PutByte(byte('5'))
	rBuf.PutByte(byte('6'))
	rBuf.PutByte(byte('7'))
	rBuf.PutByte(byte('8'))
	data := rBuf.GetBytes()
	rBuf.PutByte(byte('a'))
	rBuf.PutByte(byte('b'))
	rBuf.PutByte(byte('c'))
	rBuf.PutByte(byte('d'))
	if rBuf.AvailByteCnt() != 4 {
        t.Error("Expected 4, got ", rBuf.AvailByteCnt())
    }
	data = rBuf.GetBytes()
	if string(data) != "abcd" {
		t.Error("Expected 'abcd', got ", string(data))
	}
	if rBuf.AvailByteCnt() != 0 {
        t.Error("Expected 0, got ", rBuf.AvailByteCnt())
    }
}