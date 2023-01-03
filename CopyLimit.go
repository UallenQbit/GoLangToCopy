package GoLangToCopy

import (
	"io"
	"net"
	"sync"
)

func CopyLimit(Writer net.Conn, Reader net.Conn, Limit int64) int64 {
	Size := new(int64)
	WaitGroup := new(sync.WaitGroup)

	for Index := 0; Index < 2; Index++ {
		WaitGroup.Add(1)
		go func(Writer net.Conn, Reader net.Conn, Limit int64, Index int, Size *int64, WaitGroup *sync.WaitGroup) {
			defer func(Writer net.Conn, Reader net.Conn) {
				Reader.Close()
				Writer.Close()
			}(Writer, Reader)

			var WriterSize int64

			if Index == 0 {
				WriterSize, _ = io.Copy(Writer, io.LimitReader(Reader, Limit))
			} else {
				WriterSize, _ = io.Copy(Reader, io.LimitReader(Writer, Limit))
			}

			*Size += WriterSize

			WaitGroup.Done()
		}(Writer, Reader, Limit, Index, Size, WaitGroup)
	}

	WaitGroup.Wait()

	return *Size
}

func CopyLimitBuffer(Writer net.Conn, Reader net.Conn, Limit int64, Buffer []byte) int64 {
	Size := new(int64)
	WaitGroup := new(sync.WaitGroup)

	for Index := 0; Index < 2; Index++ {
		WaitGroup.Add(1)
		go func(Writer net.Conn, Reader net.Conn, Limit int64, Buffer []byte, Index int, Size *int64, WaitGroup *sync.WaitGroup) {
			defer func(Writer net.Conn, Reader net.Conn) {
				Reader.Close()
				Writer.Close()
			}(Writer, Reader)

			var WriterSize int64

			if Index == 0 {
				WriterSize, _ = io.CopyBuffer(Writer, io.LimitReader(Reader, Limit), Buffer)
			} else {
				WriterSize, _ = io.CopyBuffer(Reader, io.LimitReader(Writer, Limit), Buffer)
			}

			*Size += WriterSize

			WaitGroup.Done()
		}(Writer, Reader, Limit, Buffer, Index, Size, WaitGroup)
	}

	WaitGroup.Wait()

	return *Size
}
