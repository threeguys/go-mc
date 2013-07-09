package nbt

import (
	"encoding/binary"
	"fmt"
	"io"
)

/////////////////////////////////////////////////////////////////////////////////////////////////////
// ByteArrayTag
//
type ByteArrayTag struct {
	BaseTag
	length int32
	values []byte
}

func (tag *ByteArrayTag) GetType() int {
	return TAG_Byte_Array
}

func NewByteArrayTag(name string) *ByteArrayTag {
	tag := new(ByteArrayTag)
	tag.BaseTag.name = name
	tag.length = 0
	return tag
}

func (tag *ByteArrayTag) read(rdr *NBTReader) error {
	var len int32
	err := binary.Read(rdr.reader, binary.BigEndian, &len)
	
	if err != nil {
		return err
	}
	
	if len < 0 {
		return NBTError{"Negative length encountered reading byte array tag"}
	}
	
	
	if len > 0 {
		bytes := make([]byte, len)
		count, err := io.ReadFull(rdr.reader, bytes)
		if err != nil {
			return err
		}
		
		if count != int(len) {
			return NBTError{ fmt.Sprintf("Expected %d bytes but read %d bytes when reading byte array", len, count) }
		}
		
		tag.length = len
		tag.values = bytes
	}
	
	return nil
}

func (tag *ByteArrayTag) Print(indent string) {
	fmt.Printf("%sBYTES[%s]: {%d} bytes\n", indent, tag.BaseTag.name, tag.length)
}
