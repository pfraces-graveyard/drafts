package io


import "os"


type Io struct {
	file *os.File
}


func Open(path string) *Io {
	f, _ := os.Open(path, os.O_RDWR|os.O_CREATE, 0666)
	io := &Io{f}
	if io.Size() == 0 {
		io.Write(0, []byte{0})
	}
	return io
}

func (io *Io) Close() {
	io.file.Close()
}

func (io *Io) Read(entry byte, length byte) []byte {
	data := make([]byte, length)
	io.file.ReadAt(data, int64(entry))
	return data
}

func (io *Io) Write(entry byte, data []byte) {
	io.file.WriteAt(data, int64(entry))
}

func (io *Io) Size() byte {
	fi, _ := io.file.Stat()
	return byte(fi.Size)
}

