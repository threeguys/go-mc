package nbt

import (
	"encoding/binary"
	"fmt"
)

/////////////////////////////////////////////////////////////////////////////////////////////////////
// IntTag
//
type IntTag struct {
	BaseTag
	value int32
}

func (tag *IntTag) GetType() int {
	return TAG_Int
}

func NewIntTag(name string) *IntTag {
	tag := new(IntTag)
	tag.BaseTag.name = name
	return tag
}

func (tag *IntTag) read(rdr *NBTReader) error {
	var data int32
	err := binary.Read(rdr.reader, binary.BigEndian, &data)
	if err != nil {
		return err
	}
	tag.value = data
	return nil
}

func (tag *IntTag) Print(indent string) {
	fmt.Printf("%sINT[%s] %d\n", indent, tag.BaseTag.name, tag.value)
}
