package tohttps

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"sync"

	"github.com/npmania/bong/internal/bong"
)

func FilesToHttps(files []string) {
	if !hasInternetConnection() {
		panic(fmt.Errorf("cannot access internet"))
	}
	for _, path := range files {
		fileToHttps(path)
	}
}

func fileToHttps(filename string) {
	fi, err := os.Stat(filename)
	if err != nil {
		if os.IsNotExist(err) {
			fmt.Printf("file %s does not exist. skipping...\n", filename)
			return
		} else {
			fmt.Printf("failed loading file %s: error %s\n", filename, err.Error())
		}
	}

	newFilename := fi.Name() + "_httpsfixed" + filepath.Ext(filename)

	bm, err := bong.LoadBongs(filename)
	if err != nil {
		fmt.Printf("failed reading yaml from %s: error %s\n", filename, err.Error())
		return
	}

	var wg sync.WaitGroup

	semaphore := make(chan bool, 50)

	bongs := bm.ToSlice()
	shuffle(bongs)

	for bg := range bongs {
		wg.Add(1)
		semaphore <- true

		b := &bongs[bg]
		go func() {
			bongToHttps(b)
			<-semaphore

			wg.Done()
		}()
	}
	wg.Wait()

	if err = bong.SaveBongs(newFilename, bong.SliceToBongMap(bongs)); err != nil {
		fmt.Printf("failed saving file %s: %s", newFilename, err)
	}
}

func bongToHttps(b *bong.Bong) {
	if strings.HasPrefix(b.MainUrl, "https://") && strings.HasPrefix(b.BongUrl, "https://") {
		return
	}

	mUrl, bUrl := b.MainUrl, b.BongUrl

	if strings.Contains(mUrl, "%") {
		mUrl = fmt.Sprintf(mUrl, "randomterm")
	}
	bUrl = fmt.Sprintf(bUrl, "randomterm")

	if httpsSupported(mUrl) {
		b.MainUrl = urlToHttps(b.MainUrl)
	} else {
		b.MainUrl = urlToHttp(b.MainUrl)
	}

	if strings.Split(mUrl, "/")[0] != strings.Split(bUrl, "/")[0] {
		if httpsSupported(bUrl) {
			b.BongUrl = urlToHttps(b.BongUrl)
		} else {
			b.BongUrl = urlToHttp(b.BongUrl)
		}
	} else if strings.HasPrefix(b.MainUrl, "https://") {
		b.BongUrl = urlToHttps(b.BongUrl)
	} else {
		b.BongUrl = urlToHttp(b.BongUrl)
	}
}
