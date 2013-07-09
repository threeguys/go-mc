package nbt

import (
	"encoding/binary"
	"fmt"
)

/////////////////////////////////////////////////////////////////////////////////////////////////////
// IntArrayTag
//
type IntArrayTag struct {
	BaseTag
	length int32
	values []int32
}

func (tag *IntArrayTag) GetType() int {
	return TAG_Int_Array
}

func NewIntArrayTag(name string) *IntArrayTag {
	tag := new(IntArrayTag)
	tag.BaseTag.name = name
	tag.length = 0
	return tag
}

func (tag *IntArrayTag) read(rdr *NBTReader) error {
	var len int32
	err := binary.Read(rdr.reader, binary.BigEndian, &len)
	
	if err != nil {
		return err
	}
	
	if len < 0 {
		return NBTError{"Negative length encountered reading int array tag"}
	}
	
	
	if len > 0 {
		data := make([]int32, len)
		err := binary.Read(rdr.reader, binary.BigEndian, data)
		if err != nil {
			return err
		}
		
		//if count != int(len) {
		//	return NBTError{ fmt.Sprintf("Expected %d ints but read %d ints when reading int array", len, count) }
		//}
		
		tag.length = len
		tag.values = data
	}
	
	return nil
}

func (tag *IntArrayTag) Print(indent string) {
	fmt.Printf("%sINTS[%s]: {%d} elements\n", indent, tag.BaseTag.name, tag.length)
}
