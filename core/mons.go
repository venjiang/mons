package core

type Data interface{}
type Mons map[string]func(...interface{}) Data

func NewMons() Mons {
	return make(map[string]func(...interface{}) Data)
}
