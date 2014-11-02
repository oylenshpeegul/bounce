package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path"
	"strings"
	"sync"
)

const NumWorkers = 16

func main() {

	// TODO: Get input dir from the command line?
	indir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	// TODO: Get input file ext from the command line?
	inext := ".flac"

	// TODO: Get output file ext from the command line?
	outext := ".mp3"

	files := getFiles(indir, inext)

	if len(files) == 0 {
		fmt.Println("No files to process in", indir)
		os.Exit(1)
	}

	// TODO: Get output dir from the command line?
	outdir := path.Join(indir, "mono")

	err = os.Mkdir(outdir, 0770)
	if err != nil {
		log.Fatal(err)
	}

	tasks := make(chan *exec.Cmd, 64)

	// spawn worker goroutines
	var wg sync.WaitGroup
	for i := 0; i < NumWorkers; i++ {
		wg.Add(1)
		go func() {
			for cmd := range tasks {
				cmd.Run()
			}
			wg.Done()
		}()
	}

	// generate some tasks
	for i, f := range files {

		outfile := path.Join(
			outdir,
			strings.Replace(f.Name(), inext, outext, -1),
		)

		fmt.Println("Task", i)
		fmt.Println("  Converting", f.Name())
		fmt.Println("  To", outfile)

		tasks <- exec.Command("sox", "-G", f.Name(), "-c", "1", outfile)
	}
	close(tasks)

	// wait for the workers to finish
	wg.Wait()
}

// getFiles reads the given directory and returns a list of those
// entries that end in the given fileExt.
func getFiles(dirName string, fileExt string) []os.FileInfo {

	entries, err := ioutil.ReadDir(dirName)
	if err != nil {
		log.Fatal(err)
	}

	files := []os.FileInfo{}

	for _, entry := range entries {
		if path.Ext(entry.Name()) == fileExt {
			files = append(files, entry)
		}
	}

	return files
}
