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