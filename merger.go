package main

import (
	"fmt"
	"io"
	"io/fs"
	"log"
	"os"
	"slices"
)

type Merger struct {
	dirLocation    string
	outputFilename string
}

func NewMerger(dirLocation, outputFilename string) *Merger {
	return &Merger{dirLocation: dirLocation, outputFilename: outputFilename}
}

func (m Merger) Execute() error {
	filenameChan := make(chan string)
	go m.mustReadDir(m.dirLocation, filenameChan)
	for filename := range filenameChan {
		err := func() error {
			destinationFile, err := os.OpenFile(m.outputFilename, os.O_CREATE|os.O_APPEND|os.O_WRONLY, os.ModePerm)
			if err != nil {
				log.Fatalln(err)
			}
			defer destinationFile.Close()
			sourceFile, err := os.Open(fmt.Sprintf("./%s/%s", m.dirLocation, filename))
			if err != nil {
				log.Fatalln(err)
			}
			defer sourceFile.Close()
			_, err = io.Copy(destinationFile, sourceFile)
			destinationFile.WriteString("\n")
			if err != nil {
				return err
			}
			return nil
		}()
		if err != nil {
			return err
		}
	}
	return nil
}

func (m Merger) mustReadDir(dirLocation string, filenameChan chan string) {
	items, err := os.ReadDir(dirLocation)
	if err != nil {
		log.Panicln(err)
	}
	slices.SortFunc(items, func(a, b fs.DirEntry) int {
		aInfo, _ := a.Info()
		bInfo, _ := b.Info()
		return bInfo.ModTime().Compare(aInfo.ModTime())
	})
	for _, item := range items {
		if item.IsDir() {
			continue
		}
		filenameChan <- item.Name()
	}
	close(filenameChan)
}
