package main

import (
	"archive/zip"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
)

var nameOfFile string
var dest = "copyofvid.mp4"

type UrlVideo struct {
	url string
}

type FileSystemVideo struct {
	source string
}
type Zip struct {
}

func (v UrlVideo) Download() error {
	// Create blank file

	fileURL, err := url.Parse(v.url)
	if err != nil {
		log.Fatal(err)
	}
	path := fileURL.Path
	segments := strings.Split(path, "/")
	nameOfFile = segments[len(segments)-1]

	file, err := os.Create(nameOfFile)
	if err != nil {
		return err
	}

	client := http.Client{}

	// Put content on file

	resp, err := client.Get(v.url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	size, err := io.Copy(file, resp.Body)

	defer file.Close()

	fmt.Printf("Downloaded a file %s with size %d \n", nameOfFile, size)
	return nil
}

func (v2 FileSystemVideo) Download() error {
	// copying content of source file into destination file

	bytesRead, err := ioutil.ReadFile(v2.source)

	if err != nil {
		return err
	}

	err = ioutil.WriteFile(dest, bytesRead, 0644)

	if err != nil {
		return err
	}

	fmt.Println("Downloaded a file from local system")
	return nil
}

func (z Zip) Archive(names []string) error {

	//making zip file
	fmt.Println("creating zip archive...")

	archive, err := os.Create("archive.zip")
	if err != nil {
		return err
	}
	zipWriter := zip.NewWriter(archive)

	defer archive.Close()

	//adding 2 files in the zip file
	for i, ch := range names {
		fmt.Printf("opening file  %d...", i)
		f1, err := os.Open(ch)
		if err != nil {
			return err
		}
		defer f1.Close()

		fmt.Println("writing file to archive...")
		w1, err := zipWriter.Create(ch)
		if err != nil {
			return err
		}
		if _, err := io.Copy(w1, f1); err != nil {
			return err
		}
	}
	fmt.Println("closing zip archive...")
	zipWriter.Close()
	return nil
}

func main() {

	fullURLFile := "https://filesamples.com/samples/video/mp4/sample_640x360.mp4"
	src := "/Users/nikita.mogha/Downloads/sample_960x400_ocean_with_audio.mp4"

	var v Downloader
	v = UrlVideo{url: fullURLFile}

	err := v.Download()

	if err != nil {
		log.Fatal(err)
	}

	v = FileSystemVideo{source: src}

	err = v.Download()

	if err != nil {
		log.Fatal(err)
	}

	nameOfAllFile := []string{nameOfFile, dest}

	var z Archiver
	z = Zip{}

	err = z.Archive(nameOfAllFile)

	if err != nil {
		log.Fatal(err)
	}

}
