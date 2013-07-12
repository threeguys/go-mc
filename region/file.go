package region

import (
	"bufio"
	"container/list"
	"encoding/binary"
	"fmt"
	"io"
	"github.com/threeguys/go-mc/nbt"
	"os"
)

type Region struct {
	locations [1024]uint32
	timestamps [1024]uint32
	chunks [1024]*Chunk
}

func ReadRegion(fileName string) (*Region, error) {

	f, err := os.Open(fileName)
	
	if err != nil {
		return nil, err
	}
	
	totalBytes := 0
	totalBlocks := 0
	
	rdr := bufio.NewReader(f)
	
	locations := make([]uint32, 1024)
	err = binary.Read(rdr, binary.BigEndian, locations)
	
	totalBytes += 4 * 1024
	
	if err != nil {
		fmt.Printf("Error reading region file %s\n", err.Error())
		return nil, err
	}
	
	timestamps := make([]uint32, 1024)
	err = binary.Read(rdr, binary.BigEndian, timestamps)
	
	totalBytes += 4 * 1024
	
	if err != nil {
		fmt.Printf("Error reading timestamps %s\n", err.Error())
		return nil, err
	}
	
	tags := list.New()

	for err == nil && totalBlocks < 1024 {
		var len int32
		err = binary.Read(rdr, binary.BigEndian, &len)
		totalBytes += 4
		
		if err != nil {
			fmt.Printf("Error reading chunk length %s\n", err.Error())
			return nil, err
		}
		
		c, err := rdr.ReadByte()
		totalBytes++
		if err != nil {
			fmt.Printf("Error reading compression type %s\n", err.Error())
			return nil, err
		}
		
		var nbtRdr *nbt.NBTReader
		switch c {
			case 1 :
				nbtRdr, err = nbt.NewNBTGzipReader(rdr)
				
			case 2 :
				nbtRdr, err = nbt.NewNBTZlibReader(rdr)
				
			default :
				msg := fmt.Sprintf("Unknown compression type %d\n", c)
				fmt.Printf(msg)
				fmt.Printf("Total bytes: %d\n", totalBytes)
				fmt.Printf("Total blocks: %d\n", totalBlocks)
				err = nbt.NewError(msg)
				return nil, err
		}
		
		totalBytes += int(len)
		
		if err != nil {
			fmt.Printf("Reader error %s\n", err.Error())
		} else {
			tag, err := nbtRdr.Read()
		
			if err != nil {
				return nil, err
			}
		
			if err == nil && tag != nil {
				totalBlocks++
				tags.PushBack(tag)
			}
		}
		
		diff := 4096 - (int(len) % 4096)
		if diff > 0 {
			bytes := make([]byte, diff)
			count, e := io.ReadFull(rdr, bytes)
			err = e
			totalBytes += count
			
			if count != diff {
				msg := fmt.Sprintf("Expected %d bytes but read %d while reading padding\n", diff, count)
				fmt.Printf(msg)
				err = nbt.NewError(msg)
				return nil, err
			}
		}
	}
	
	if err != nil {
		fmt.Printf("Finished reading blocks: Error = %s\n", err.Error())
		return nil, err
	} else {
		fmt.Printf("Read %d blocks with no errors\n", totalBlocks)
	}
	
	f.Close()
	return nil, nil
}
