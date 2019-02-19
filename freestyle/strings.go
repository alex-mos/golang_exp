package main

import (
	"fmt"
)

func main() {
	// const sample = "\xbd\xb2\x3d\xbc\x20\xe2\x8c\x98"
	const sample = "\xe2\x8c\x98"

	fmt.Printf("% x\n", sample)
	fmt.Printf("%q\n", sample)

	// for i := 0; i < len(sample); i++ {
	// fmt.Printf("%q ", sample[i])
	// }

	const sign = "\u2318"

	fmt.Printf("%+q\n", sign)

	const face = "\u1F62C"

	fmt.Printf("%q\n", face)

	const eyes = "\u1F644"

	fmt.Printf("%+q\n", eyes)

	const face2 = "😬"

	fmt.Printf("%q\n", face2)
}
