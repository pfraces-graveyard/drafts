package list


import (
	"./item"
	"./io"
	"fmt"
)


type List struct {
	item []*item.Item
	io *io.Io
	entry byte
}

func New(io *io.Io) *List {
	return &List{[]*item.Item{}, io, 0}
}

func (list *List) Read(entry byte) {
	list.entry = entry
	list.item = []*item.Item{}
	next := list.entry
	for next > 0 {
		i := item.New(list.io)
		i.Read(next)
		if i.Next() != byte(0) || i.Length() != byte(0) {
			list.item = append(list.item, i)
		}
		next = i.Next()
	}
}

func (list *List) Entry() byte { return list.entry }

func (list *List) At(idx byte) *item.Item {
	if idx < 0 || idx >= list.Count() { return nil }
	return list.item[idx]
}

func (list *List) Count() byte {
	return byte(len(list.item))
}

func (list *List) Insert(idx byte, data []byte) {
	if idx < 0 || idx > list.Count() { return }
	newItem := item.New(list.io)
	newNext := byte(0)
	if atItem := list.At(idx); atItem != nil {
		newNext = atItem.Entry()
	}
	newItem.Append(newNext, data)
	if idx == 0 {
		list.entry = newItem.Entry()
	} else {
		prevItem := list.At(idx - 1)
		prevItem.SetNext(prevItem.Entry(), newItem.Entry())
	}
	list.item = append(list.item[:idx], append([]*item.Item{newItem}, list.item[idx:]...)...)
}

func (list *List) Append(data []byte) {
	list.Insert(list.Count(), data)
}

// ? Debe o no el update generar un item extra por falta de espacio o por el contrario hacer un delete y un insert
// -> No debe. debe usar el espacio reservado si puede o reservar espacio para el item entero si no puede. Item es datos continuos
// ? Debe retornar el area eliminada?
// -> No, ya se encargara de las areas libres una capa superior
func (list *List) Update(idx byte, data []byte) {
	if idx < 0 || idx >= list.Count() { return }
	target := list.At(idx)
	if len(data) <= int(target.Length()) {
		target.Write(target.Entry(), target.Next(), data)
	} else {
		item := item.New(list.io)
		item.Append(target.Next(), data)
		if idx == 0 {
			list.entry = item.Entry()
		} else {
			prevItem := list.At(idx - 1)
			prevItem.SetNext( prevItem.Entry(), item.Entry())
		}
		list.item = append(append(list.item[:idx], item), list.item[idx + 1:]...)
	}
}

func (list *List) Delete(idx byte) {
	if idx < 0 || idx >= list.Count() { return }
	if idx == 0 {
		list.entry = list.At(1).Entry()
	} else {
		prevItem := list.At(idx - 1)
		prevItem.SetNext(prevItem.Entry(), list.At(idx).Next())
	}
	list.item = append(list.item[:idx], list.item[idx + 1:]...)
}

func (list *List) Dump() {
	for i, v := range list.item {
		fmt.Println(i, int(v.Entry()), string(v.Data()))
	}
}
