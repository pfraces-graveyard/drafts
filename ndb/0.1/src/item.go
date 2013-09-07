package item


import (
	"./block"
	"./io"
	"fmt"
)


type Item struct {
	block *block.Block
	next byte
	entry byte
}

func New(io *io.Io) *Item {
	return &Item{block.New(io), 0, 0}
}

func (item *Item) Data() []byte {
	return item.block.Data()[1:]
}

func (item *Item) Next() byte { return item.next }
func (item *Item) Entry() byte { return item.entry }

func (item *Item) Length() byte {
	l := item.block.Length()
	if l == 0 { return 0 }
	return l - 1
}

func (item *Item) Read(entry byte) {
	if entry == 0 { entry++	}
	item.entry = entry
	item.block.Read(item.entry)
	if item.block.Length() == 0 {
		item.next = 0
	} else {
		item.next = item.block.Data()[0]
	}
}

func (item *Item) SetNext(entry byte, next byte) {
	if entry == 0 { entry++	}
	item.entry = entry
	item.next = next
	item.block.Io().Write(item.entry + 1, []byte{item.next})
}

func (item *Item) Write(entry byte, next byte, data []byte) {
	if entry == 0 { entry++	}
	item.entry = entry
	item.next = next
	item.block.Write(item.entry, append([]byte{item.next}, data...))
}

func (item *Item) Append(next byte, data []byte) {
	item.Write(item.block.Io().Size(), next, data)
}

func (item *Item) Dump() {
	fmt.Println("Length:", item.Length(), "Next:", item.Next())
	for i, v := range item.Data() {
		fmt.Println(i, v, string(v))
	}
}

