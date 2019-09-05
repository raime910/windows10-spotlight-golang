package main

import (
	"image/jpeg"
	"io"
	"log"
	"os"
	"os/user"
	"path/filepath"
)

func main() {

	root := getAssetsFolder()
	log.Printf("Scanning root folder (%s) for HD images", root)

	wallpapers := getHdWallpapers(root)
	log.Printf("Found %d HD wallpapers within the root folder", len(wallpapers))

	copySuccess := createCopies(wallpapers)

	if copySuccess {
		log.Println("Yay! new wallpapers!")
	}
}

func createCopies(wallpapers []string) bool {
	for _, wallpaper := range wallpapers {
		// read file.
		from, err := os.Open(wallpaper)

		// check for errors.
		if err != nil {
			log.Fatal(err)
		}
		defer from.Close()

		fileName := filepath.Base(wallpaper)

		to, err := os.OpenFile("./wallpapers/"+fileName+".jpeg", os.O_CREATE|os.O_RDWR, 0666)

		// check for errors.
		if err != nil {
			log.Fatal(err)
		}
		defer to.Close()

		_, err = io.Copy(to, from)
		if err != nil {
			log.Fatal(err)
		}
	}

	return true
}

func isHighResImage(path string) bool {
	// read file.
	file, err := os.Open(path)

	// check for errors.
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	image, err := jpeg.Decode(file)

	if err != nil {
		return false
	}

	bounds := image.Bounds()
	isHD := bounds.Max.X >= 1920

	return isHD
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

func getHdWallpapers(root string) []string {
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

	// lets find 1080p images.
	var wallpapers []string
	for _, path := range files {
		if isHighResImage(path) {
			wallpapers = append(wallpapers, path)
		}
	}

	return wallpapers
}
