package srbuf
/* Simple Ring Buffer
 * http://github.com/coreyshuman/srbuf
 * (C) 2016 Corey Shuman
 * 4/26/16
 *
 * License: MIT
 */


type SimpleRingBuff struct {
    d       []byte  // data buffer
    size    int     // max buffer size, declared when instantiated
    rIdx    int     // read location index
    wIdx    int     // write location index
}

func Create(bufSize int) *SimpleRingBuff {
    b := &SimpleRingBuff {
        size: bufSize,
        rIdx: 0,
        wIdx: 0,
    }
    
    b.d = make([]byte, bufSize, bufSize)
    return b
}

func (b *SimpleRingBuff) PutByte(data byte) {
    b.d[b.wIdx] = data
    b.wIdx ++
    if(b.wIdx >= b.size) {
        b.wIdx = 0
    }
}

func (b *SimpleRingBuff) GetByte() byte {
    data := b.d[b.rIdx]
    b.rIdx ++
    if(b.rIdx >= b.size) {
        b.rIdx = 0
    }
    return data
}

/* Return Slice of all available bytes */
func (b *SimpleRingBuff) GetBytes() []byte {
    cnt := b.AvailByteCnt()
	if(b.rIdx + cnt < b.size) { // best case scenario, most efficient
		data := b.d[b.rIdx:b.rIdx+cnt]
		b.rIdx += cnt
		return data
	} else { // data wraps, we need to copy to a new buffer
		data := make([]byte, cnt)
		copy(data, b.d[b.rIdx:b.size])
		cnt -= (b.size - b.rIdx)
		copy(data[(b.size-b.rIdx):], b.d[0:cnt])
		b.rIdx = cnt
		return data
	}
}

func (b *SimpleRingBuff) AvailByteCnt() int {
    if(b.rIdx <= b.wIdx) {
        return b.wIdx - b.rIdx
    } else {
        return b.size - b.rIdx + b.wIdx
    }
}

func (b *SimpleRingBuff) Clear() {
    b.rIdx = 0
    b.wIdx = 0
}