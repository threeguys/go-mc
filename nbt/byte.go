package nbt

import (
	"fmt"
)

/////////////////////////////////////////////////////////////////////////////////////////////////////
// ByteTag
//
type ByteTag struct {
	BaseTag
	value byte
}

func (tag *ByteTag) GetType() int {
	return TAG_Byte
}

func NewByteTag(name string) *ByteTag {
	tag := new(ByteTag)
	tag.BaseTag.name = name
	return tag
}

func (tag *ByteTag) read(rdr *NBTReader) error {
	c, e := rdr.reader.ReadByte()
	if e != nil {
		return e
	}
	tag.value = c
	return nil
}

func (tag *ByteTag) Print(indent string) {
	fmt.Printf("%sBYTE[%s] %d\n", indent, tag.BaseTag.name, tag.value)
}
