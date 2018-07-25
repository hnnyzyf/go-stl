package value

import (
	"fmt"
	"testing"
)

func dasd() Value{
	var a Int
	return a
}


func Test_nil(t *testing.T){
	b:=dasd()

	if b==nil{
		fmt.Println(b)
	}
}
