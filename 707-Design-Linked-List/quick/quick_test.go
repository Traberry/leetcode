package quick

import (
	"testing"
	"fmt"
)

func Test_quick(t *testing.T){
	list := Constructor()
	list.AddAtHead(1)
	list.AddAtIndex(1, 2)
	fmt.Println(list.List)

	get := list.Get(1)
	fmt.Println(get)
}