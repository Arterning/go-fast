package base

/**
字节码阅读器
*/
type BytecodeReader struct {
	code []byte // bytecodes 存放字节码
	pc   int    //记录读取到了哪个字节
}

func (self *BytecodeReader) Reset(code []byte, pc int) {
	self.code = code
	self.pc = pc
}

func (self *BytecodeReader) PC() int {
	return self.pc
}

func (self *BytecodeReader) ReadInt8() int8 {
	return int8(self.ReadUint8())
}

//从字节码中读取一个字节的数据，同时pc+1
func (self *BytecodeReader) ReadUint8() uint8 {
	i := self.code[self.pc]
	self.pc++
	return i
}

func (self *BytecodeReader) ReadInt16() int16 {
	return int16(self.ReadUint16())
}
func (self *BytecodeReader) ReadUint16() uint16 {
	byte1 := uint16(self.ReadUint8())
	byte2 := uint16(self.ReadUint8())
	return (byte1 << 8) | byte2
}

func (self *BytecodeReader) ReadInt32() int32 {
	byte1 := int32(self.ReadUint8())
	byte2 := int32(self.ReadUint8())
	byte3 := int32(self.ReadUint8())
	byte4 := int32(self.ReadUint8())
	return (byte1 << 24) | (byte2 << 16) | (byte3 << 8) | byte4
}

// used by lookupswitch and tableswitch
func (self *BytecodeReader) ReadInt32s(n int32) []int32 {
	ints := make([]int32, n)
	for i := range ints {
		ints[i] = self.ReadInt32()
	}
	return ints
}

// used by lookupswitch and tableswitch
func (self *BytecodeReader) SkipPadding() {
	for self.pc%4 != 0 {
		self.ReadUint8()
	}
}
