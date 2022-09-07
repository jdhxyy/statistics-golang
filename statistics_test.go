package statistics_golang

import (
	"fmt"
	"testing"
)

func TestCase1(t *testing.T) {
	id1 := Register("item1")
	id2 := Register("item2")
	id3 := Register("item3")

	Add(id1)
	Set(id2, 5)
	Add(id3)
	Clear(id3)
	Print()
	ClearAll()
	fmt.Println(Output())
}
