package column


import (
	"fmt"
	"./io"
	"./list"
	"bytes"
)


const (
	ENONE = iota
	EFOUND
	ENOTFOUND
	EOUTOFRANGE
)


type Column struct {
	io *io.Io
}


func New(io *io.Io) *Column {
	return &Column{io}
}

func (col *Column) entry(idx byte) (err int, entry byte) {
	l := list.New(col.io)
	l.Read(1)
	if idx < 0 || idx >= l.Count() {return EOUTOFRANGE, 0}
	return ENONE, l.At(idx).Data()[0]
}

func (col *Column) name(idx byte) (err int, name []byte) {
	l := list.New(col.io)
	l.Read(1)
	if idx < 0 || idx >= l.Count() {return EOUTOFRANGE, []byte{}}
	return ENONE, l.At(idx).Data()[1:]
}

func (col *Column) setEntry(idx byte, entry byte) int {
	l := list.New(col.io)
	l.Read(1)
	if idx < 0 || idx >= l.Count() {return EOUTOFRANGE}
	l.Update(idx, append([]byte{entry}, l.At(idx).Data()[1:]...))
	return ENONE
}

func (col *Column) find(uid []byte) (err int, idx byte) {
	l := list.New(col.io)
	l.Read(1)
	if l.Count() == 0 {return ENOTFOUND, 0}
	i, f := byte(0), l.Count() - 1
	for i <= f {
		m := (i + f) / 2
		_, mname := col.name(m)
		switch bytes.Compare(uid, mname) {
		case 0:
			return ENONE, m
		case -1:
			f = m - 1
		case 1:
			i = m + 1
		}
	}
	return ENOTFOUND, i
}

func (col *Column) Create(uid []byte) int {
	err, idx := col.find(uid)
	if err == ENONE {return EFOUND}
	l := list.New(col.io)
	l.Read(1)
	l.Insert(idx, append([]byte{0}, uid...))
	return ENONE
}

func (col *Column) Drop(uid []byte) int {
	err, idx := col.find(uid)
	if err != ENONE {return err}
	l := list.New(col.io)
	l.Read(1)
	l.Delete(idx)
	return ENONE
}

func (col *Column) Count(uid []byte) (err int, count byte) {
	err, idx := col.find(uid)
	if err != ENONE {return err, 0}
	_, entry := col.entry(idx)
	if entry == 0 {return ENONE, 0}
	l := list.New(col.io)
	l.Read(entry)
	return ENONE, l.Count()
}

func (col *Column) Insert(uid []byte, idx byte, data []byte) int {
	err, i := col.find(uid)
	if err != ENONE {return err}
	_, colCount := col.Count(uid)
	if idx < 0 || idx > colCount {return EOUTOFRANGE}
	l := list.New(col.io)
	_, entry := col.entry(i)
	if entry != 0 {
		l.Read(entry)
	}
	l.Insert(idx, data)
	return col.setEntry(i, l.Entry())
	return ENONE
}

func (col *Column) Append(uid []byte, data []byte) int {
	err, colCount := col.Count(uid)
	if err != ENONE {return err}
	return col.Insert(uid, colCount, data)
}

func (col *Column) Delete(uid []byte, idx byte) int {
	err, i := col.find(uid)
	if err != ENONE {return err}
	_, colCount := col.Count(uid)
	if idx < 0 || idx >= colCount {return EOUTOFRANGE}
	l := list.New(col.io)
	_, entry := col.entry(i)
	l.Read(entry)
	l.Delete(idx)
	return col.setEntry(i, l.Entry())
}

func (col *Column) Update(uid []byte, idx byte, newData []byte) int {
	err, i := col.find(uid)
	if err != ENONE {return err}
	_, colCount := col.Count(uid)
	if idx < 0 || idx >= colCount {return EOUTOFRANGE}
	l := list.New(col.io)
	_, entry := col.entry(i)
	l.Read(entry)
	l.Update(idx, newData)
	return col.setEntry(i, l.Entry())
}

func (col *Column) Select(uid []byte, idx byte) (err int, data []byte) {
	err, i := col.find(uid)
	if err != ENONE {return err, []byte{}}
	_, colCount := col.Count(uid)
	if idx < 0 || idx >= colCount {return EOUTOFRANGE, []byte{}}
	l := list.New(col.io)
	_, entry := col.entry(i)
	l.Read(entry)
	return ENONE, l.At(idx).Data()
}

func (col *Column) Dump(uid []byte) {
	err, colCount := col.Count(uid)
	if err != ENONE {
		fmt.Println("err:", err)
		return
	}
	fmt.Println(string(uid))
	for i := byte(0); i < colCount; i++ {
		_, data := col.Select(uid, i)
		fmt.Println(i, string(data))
	}
}

