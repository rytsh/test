package variadic

import (
	"fmt"
	"io"
)

type Data struct {
	Name string
	Age  int
}

func RandomDataGen(count int) []Data {
	data := make([]Data, count)
	for i := 0; i < count; i++ {
		data[i] = Data{
			Name: fmt.Sprintf("Name %d", i),
			Age:  i,
		}
	}
	return data
}

func CallVariadic(data ...Data) {
	for i := range data {
		dummy(&data[i].Name)
	}
}

func CallSlice(data []Data) {
	for i := range data {
		dummy(&data[i].Name)
	}
}

func dummy(name *string) {
	io.Discard.Write([]byte(*name))
}
