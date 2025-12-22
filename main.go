package main

import "fmt"

func main() {
 i, j := 42, 2701
 fmt.Println(&i, &j)
 
 p := &i

 fmt.Println(*p)
 *p = 21
 fmt.Println(i)

 p = &j
 *p = *p / 37
 fmt.Println(j)
}