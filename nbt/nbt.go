package nbt

import (
	"fmt"
	"io"
	"bufio"
	"os"
	"compress/gzip"
	"compress/zlib"
	"encoding/binary"
)

const (
	TAG_End = 0
	TAG_Byte = 1
	TAG_Short = 2
	TAG_Int = 3
	TAG_Long = 4
	TAG_Float = 5
	TAG_Double = 6
	TAG_Byte_Array = 7
	TAG_String = 8
	TAG_List = 9
	TAG_Compound = 10
	TAG_Int_Array = 11
)

type NBTError struct {
	msg string
}

type NBTTag interface {
	GetType() int
	Print(indent string)
	read(*NBTReader) error
}

func (err NBTError) Error() string {
	return err.msg
}

func readString(reader io.Reader, size int) (string, error) {
	if size > 0 {
		bytes := make([]byte, size)
		count, err := io.ReadFull(reader, bytes)
		if err != nil {
			fmt.Printf("Error reading name %s\n", err.Error())
			return "", err
		}
		
		if count != int(size) {
			e := NBTError{fmt.Sprintf("Could not read total amount of data for name, read %d, expected %d\n", count, size)}
			fmt.Printf(e.Error())
			return "", e
		}
		
		return string(bytes), nil
	}
	
	return "", nil
}

/////////////////////////////////////////////////////////////////////////////////////////////////////
// BaseTag
//
type BaseTag struct {
	name string
}

/////////////////////////////////////////////////////////////////////////////////////////////////////
// NBTReader
//
type NBTReader struct {
	file *os.File
	reader *bufio.Reader
}

func NewNBTZlibReader(reader io.Reader) (*NBTReader, error) {
	rdr, e := zlib.NewReader(reader)
	
	if e != nil {
		return nil, e
	}
	
	return NewNBTReader(rdr)
}

func NewNBTGzipReader(reader io.Reader) (*NBTReader, error) {
	rdr, e := gzip.NewReader(reader)
	
	if e != nil {
		return nil, e
	}
	
	return NewNBTReader(rdr)
}

func NewNBTReader(reader io.Reader) (*NBTReader, error) {
	rdr := bufio.NewReader(reader)
	
	retVal := new(NBTReader)
	retVal.reader = rdr
	return retVal, nil
}

func NewNBTFileReader(fileName string) (*NBTReader, error)  {
	f, e := os.Open(fileName)
	if e != nil {
		fmt.Printf("Error [%s] opening file %s for NBT reading\n", e.Error(), fileName)
		return nil, e
	}

	retVal, err := NewNBTGzipReader(f)
	
	if retVal != nil {
		retVal.file = f
	}
	
	return retVal, err 
}

func (rdr *NBTReader) Close() {
	if rdr.file != nil {
		rdr.file.Close()
	}
}

func (rdr *NBTReader) readTagData(c byte, name string) (NBTTag, error) {
	var tag NBTTag

	switch c {
		case TAG_End :
			fmt.Printf("TAG_End read\n");
			
		case TAG_Byte :
			fmt.Printf("TAG_Byte[%s]\n", name);
			tag = NewByteTag(name)
			
		case TAG_Short :
			fmt.Printf("TAG_Short[%s]\n", name);
			tag = NewShortTag(name)
			
		case TAG_Int :
			fmt.Printf("TAG_Int[%s]\n", name);
			tag = NewIntTag(name)
			
		case TAG_Long :
			fmt.Printf("TAG_Long[%s]\n", name);
			tag = NewLongTag(name)
			
		case TAG_Float :
			fmt.Printf("TAG_Float[%s]\n", name);
			tag = NewFloatTag(name)
			
		case TAG_Double :
			fmt.Printf("TAG_Double[%s]\n", name);
			tag = NewDoubleTag(name)
			
		case TAG_Byte_Array :
			fmt.Printf("TAG_Byte_Array read\n");
			tag = NewByteArrayTag(name)
			
		case TAG_String :
			fmt.Printf("TAG_String[%s]\n", name);
			tag = NewStringTag(name)
			
		case TAG_List :
			fmt.Printf("TAG_List[%s]\n", name);
			tag = NewListTag(name)
			
		case TAG_Compound :
			fmt.Printf("TAG_Compound[%s]\n", name);
			tag = NewCompoundTag(name)
			
		case TAG_Int_Array :
			fmt.Printf("TAG_Int_Array read\n");
			tag = NewIntArrayTag(name)
	}

	if tag != nil {
		err := tag.read(rdr)
		if err != nil {
			fmt.Printf("Error reading tag! %s\n", err.Error())
			return nil, err
		}
	}
	return tag, nil
}

func (rdr *NBTReader) Read() (NBTTag, error) {
	c, err := rdr.reader.ReadByte()
	
	if err != nil {
		fmt.Printf("Error reading tag %s\n", err.Error())
		return nil, err
	}
	
	if c < TAG_End || c > TAG_Int_Array {
		fmt.Printf("Unknown byte %i read\n", c)
		return nil, NBTError{fmt.Sprintf("Unknown byte %i read\n", c)}
	}
	
	var size int16
	err = binary.Read(rdr.reader, binary.BigEndian, &size)
	if err != nil {
		fmt.Printf("Error reading size of name %s\n", err.Error())
		return nil, err
	}
	
	name, err := readString(rdr.reader, int(size))
	if err != nil {
		return nil, err
	}
	
	return rdr.readTagData(c, name)
}

func readRegion(fileName string) {
	f, err := os.Open(fileName)
	
	rdr := bufio.NewReader(f)
	
	locations := make([]int32, 1024)
	err = binary.Read(rdr, binary.BigEndian, locations)
	
	if err != nil {
		fmt.Printf("Error reading region file %s\n", err.Error())
	} else {
		fmt.Printf("Successfully read locations\n");
	}
	
	timestamps := make([]int32, 1024)
	err = binary.Read(rdr, binary.BigEndian, timestamps)
	
	if err != nil {
		fmt.Printf("Error reading timestamps %s\n", err.Error())
	} else {
		fmt.Printf("Successfully read timestamps\n");
	}
	
	for i := 0; i<1024; i++ {
		offset := locations[i] >> 8
		if offset > 0 {
			fmt.Printf("Offset: %d\n", offset)
		}
	}
	
	var len int32
	err = binary.Read(rdr, binary.BigEndian, &len)
	if err != nil {
		fmt.Printf("Error reading chunk length %s\n", err.Error())
		return
	}
	
	fmt.Printf("Length: %d\n", len)
	
	c, err := rdr.ReadByte()
	if err != nil {
		fmt.Printf("Error reading compression type %s\n", err.Error())
		return
	}
	
	var nbtRdr *NBTReader
	switch c {
		case 1 :
			nbtRdr, err = NewNBTGzipReader(rdr)
			
		case 2 :
			nbtRdr, err = NewNBTZlibReader(rdr)
			
		default :
			fmt.Printf("Unknown compression type %d\n", c)
			return
	}
	
	if err != nil {
		fmt.Printf("Reader error %s\n", err.Error())
	} else {
	
		tag, err := nbtRdr.Read()
		
		if err == nil && tag != nil {
			tag.Print("")
		}
		
	}
	
	f.Close()
}

