package bong

import "fmt"

type BongMap map[string]Bong

func (bm BongMap) validate() error {
	for _, b := range bm {
		err := b.validate()
		if err != nil {
			return err
		}
	}
	return nil
}

type Bong struct {
	Title   string `yaml:"Title"`
	MainUrl string `yaml:"Main URL"`
	BongUrl string `yaml:"Bong URL"`
	Bongus  string `yaml:"Bongus"`
}

func (b Bong) validate() error {
	if b.Title == "" {
		return fmt.Errorf("no Title in Bong %+v", b)
	} else if b.MainUrl == "" {
		return fmt.Errorf("no MainUrl in Bong %+v", b)
	} else if b.BongUrl == "" {
		return fmt.Errorf("no BongUrl in Bong %+v", b)
	} else if b.Bongus == "" {
		return fmt.Errorf("no Bongus in Bong %+v", b)
	}

	return nil
}
