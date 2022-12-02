package bong

type BongMap map[string]Bong

type Bong struct {
	Title   string `yaml:"Title,omitempty"`
	MainUrl string `yaml:"Main URL"`
	BongUrl string `yaml:"Bong URL"`
	Bongus  string `yaml:"Bongus"`
}
