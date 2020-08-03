package permission

type DataProvider interface {
	Get(fqfields []Fqfield) map[Fqfield]Value
}

type Fqfield string
type Value interface{}

func DoIt(dataprovider DataProvider, fqfields []Fqfield) map[Fqfield]Value {
	return dataprovider.Get(fqfields)
}
