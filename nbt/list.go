package nbt

import (
	"container/list"
	"encoding/binary"
	"fmt"
)

/////////////////////////////////////////////////////////////////////////////////////////////////////
// ListTag
//
type ListTag struct {
	BaseTag
	values *list.List
}

func (tag *ListTag) GetType() int {
	return TAG_List
}

func NewListTag(name string) *ListTag {
	tag := new(ListTag)
	tag.BaseTag.name = name
	tag.values = list.New()
	return tag
}

func (tag *ListTag) read(rdr *NBTReader) error {
	c, err := rdr.reader.ReadByte()
	if err != nil {
		return err
	}

	var len int32
	err = binary.Read(rdr.reader, binary.BigEndian, &len)
	if err != nil {
		return err
	}
	
	for i := 0; i < int(len); i++ {
		t, err := rdr.readTagData(c, "")
		if err != nil {
			return err
		}
		if t == nil {
			return NBTError{"nil tag value while reading list"}
		}
		tag.values.PushBack(t)
	}
	return nil
}

func (tag *ListTag) Print(indent string) {
	fmt.Printf("%sLIST[%s]\n", indent, tag.BaseTag.name)
	elem := tag.values.Front()
	for elem != nil {
		v := elem.Value.(NBTTag)
		v.Print(indent + "\t")
		elem = elem.Next()
	}
	fmt.Printf("%sEND[%s]\n", indent, tag.BaseTag.name)
}
