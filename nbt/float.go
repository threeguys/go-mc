package nbt

import (
	"encoding/binary"
	"fmt"
)

/////////////////////////////////////////////////////////////////////////////////////////////////////
// FloatTag
//
type FloatTag struct {
	BaseTag
	value float32
}

func (tag *FloatTag) GetType() int {
	return TAG_Float
}

func NewFloatTag(name string) *FloatTag {
	tag := new(FloatTag)
	tag.BaseTag.name = name
	return tag
}

func (tag *FloatTag) read(rdr *NBTReader) error {
	var data float32
	err := binary.Read(rdr.reader, binary.BigEndian, &data)
	if err != nil {
		return err
	}
	tag.value = data
	return nil
}

func (tag *FloatTag) Print(indent string) {
	fmt.Printf("%sFLOAT[%s] %f\n", indent, tag.BaseTag.name, tag.value)
}
