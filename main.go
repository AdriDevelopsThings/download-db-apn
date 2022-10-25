package main

import (
	"errors"
	"fmt"
	"os"
	"path"

	"github.com/adridevelopsthings/download-db-apn/pkg"
	"github.com/akamensky/argparse"
	"github.com/schollz/progressbar/v3"
)

func main() {
	parser := argparse.NewParser("download-db-apn", "Download all available apn sketches from trassenfinder.de")
	infrastructure_id := parser.Int("i", "infrastructure-id", &argparse.Options{Default: 12, Help: "Change the ID of the infrastructure, but I don't think you'll need this"})
	target_directory := parser.String("t", "target-directory", &argparse.Options{Default: "target", Help: "The directory where the apk sketches will be saved in"})

	err := parser.Parse(os.Args)
	if err != nil {
		fmt.Printf("Error while parsing cli arguments: %v\n", err)
		return
	}

	if _, err := os.Stat(*target_directory); errors.Is(err, os.ErrNotExist) {
		fmt.Printf("Directory %s does not exist, creating...\n", *target_directory)
		err = os.Mkdir(*target_directory, os.ModePerm)
		if err != nil {
			fmt.Printf("Error while creating directory %s: %v\n", *target_directory, err)
			return
		}
	}

	fmt.Println("Downloading infrastructure infos...")
	infrastructure, err := pkg.GetInfrastructure(*infrastructure_id)
	if err != nil {
		fmt.Printf("Error while downloading infrastructure: %v\n", err)
		return
	}
	betriebsstellen := infrastructure.Ordnungsrahmen.Betriebsstellen
	documents := make([]*pkg.Document, 0)
	alreadyDownloaded := 0
	fmt.Printf("Found %d betriebsstellen\n", len(betriebsstellen))
	documentBar := progressbar.Default(int64(len(betriebsstellen)), "Fetching document information")
	for _, betriebsstelle := range betriebsstellen {
		document, err := pkg.GetDocument(infrastructure.ID, &betriebsstelle)
		if err != nil {
			fmt.Printf("Error while getting document information about betriebsstelle %s: %v\n", betriebsstelle.DS100, err)
		}
		if document != nil {
			if _, err := os.Stat(path.Join(*target_directory, document.Filename)); errors.Is(err, os.ErrNotExist) {
				documents = append(documents, document)
			} else {
				alreadyDownloaded += 1
			}
		}
		documentBar.Add(1)
	}
	documentBar.Finish()
	if alreadyDownloaded > 0 {
		fmt.Printf("%d documents were already downloaded, skipping...\n", alreadyDownloaded)
	}
	bar := progressbar.Default(int64(len(documents)), "Downloading documents")
	for _, document := range documents {
		err := document.Download(*target_directory)
		if err != nil {
			fmt.Printf("Error while downloading document %s: %v", document.Filename, err)
		}
		bar.Add(1)
	}
	bar.Finish()
	fmt.Printf("The downloaded files were saved to %s\n", *target_directory)
}
