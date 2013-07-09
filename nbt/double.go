package nbt

import (
	"encoding/binary"
	"fmt"
)

/////////////////////////////////////////////////////////////////////////////////////////////////////
// DoubleTag
//
type DoubleTag struct {
	BaseTag
	value float64
}

func (tag *DoubleTag) GetType() int {
	return TAG_Double
}

func NewDoubleTag(name string) *DoubleTag {
	tag := new(DoubleTag)
	tag.BaseTag.name = name
	return tag
}

func (tag *DoubleTag) read(rdr *NBTReader) error {
	var data float64
	err := binary.Read(rdr.reader, binary.BigEndian, &data)
	if err != nil {
		return err
	}
	tag.value = data
	return nil
}

func (tag *DoubleTag) Print(indent string) {
	fmt.Printf("%sDOUBLE[%s] %f\n", indent, tag.BaseTag.name, tag.value)
}
