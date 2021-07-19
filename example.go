package main

import (
	"fmt"
	"github.com/NubeIO/nubeio-rubix-lib-helpers-go/pkg/bools"
	"github.com/NubeIO/nubeio-rubix-lib-helpers-go/pkg/strings"
	"github.com/NubeIO/nubeio-rubix-lib-helpers-go/pkg/uuid"
)

func main() {

	str := strings.New("what$ up !n the hood ")
	fmt.Println(str.RemoveSpecialCharacter())

	b, err := bools.Boolean("on on");if err != nil {
		fmt.Println(err)
	}
	fmt.Println(b)
	bb, _ := bools.Boolean("0")
	fmt.Println(bb)
	bbb, _ := bools.Boolean("True")
	fmt.Println(bbb)

	u, _ := uuid.MakeUUID()
	fmt.Println(u)



}

