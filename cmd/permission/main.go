package main

import (
	"fmt"

	"github.com/openslides/openslides-permission-service/internal/permission"
)

type fakeDataProvider struct{}

func (dp *fakeDataProvider) Get(fqfields []permission.Fqfield) map[permission.Fqfield]permission.Value {
	m := make(map[permission.Fqfield]permission.Value)
	for i, val := range fqfields {
		m[val] = i
	}
	return m
}

func main() {
	myDataProvider := &fakeDataProvider{}
	myFakeData := []permission.Fqfield{"Foo", "Bar"}
	result := permission.DoIt(myDataProvider, myFakeData)
	fmt.Println(result)
}
