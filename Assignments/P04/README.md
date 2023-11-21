## P04 - Concurrent Image Downloader
### Bao Trang

### Description:
Program 4 - Concurrent Image Downloader is a Go-based project designed to demonstrate the power of concurrency. It features two different image downloading methods: a sequential version and a concurrent version.

### Instructions

#### Pre-requisites:
- Install the Go programming language.
- Ensure Git is installed for version control.
- A working internet connection is required to download images.

#### Setting up and Running:
1. **Navigate to the project directory.**
2. **Download Dependencies:** go mod download
3. **Run the program:** go run main.go

#### Usage:
- Place a text file named `urls.txt` in the project directory with image URLs for downloading.
- The program will read the URLs from this file and download the images.
- Two sets of images will be downloaded: one using sequential downloading and the other using concurrent downloading.
- The program will display the time taken for each method, showcasing the efficiency of concurrent operations in Go.

#### Notes:
- The images will be saved in the project directory with unique names.
- Ensure the URLs in `urls.txt` are direct links to image files.
- The program includes error handling for failed downloads or invalid URLs.
