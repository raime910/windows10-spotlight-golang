package main

import (
	"fmt"
	"image"
	"log"
	"os"
	"os/user"
	"path/filepath"
)

func main() {

	var root = getAssetsFolder()
	var allFiles = getFilePaths(root)

	// lets find 1080p images.
	for _, filePath := range allFiles {
		fmt.Println(filePath)
		isHighResImage(filePath)
	}
}

func isHighResImage(filePath string) bool {
	// read file.
	file, err := os.Open(filePath)

	// check for errors.
	if err != nil {
		log.Println("Failed to open the source file.")
		log.Fatal(err)
	}
	defer file.Close()

	image, _, err := image.Decode(file)

	if err != nil {
		log.Println("Failed to decode the file.")
		log.Fatal(err)
	}

	bounds := image.Bounds()
	fmt.Printf("%s is %dx%d", file.Name(), bounds.Max.X, bounds.Max.Y)

	return false
}

func getAssetsFolder() string {
	user, err := user.Current()

	if err != nil {
		log.Fatal(err)
	}

	// build the root path to the assets folder.
	root := user.HomeDir + "\\AppData\\Local\\Packages\\Microsoft.Windows.ContentDeliveryManager_cw5n1h2txyewy\\LocalState\\Assets"

	return root
}

func getFilePaths(root string) []string {
	// start collecting the files within it.
	var files []string
	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if info.IsDir() == false {
			files = append(files, path)
		}
		return nil
	})

	// check for errors.
	if err != nil {
		log.Fatal(err)
	}

	return files
}
