package main

import (
	"flag"
	"image"
	"image/jpeg"
	"io/ioutil"
	"log"
	"os"
	"path"
	"strings"
	"sync"

	"github.com/nfnt/resize"
)

var v = flag.Int("v", 1, "thumbnail version")

func main() {
	flag.Parse()
	dir := "images/"
	var files []string
	fi, err := ioutil.ReadDir(dir)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	for _, fInfo := range fi {
		n := fInfo.Name()
		if strings.HasSuffix(n, ".jpg") || strings.HasSuffix(n, ".jpeg") {
			files = append(files, dir+fInfo.Name())
		}
	}

	log.Println("start converting...")
	switch *v {
	case 1:
		makeThumbnails(files)
	case 2:
		makeThumbnails2(files)
	case 3:
		makeThumbnails3(files)
	case 4:
		makeThumbnails4(files)
	case 5:
		_, err := makeThumbnails5(files)
		if err != nil {
			log.Fatal(err)
		}
	case 6:
		makeThumbnails6(files)
	default:
		makeThumbnails(files)
	}
	log.Println("finished converting.")
}

func makeThumbnails(filenames []string) {
	for _, filename := range filenames {
		if _, err := ImageFile(filename); err != nil {
			log.Println(filename, err)
		}
	}
}

func makeThumbnails2(filenames []string) {
	for _, filename := range filenames {
		go ImageFile(filename)
	}
}

func makeThumbnails3(filenames []string) {
	ch := make(chan struct{})
	for _, filename := range filenames {
		go func(f string) {
			ImageFile(f)
			ch <- struct{}{}
		}(filename)
	}

	for range filenames {
		<-ch
	}
}

func makeThumbnails4(filenames []string) error {
	errors := make(chan error)
	for _, filename := range filenames {
		go func(f string) {
			_, err := ImageFile(f)
			errors <- err
		}(filename)
	}

	for range filenames {
		if err := <-errors; err != nil {
			return err // goroutine leak
		}
	}
	return nil
}

func makeThumbnails5(filenames []string) (thumbfiles []string, err error) {
	type item struct {
		thumbfile string
		err       error
	}
	ch := make(chan item, len(filenames))

	for _, f := range filenames {
		go func(f string) {
			var it item
			it.thumbfile, it.err = ImageFile(f)
			ch <- it
		}(f)
	}

	for range filenames {
		it := <-ch
		if it.err != nil {
			return nil, it.err
		}
		thumbfiles = append(thumbfiles, it.thumbfile)
	}
	return thumbfiles, nil
}

// makeThumbnails6 makes thumbnails for each file received from the channel.
// It returns the number of bytes occupied by the files it creates.
func makeThumbnails6(filenames []string) int64 {
	sizes := make(chan int64)
	var wg sync.WaitGroup // counter. number of working goroutines
	for _, f := range filenames {
		wg.Add(1)
		go func(f string) {
			defer wg.Done()
			thumb, err := ImageFile(f)
			if err != nil {
				log.Println(err)
				return
			}
			info, _ := os.Stat(thumb)
			sizes <- info.Size()
		}(f)
	}

	go func() {
		wg.Wait()
		close(sizes)
	}()

	var total int64
	for size := range sizes {
		total += size
	}
	return total
}

func ImageFile(infile string) (string, error) {
	r, err := os.Open(infile)
	if err != nil {
		return "", err
	}
	defer r.Close()

	// Load image
	img, _, err := image.Decode(r)
	if err != nil {
		return "", err
	}
	newImg := resize.Resize(160, 0, img, resize.Lanczos3)

	// Create new file name with 'tmb'
	dir := path.Dir(infile)
	file := path.Base(infile)
	extension := path.Ext(infile)
	file = strings.Replace(file, extension, ".tmb"+extension, 1)
	filePath := path.Join(dir, file)

	// Write image
	w, err := os.Create(filePath)
	if err != nil {
		return "", err

	}
	defer w.Close()
	err = jpeg.Encode(w, newImg, nil)
	if err != nil {
		return "", err
	}
	return filePath, nil
}
