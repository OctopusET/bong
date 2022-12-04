package tohttps

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"sync"

	"github.com/npmania/bong/internal/bong"
	log "github.com/sirupsen/logrus"
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
	_, err := os.Stat(filename)
	if err != nil {
		if os.IsNotExist(err) {
			log.Warnf("file %s does not exist. skipping...\n", filename)
			return
		} else {
			log.Warnf("failed loading file %s: error %s\n", filename, err.Error())
		}
	}

	ext := filepath.Ext(filename)
	newFilename := filename[:len(filename)-len(ext)] + "_httpsfixed" + ext

	bm, err := bong.LoadBongs(filename)
	if err != nil {
		log.Warnf("failed loading yaml from %s: error %s\n", filename, err.Error())
		return
	}

	var wg sync.WaitGroup

	semaphore := make(chan bool, 5)

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
