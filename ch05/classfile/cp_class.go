package classfile

type ConstantClassInfo struct {
	cp        ConstantPool
	nameIndex uint16
}

func (self *ConstantClassInfo) readInfo(reader *ClassReader) {
	self.nameIndex = reader.readUint16()
	// fmt.Print(self.nameIndex)
	// fmt.Printf("ConstantClassInfo.nameIndex :%v\n", self.nameIndex)
	// fmt.Printf("ConstantClassInfo.nameIndex :%v\n", self.nameIndex)

}
func (self *ConstantClassInfo) Name() string {
	return self.cp.getUtf8(self.nameIndex)
}
