package nbt

import (
	"encoding/binary"
	"fmt"
)

/////////////////////////////////////////////////////////////////////////////////////////////////////
// ShortTag
//
type ShortTag struct {
	BaseTag
	value int16
}

func (tag *ShortTag) GetType() int {
	return TAG_Short
}

func NewShortTag(name string) *ShortTag {
	tag := new(ShortTag)
	tag.BaseTag.name = name
	return tag
}

func (tag *ShortTag) read(rdr *NBTReader) error {
	var data int16
	err := binary.Read(rdr.reader, binary.BigEndian, &data)
	if err != nil {
		return err
	}
	tag.value = data
	return nil
}

func (tag *ShortTag) Print(indent string) {
	fmt.Printf("%sSHORT[%s] %d\n", indent, tag.BaseTag.name, tag.value)
}
