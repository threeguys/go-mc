package nbt

import (
	"encoding/binary"
	"fmt"
)

/////////////////////////////////////////////////////////////////////////////////////////////////////
// StringTag
//
type StringTag struct {
	BaseTag
	value string
}

func (tag *StringTag) GetType() int {
	return TAG_String
}

func NewStringTag(name string) *StringTag {
	tag := new(StringTag)
	tag.BaseTag.name = name
	return tag
}

func (tag *StringTag) read(rdr *NBTReader) error {
	var len int16
	err := binary.Read(rdr.reader, binary.BigEndian, &len)
	if err != nil {
		return err
	}
	if len < 0 {
		return NBTError{fmt.Sprintf("Read length of less than zero for string!")}
	}
	
	data, err := readString(rdr.reader, int(len))
	if err != nil {
		return err
	}
	tag.value = data
	return nil
}

func (tag *StringTag) Print(indent string) {
	fmt.Printf("%sSTRING[%s] \"%s\"\n", indent, tag.BaseTag.name, tag.value)
}
