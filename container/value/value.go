package value

type Value interface {
	Less(Value) bool
	More(Value) bool
	Equal(Value) bool
}

//we provide some  to use

type Uint64 uint64

func (i Uint64) Less(v Value) bool {
	return i < v.(Uint64)
}

func (i Uint64) More(v Value) bool {
	return i > v.(Uint64)
}

func (i Uint64) Equal(v Value) bool {
	return i == v.(Uint64)
}

type Int uint64

func (i Int) Less(v Value) bool {
	return i < v.(Int)
}

func (i Int) More(v Value) bool {
	return i > v.(Int)
}

func (i Int) Equal(v Value) bool {
	return i == v.(Int)
}

type String string

func (i String) Less(v Value) bool {
	return i < v.(String)
}

func (i String) More(v Value) bool {
	return i > v.(String)
}

func (i String) Equal(v Value) bool {
	return i == v.(String)
}

type Float64 float64

func (i Float64) Less(v Value) bool {
	return i < v.(Float64)
}

func (i Float64) More(v Value) bool {
	return i > v.(Float64)
}

func (i Float64) Equal(v Value) bool {
	return i == v.(Float64)
}

type Rune rune

func (i Rune) Less(v Value) bool {
	return i < v.(Rune)
}

func (i Rune) More(v Value) bool {
	return i > v.(Rune)
}

func (i Rune) Equal(v Value) bool {
	return i == v.(Rune)
}
