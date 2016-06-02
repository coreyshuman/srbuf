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
	rBuf.GetBytes(0)
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
	data := rBuf.GetBytes(0)
	if string(data) != "test" {
		t.Error("Expected 'test', got ", string(data))
	}
	if rBuf.AvailByteCnt() != 0 {
        t.Error("Expected 0, got ", rBuf.AvailByteCnt())
    }
}

func TestReadBytesN(t *testing.T) {
	rBuf := Create(10)
    rBuf.PutByte(byte('t'))
    rBuf.PutByte(byte('e'))
    rBuf.PutByte(byte('s'))
    rBuf.PutByte(byte('t'))
	rBuf.PutByte(byte('1'))
    rBuf.PutByte(byte('2'))
    rBuf.PutByte(byte('3'))
    rBuf.PutByte(byte('4'))
	if rBuf.AvailByteCnt() != 8 {
        t.Error("Expected 8, got ", rBuf.AvailByteCnt())
    }
	data := rBuf.GetBytes(5)
	if string(data) != "test1" {
		t.Error("Expected 'test1', got ", string(data))
	}
	if rBuf.AvailByteCnt() != 3 {
        t.Error("Expected 3, got ", rBuf.AvailByteCnt())
    }
	data = rBuf.GetBytes(2)
	if string(data) != "23" {
		t.Error("Expected '23', got ", string(data))
	}
	if rBuf.AvailByteCnt() != 1 {
        t.Error("Expected 1, got ", rBuf.AvailByteCnt())
    }
}

func TestPeekBytes(t *testing.T) {
	rBuf := Create(10)
    rBuf.PutByte(byte('t'))
    rBuf.PutByte(byte('e'))
    rBuf.PutByte(byte('s'))
    rBuf.PutByte(byte('t'))
	if rBuf.AvailByteCnt() != 4 {
        t.Error("Expected 4, got ", rBuf.AvailByteCnt())
    }
	data := rBuf.PeekBytes(2)
	if string(data) != "te" {
		t.Error("Expected 'te', got ", string(data))
	}
	if rBuf.AvailByteCnt() != 4 {
        t.Error("Expected 4, got ", rBuf.AvailByteCnt())
    }
	data = rBuf.PeekBytes(0)
	if string(data) != "test" {
		t.Error("Expected 'test', got ", string(data))
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
	data := rBuf.GetBytes(0)
	rBuf.PutByte(byte('a'))
	rBuf.PutByte(byte('b'))
	rBuf.PutByte(byte('c'))
	rBuf.PutByte(byte('d'))
	if rBuf.AvailByteCnt() != 4 {
        t.Error("Expected 4, got ", rBuf.AvailByteCnt())
    }
	data = rBuf.GetBytes(0)
	if string(data) != "abcd" {
		t.Error("Expected 'abcd', got ", string(data))
	}
	if rBuf.AvailByteCnt() != 0 {
        t.Error("Expected 0, got ", rBuf.AvailByteCnt())
    }
}