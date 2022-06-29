package main

import (
	"fmt"
	"github.com/rqixy/moduleTest/module1"
	"github.com/rqixy/moduleTest/module2"
)

func main() {
	// module1から呼び出す
	m1 := module1.ReturnHello()
	fmt.Println(m1)

	// module2から呼び出す
	module2.StartServer()
}
