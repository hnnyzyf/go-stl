package stl

type Pair = Entry

type Entry interface {
	Value
	GetKey() interface{}
	GetValue() interface{}
}

//The Entry could be bool,string,interge,float,pointers and so on,we implement some entry we always use

//bool Entry
type BoolEntry struct {
	k bool
	v interface{}
}

func NewBoolEntry(k bool, v interface{}) *BoolEntry {
	return &BoolEntry{k, v}
}

func (e *BoolEntry) Less(v Value) bool {
	if e.k == true && v.(*BoolEntry).k == false {
		return true
	} else {
		return false
	}
}

func (e *BoolEntry) More(v Value) bool {
	if e.k == false && v.(*BoolEntry).k == true {
		return true
	} else {
		return false
	}
}

func (e *BoolEntry) GetKey() interface{} {
	return e.k
}

func (e *BoolEntry) GetValue() interface{} {
	return e.v
}

func (e *BoolEntry) Key() bool {
	return e.k
}

//String Entry
type StringEntry struct {
	k string
	v interface{}
}

func NewStringEntry(k string, v interface{}) *StringEntry {
	return &StringEntry{k, v}
}

func (e *StringEntry) Less(v Value) bool {
	return e.k < v.(*StringEntry).k
}

func (e *StringEntry) More(v Value) bool {
	return e.k > v.(*StringEntry).k
}

func (e *StringEntry) GetKey() interface{} {
	return e.k
}

func (e *StringEntry) GetValue() interface{} {
	return e.v
}

func (e *StringEntry) Key() string {
	return e.k
}

//Uint64 Entry
type Uint64Entry struct {
	k uint64
	v interface{}
}

func NewUint64Entry(k uint64, v interface{}) *Uint64Entry {
	return &Uint64Entry{k, v}
}

func (e *Uint64Entry) Less(v Value) bool {
	return e.k < v.(*Uint64Entry).k
}

func (e *Uint64Entry) More(v Value) bool {
	return e.k > v.(*Uint64Entry).k
}

func (e *Uint64Entry) GetKey() interface{} {
	return e.k
}

func (e *Uint64Entry) GetValue() interface{} {
	return e.v
}

func (e *Uint64Entry) Key() uint64 {
	return e.k
}

//Uint64 Entry
type IntEntry struct {
	k int
	v interface{}
}

func NewIntEntry(k int, v interface{}) *IntEntry {
	return &IntEntry{k, v}
}

func (e *IntEntry) Less(v Value) bool {
	return e.k < v.(*IntEntry).k
}

func (e *IntEntry) More(v Value) bool {
	return e.k > v.(*IntEntry).k
}

func (e *IntEntry) GetKey() interface{} {
	return e.k
}

func (e *IntEntry) GetValue() interface{} {
	return e.v
}

func (e *IntEntry) Key() int {
	return e.k
}

//float64 Entry
type Float64Entry struct {
	k float64
	v interface{}
}

func NewFloat64Entry(k float64, v interface{}) *Float64Entry {
	return &Float64Entry{k, v}
}

func (e *Float64Entry) Less(v Value) bool {
	return e.k < v.(*Float64Entry).k
}

func (e *Float64Entry) More(v Value) bool {
	return e.k > v.(*Float64Entry).k
}

func (e *Float64Entry) GetKey() interface{} {
	return e.k
}

func (e *Float64Entry) GetValue() interface{} {
	return e.v
}

func (e *Float64Entry) Key() float64 {
	return e.k
}

//rune Entry
type RuneEntry struct {
	k rune
	v interface{}
}

func NewRuneEntry(k rune, v interface{}) *RuneEntry {
	return &RuneEntry{k, v}
}

func (e *RuneEntry) Less(v Value) bool {
	return e.k < v.(*RuneEntry).k
}

func (e *RuneEntry) More(v Value) bool {
	return e.k > v.(*RuneEntry).k
}

func (e *RuneEntry) GetKey() interface{} {
	return e.k
}

func (e *RuneEntry) GetValue() interface{} {
	return e.v
}

func (e *RuneEntry) Key() rune {
	return e.k
}
