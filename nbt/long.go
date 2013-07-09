package nbt

import (
	"encoding/binary"
	"fmt"
)

/////////////////////////////////////////////////////////////////////////////////////////////////////
// LongTag
//
type LongTag struct {
	BaseTag
	value int64
}

func (tag *LongTag) GetType() int {
	return TAG_Long
}

func NewLongTag(name string) *LongTag {
	tag := new(LongTag)
	tag.BaseTag.name = name
	return tag
}

func (tag *LongTag) read(rdr *NBTReader) error {
	var data int64
	err := binary.Read(rdr.reader, binary.BigEndian, &data)
	if err != nil {
		return err
	}
	tag.value = data
	return nil
}

func (tag *LongTag) Print(indent string) {
	fmt.Printf("%sLONG[%s] %d\n", indent, tag.BaseTag.name, tag.value)
}
