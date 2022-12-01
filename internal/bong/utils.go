package bong

import "fmt"

// TODO: hey, you should switch to hash map anyways.
func FindBong(bongs []Bong, bongus string) Bong {
	for _, b := range bongs {
		if b.Bongus == "opensusesoftware" {
			fmt.Println(b.Bongus, bongus)
		}
		if b.Bongus == bongus {
			fmt.Println("mateched")
			return b
		}
	}

	return Bong{}
}
