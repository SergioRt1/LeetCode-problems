package main

import (
	"fmt"
	"io"
	"math/rand"
	"net/http"
	"net/url"
	"os"
	"path"
	"strconv"
	"sync"
)

const (
	concurrencyLimit = 10
)

type nothing struct{}

func readInput() []string {
	input := make([]string, 0)
	var line string
	for {
		_, err := fmt.Scanln(&line)
		if err != nil {
			return input
		}
		input = append(input, line)
	}
}

func createFolder(destination string) error {
	err := os.MkdirAll(destination, os.ModePerm)
	if err != nil {
		return err
	}
	err = os.Chdir(destination)
	if err != nil {
		return err
	}
	return nil
}

func getFileName(urlRaw string) string {
	URL, err := url.Parse(urlRaw)
	if err != nil {
		return strconv.Itoa(rand.Int())
	} else {
		return path.Base(URL.Path)
	}
}

func makeHttpAndSaveInFile(url string, filepath string) error {
	out, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer out.Close()

	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	_, err = io.Copy(out, resp.Body)
	if err != nil {
		return err
	}

	return nil
}

func download(urls []string) {
	concurrencyLimiter := make(chan nothing, concurrencyLimit)
	wg := sync.WaitGroup{}

	for _, u := range urls {
		concurrencyLimiter <- nothing{}
		wg.Add(1)
		go func(URL string) {
			defer func() {
				<-concurrencyLimiter
				wg.Done()
			}()
			name := getFileName(URL)
			err := makeHttpAndSaveInFile(URL, name)
			if err != nil {
				fmt.Printf("error downloading %s in path %s: %s\n", URL, name, err.Error())
			} else {
				fmt.Printf("succesfuly downloaded %s, as %s\n", URL, name)
			}
		}(u)
	}

	wg.Wait()
}

func main() {
	/* Sample input:
	https://img.freepik.com/free-photo/fluffy-kitten-sitting-grass-staring-sunset-playful-generated-by-artificial-intelligence_25030-67836.jpg
	https://images.unsplash.com/photo-1529778873920-4da4926a72c2
	*/
	destination := "dest/"
	urls := readInput()

	err := createFolder(destination)
	if err != nil {
		panic(err)
	}

	download(urls)
}
