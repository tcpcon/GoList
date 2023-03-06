# GoList
Data structure that adds QOL improvements to the standard slice type.

## Basic Example
```go
package main

import (
	"fmt"

	"github.com/ox-y/GoList"
	// go get github.com/ox-y/GoList
)

func main() {
	gl := golist.New[int]()

	gl.Add([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}...)
	
	for i, v := range gl.Unpack() {
		if v % 2 == 0 {
			gl.Set(i, 0)
		}
	}
	
	gl.Insert(2, []int{6, 9, 6, 9}...)

	fmt.Print(gl.Unpack()) // [1 0 6 9 6 9 3 0 5 0 7 0 9 0]
}
```

## Functions
```go
- New()
- Next()
- Insert()
- Remove()
- Replace()
- RemoveAll()
- ReplaceAll()
- RemoveAt()
- Set()
- Add()
- Clear()
- Get()
- Len()
- Unpack()
- Contains()
- IndexOf()
```
