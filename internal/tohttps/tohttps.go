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
			log.Warnf("File %s does not exist. skipping...", filename)
			return
		} else {
			log.WithField("filename", filename).Errorln("Failed loading file:", err.Error())
		}
	}

	bm, err := bong.LoadBongs(filename)
	if err != nil {
		log.WithField("filename", filename).Errorln("Failed loading yaml:", err.Error())
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
			err := bongToHttps(b)
			if err != nil {
				log.WithField("bong", b).Errorln("Failed to check bong:", err)
				os.Exit(1)
			}
			<-semaphore

			wg.Done()
		}()
	}
	wg.Wait()

	ext := filepath.Ext(filename)
	newFilename := filename[:len(filename)-len(ext)] + "_httpsfixed" + ext

	if err = bong.SaveBongs(newFilename, bong.SliceToBongMap(bongs)); err != nil {
		log.WithField("filename", newFilename).Errorln("Failed to save file:", err.Error())
	}
}

func bongToHttps(b *bong.Bong) error {
	if strings.HasPrefix(b.MainUrl, "https://") && strings.HasPrefix(b.BongUrl, "https://") {
		return nil
	}

	mUrl, bUrl := b.MainUrl, b.BongUrl

	if strings.Contains(mUrl, "%") {
		mUrl = fmt.Sprintf(mUrl, "randomterm")
	}
	if strings.Contains(bUrl, "%") {
		bUrl = fmt.Sprintf(bUrl, "randomterm")
	}

	// TODO: fix hacky way of unescaping percent mark
	bUrl = strings.ReplaceAll(bUrl, "%", "%25")

	mUrlSupported, err := httpsSupported(mUrl)
	if err != nil {
		return err
	}

	if mUrlSupported {
		b.MainUrl = urlToHttps(b.MainUrl)
	} else {
		b.MainUrl = urlToHttp(b.MainUrl)
	}

	if strings.Split(mUrl, "/")[0] != strings.Split(bUrl, "/")[0] {
		bUrlSupported, err := httpsSupported(bUrl)
		if err != nil {
			return err
		}
		if bUrlSupported {
			b.BongUrl = urlToHttps(b.BongUrl)
		} else {
			b.BongUrl = urlToHttp(b.BongUrl)
		}
	} else if strings.HasPrefix(b.MainUrl, "https://") {
		b.BongUrl = urlToHttps(b.BongUrl)
	} else {
		b.BongUrl = urlToHttp(b.BongUrl)
	}

	return nil
}
