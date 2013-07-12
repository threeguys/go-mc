package nbt

import (
	"fmt"
	"container/list"
)

/////////////////////////////////////////////////////////////////////////////////////////////////////
// CompoundTag
//
type CompoundTag struct {
	BaseTag
	values *list.List
}

func (tag *CompoundTag) GetType() int {
	return TAG_Compound
}

func NewCompoundTag(name string) *CompoundTag {
	tag := new(CompoundTag)
	tag.BaseTag.name = name
	tag.values = list.New()
	return tag
}

func (tag *CompoundTag) read(rdr *NBTReader) error {
	
	for {
		c, e := rdr.reader.ReadByte()
		if e != nil {
			return e
		}
		
		if c == TAG_End {
			return nil
		}
		
		rdr.reader.UnreadByte()
		item, e := rdr.Read()
		
		if e != nil {
			return e
		}
		
		if item != nil {
			tag.values.PushBack(item)
		}
	}
}

func (tag *CompoundTag) Print(indent string) {
	fmt.Printf("%sCOMPOUND[%s]\n", indent, tag.BaseTag.name)
	elem := tag.values.Front()
	for elem != nil {
		v := elem.Value.(NBTTag)
		v.Print(indent + "\t")
		elem = elem.Next()
	}
	fmt.Printf("%sEND[%s]\n", indent, tag.BaseTag.name)
}
