package bong

func (bm BongMap) ToSlice() []Bong {
	bs := make([]Bong, len(bm))
	i := 0
	for _, b := range bm {
		bs[i] = b
		i++
	}

	return bs
}

func SliceToBongMap(bongs []Bong) BongMap {
	bm := make(BongMap)
	for _, b := range bongs {
		bm[b.Bongus] = b
	}
	return bm
}
