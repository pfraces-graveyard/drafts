package block


import (
	"./io"
	"fmt"
)


type Block struct {
	io *io.Io
	length byte
	data []byte
}


func New(io *io.Io) *Block {
	return &Block{io, 0, []byte("")}
}

func (block *Block) Io() *io.Io { return block.io }
func (block *Block) Length() byte { return block.length }
func (block *Block) Data() []byte { return block.data }

func (block *Block) Read(entry byte) {
	block.length = block.io.Read(entry, 1)[0]
	block.data = block.io.Read(entry + 1, block.length)
}

func (block *Block) Write(entry byte, data []byte) {
	block.data = data
	block.length = byte(len(block.data))
	block.io.Write(entry, []byte{block.length})
	block.io.Write(entry + 1, block.data)
}

func (block *Block) Append (data []byte) byte {
	entry := block.io.Size()
	block.Write(entry, data)
	return entry
}

func (block *Block) Dump() {
	fmt.Println(block.Length())
	for i, v := range block.Data() {
		fmt.Println(i, v, string(v))
	}
}

