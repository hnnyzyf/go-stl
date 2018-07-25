package pair

import (
	"github.com/hnnyzyf/go-stl/container/value"
)

type Pair interface {
	value.Value
	GetKey() interface{}
	GetValue() interface{}
}

//The Pair could be bool,string,interge,float,pointers and so on,we implement some Pair we always use

//bool Pair
type BoolPair struct {
	Key bool
	Val interface{}
}

func Bool(Key bool, Val interface{}) *BoolPair {
	return &BoolPair{Key, Val}
}

func (e *BoolPair) Less(Val value.Value) bool {
	if e.Key == true && Val.(*BoolPair).Key == false {
		return true
	} else {
		return false
	}
}

func (e *BoolPair) More(Val value.Value) bool {
	if e.Key == false && Val.(*BoolPair).Key == true {
		return true
	} else {
		return false
	}
}

func (e *BoolPair) Equal(Val value.Value) bool {
	return e.Key == Val.(*BoolPair).Key
}

func (e *BoolPair) GetKey() interface{} {
	return e.Key
}

func (e *BoolPair) GetValue() interface{} {
	return e.Val
}

//String Pair
type StringPair struct {
	Key string
	Val interface{}
}

func String(Key string, Val interface{}) *StringPair {
	return &StringPair{Key, Val}
}

func (e *StringPair) Less(Val value.Value) bool {
	return e.Key < Val.(*StringPair).Key
}

func (e *StringPair) More(Val value.Value) bool {
	return e.Key > Val.(*StringPair).Key
}

func (e *StringPair) Equal(Val value.Value) bool {
	return e.Key > Val.(*StringPair).Key
}

func (e *StringPair) GetKey() interface{} {
	return e.Key
}

func (e *StringPair) GetValue() interface{} {
	return e.Val
}

//Uint64 Pair
type Uint64Pair struct {
	Key uint64
	Val interface{}
}

func Uint64(Key uint64, Val interface{}) *Uint64Pair {
	return &Uint64Pair{Key, Val}
}

func (e *Uint64Pair) Less(Val value.Value) bool {
	return e.Key < Val.(*Uint64Pair).Key
}

func (e *Uint64Pair) More(Val value.Value) bool {
	return e.Key > Val.(*Uint64Pair).Key
}

func (e *Uint64Pair) Equal(Val value.Value) bool {
	return e.Key == Val.(*Uint64Pair).Key
}

func (e *Uint64Pair) GetKey() interface{} {
	return e.Key
}

func (e *Uint64Pair) GetValue() interface{} {
	return e.Val
}

//Uint64 Pair
type IntPair struct {
	Key int
	Val interface{}
}

func Int(Key int, Val interface{}) *IntPair {
	return &IntPair{Key, Val}
}

func (e *IntPair) Less(Val value.Value) bool {
	return e.Key < Val.(*IntPair).Key
}

func (e *IntPair) More(Val value.Value) bool {
	return e.Key > Val.(*IntPair).Key
}

func (e *IntPair) Equal(Val value.Value) bool {
	return e.Key == Val.(*IntPair).Key
}

func (e *IntPair) GetKey() interface{} {
	return e.Key
}

func (e *IntPair) GetValue() interface{} {
	return e.Val
}

//float64 Pair
type Floate64Pair struct {
	Key float64
	Val interface{}
}

func Float64(Key float64, Val interface{}) *Floate64Pair {
	return &Floate64Pair{Key, Val}
}

func (e *Floate64Pair) Less(Val value.Value) bool {
	return e.Key < Val.(*Floate64Pair).Key
}

func (e *Floate64Pair) More(Val value.Value) bool {
	return e.Key > Val.(*Floate64Pair).Key
}

func (e *Floate64Pair) Equal(Val value.Value) bool {
	return e.Key == Val.(*Floate64Pair).Key
}

func (e *Floate64Pair) GetKey() interface{} {
	return e.Key
}

func (e *Floate64Pair) GetValue() interface{} {
	return e.Val
}

//rune Pair
type RunePair struct {
	Key rune
	Val interface{}
}

func Rune(Key rune, Val interface{}) *RunePair {
	return &RunePair{Key, Val}
}

func (e *RunePair) Less(Val value.Value) bool {
	return e.Key < Val.(*RunePair).Key
}

func (e *RunePair) More(Val value.Value) bool {
	return e.Key > Val.(*RunePair).Key
}

func (e *RunePair) Equal(Val value.Value) bool {
	return e.Key == Val.(*RunePair).Key
}

func (e *RunePair) GetKey() interface{} {
	return e.Key
}

func (e *RunePair) GetValue() interface{} {
	return e.Val
}
