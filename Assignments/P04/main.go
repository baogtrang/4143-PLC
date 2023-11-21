package main

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
	"os"
	"sync"
	"time"
)

func main() {
	//open the file
	file, err := os.Open("urls.txt")
	if err != nil {
		fmt.Println("Error opening file: ", err)
		return
	}

	// ensure the file is closed when the function exits
	defer file.Close()

	// create a scanner to read the file
	scanner := bufio.NewScanner(file)

	// create a slice to store URLs
	var urls []string

	// read each line (URL) to append to the slice
	for scanner.Scan() {
		url := scanner.Text()
		urls = append(urls, url)
	}

	// check for errors during scanning
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading from file:", err)
		return
	}

	// print the URLs to verify
	fmt.Println("Read URLs: ")
	for _, url := range urls {
		fmt.Println(url)
	}

	// sequential download
	start := time.Now()
	downloadImagesSequential(urls)
	fmt.Printf("Sequential download took: %v\n", time.Since(start))

	// concurrent download
	start = time.Now()
	downloadImagesConcurrent(urls)
	fmt.Printf("Concurrent download took: %v\n", time.Since(start))

}

func downloadImage(url, filename string) error {
	// send a GET request to the URL
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Create a file to save the image
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	// Copy the response body (image data) to the file
	_, err = io.Copy(file, resp.Body)
	if err != nil {
		return err
	}

	return nil
}

func downloadImagesSequential(urls []string) {
	for i, url := range urls {
		filename := fmt.Sprintf("image_%d.jpg", i) // Unique filename for each image
		err := downloadImage(url, filename)
		if err != nil {
			fmt.Printf("Failed to download %s: %s\n", url, err)
			continue
		}
		fmt.Printf("Downloaded %s to %s\n", url, filename)
	}
}

func downloadImagesConcurrent(urls []string) {
	var wg sync.WaitGroup
	errChan := make(chan error, len(urls))

	for i, url := range urls {
		wg.Add(1)
		go func(url string, i int) {
			defer wg.Done()
			filename := fmt.Sprintf("image_%d.jpg", i)
			err := downloadImage(url, filename)
			if err != nil {
				errChan <- fmt.Errorf("Failed to download %s: %s", url, err)
			} else {
				fmt.Printf("Downloaded %s to %s\n", url, filename)
			}
		}(url, i)
	}

	wg.Wait()
	close(errChan)

	
	for err := range errChan {
		if err != nil {
			fmt.Println("Download error:", err)
		}
	}
}
