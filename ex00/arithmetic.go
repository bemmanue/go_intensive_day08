package main

import (
	"errors"
	"fmt"
	"log"
	"unsafe"
)

func getElement(arr []int, idx int) (int, error) {
	if len(arr) == 0 {
		return 0, errors.New("empty array")
	} else if idx >= len(arr) || idx < 0 {
		return 0, errors.New("index is out of bounds")
	}

	start := unsafe.Pointer(&arr[0])
	size := unsafe.Sizeof(int(0))

	item := *(*int)(unsafe.Pointer(uintptr(start) + size*uintptr(idx)))

	return item, nil
}

func main() {
	array := []int{19, 30, 5, 691, 10, 2}

	res, err := getElement(array, 14)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(res)
}
