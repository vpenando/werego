package werego

import (
	"encoding/binary"
	"fmt"
	"io"
	"os"
)

func openFile(path string) ([][]byte, error) {
	buffer := make([][]byte, 0)
	file, err := os.Open(path)
	if err != nil {
		return buffer, err
	}
	var opuslen int16
	for {
		// Read opus frame length from dca file.
		err = binary.Read(file, binary.LittleEndian, &opuslen)

		// If this is the end of the file, just return.
		if err == io.EOF || err == io.ErrUnexpectedEOF {
			err := file.Close()
			if err != nil {
				return buffer, err
			}
			return buffer, nil
		}

		if err != nil {
			fmt.Println("Error reading from file :", err)
			return buffer, err
		}

		// Read encoded pcm from dca file.
		inBuf := make([]byte, opuslen)
		err = binary.Read(file, binary.LittleEndian, &inBuf)

		// Should not be any end of file errors
		if err != nil {
			fmt.Println("Error reading from dca file :", err)
			return buffer, err
		}

		// Append encoded pcm data to the buffer.
		buffer = append(buffer, inBuf)
	}
}
