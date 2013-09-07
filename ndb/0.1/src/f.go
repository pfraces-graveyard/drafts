package main


import (
	"./column"
	"./io"
)

// que reciba la db de un parametro de linea de comandos
func main() {
	db := io.Open("/home/pau/dev/go/ndb/0.1/foo.ndb")
	defer db.Close()

	col := column.New(db)
/*	col.Create([]byte("LISTA1"))
	col.Create([]byte("LISTA2"))
	col.Append([]byte("LISTA1"), []byte("Hello "))
	col.Append([]byte("LISTA2"), []byte("foo "))
	col.Append([]byte("LISTA1"), []byte("World!"))
	col.Append([]byte("LISTA2"), []byte("manchoo"))
*/	col.Dump([]byte("LISTA1"))
	col.Dump([]byte("LISTA2"))
}

