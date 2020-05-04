// main.go
package main

import (
	"./base"
	_ "./ex1"
	_ "./ex2"
)

//因為上面使用匿名導入了ex1 & ex2 package.
//main()執行前, 這兩個package的init()會被調用, 而註冊了class1 & class2
func main() {
	c1 := base.Create("Class1")
	c1.Do()
}
