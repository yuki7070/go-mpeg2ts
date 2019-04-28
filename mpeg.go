package main

import (
	"fmt"
	"os"
)

func main() {
	f, err := os.Open("test.ts")
	if err != nil {
		panic(err)
	}

	ids := []byte{0x4E, 0x50, 0x51, 0x52, 0x53, 0x54, 0x55, 0x56, 0x57, 0x58, 0x59, 0x5a, 0x5b, 0x5c, 0x5d, 0x5e, 0x5f}

	data := TSFile(f, ids)
	fmt.Println(data)
	for _, s := range data.Section {
		for k, e := range s.Event {
			fmt.Println("=========================")
			fmt.Printf("%X", k)
			fmt.Println("\n")
			fmt.Printf("%X", e.ExtendedEventDescriptor.TextChar)
			fmt.Println("\n")
			fmt.Printf("%X", e.ShortEventDescriptor.EventNameChar)
			fmt.Println("\n")
			fmt.Printf("%X", e.ShortEventDescriptor.TextChar)
			fmt.Println("\n")
			demo := AribStr(e.ShortEventDescriptor.EventNameChar)
			fmt.Println(string(demo))
		}
	}
}
