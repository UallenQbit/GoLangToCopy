package GoLangToCopy

import (
	"io"
	"net"
	"sync"
)

func Copy(Writer net.Conn, Reader net.Conn) int64 {
	Size := new(int64)
	WaitGroup := new(sync.WaitGroup)

	for Index := 0; Index < 2; Index++ {
		WaitGroup.Add(1)
		go func(Writer net.Conn, Reader net.Conn, Index int, Size *int64, WaitGroup *sync.WaitGroup) {
			defer func(Writer net.Conn, Reader net.Conn) {
				Reader.Close()
				Writer.Close()
			}(Writer, Reader)

			var WriterSize int64

			if Index == 0 {
				WriterSize, _ = io.Copy(Writer, Reader)
			} else {
				WriterSize, _ = io.Copy(Reader, Writer)
			}

			*Size += WriterSize

			WaitGroup.Done()
		}(Writer, Reader, Index, Size, WaitGroup)
	}

	WaitGroup.Wait()

	return *Size
}

func CopyBuffer(Writer net.Conn, Reader net.Conn, Buffer []byte) int64 {
	Size := new(int64)
	WaitGroup := new(sync.WaitGroup)

	for Index := 0; Index < 2; Index++ {
		WaitGroup.Add(1)
		go func(Writer net.Conn, Reader net.Conn, Buffer []byte, Index int, Size *int64, WaitGroup *sync.WaitGroup) {
			defer func(Writer net.Conn, Reader net.Conn) {
				Reader.Close()
				Writer.Close()
			}(Writer, Reader)

			var WriterSize int64

			if Index == 0 {
				WriterSize, _ = io.CopyBuffer(Writer, Reader, Buffer)
			} else {
				WriterSize, _ = io.CopyBuffer(Reader, Writer, Buffer)
			}

			*Size += WriterSize

			WaitGroup.Done()
		}(Writer, Reader, Buffer, Index, Size, WaitGroup)
	}

	WaitGroup.Wait()

	return *Size
}
