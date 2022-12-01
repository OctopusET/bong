package bong

type BongCollection struct {
	Bongs map[string]Bong `yaml:"Bongs"`
}

type Bong struct {
	Title   string `yaml:"Title,omitempty"`
	MainUrl string `yaml:"Main URL"`
	BongUrl string `yaml:"Bong URL"`
	Bongus  string `yaml:"Bongus"`
}

func newBongCollection() (bc BongCollection) {
	bc.Bongs = make(map[string]Bong)
	return
}

func toCollection(bongs []Bong) BongCollection {
	bc := newBongCollection()

	for _, b := range bongs {
		bc.Bongs[b.Bongus] = b
	}

	return bc
}

func (bc BongCollection) toBongs() (bongs []Bong) {
	for _, b := range bc.Bongs {
		bongs = append(bongs, b)
	}

	return
}
