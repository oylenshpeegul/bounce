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

func main() {

	tasks := make(chan *exec.Cmd, 64)

	// spawn four worker goroutines
	var wg sync.WaitGroup
	for i := 0; i < 4; i++ {
		wg.Add(1)
		go func() {
			for cmd := range tasks {
				cmd.Run()
			}
			wg.Done()
		}()
	}

	indir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	outdir := path.Join(indir, "mono")

	err = os.Mkdir(outdir, 0770)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(indir)
	fmt.Println(outdir)

	// generate some tasks
	files := getFlacs(indir)
	for i, f := range files {

		outfile := path.Join(
			outdir,
			strings.Replace(f.Name(), ".flac", ".mp3", -1),
		)

		fmt.Println(f.Name())
		fmt.Println(outfile)

		fmt.Println("task", i, "converting", f.Name(), "to", outfile)
		tasks <- exec.Command("sox", "-G", f.Name(), "-c", "1", outfile)
	}
	close(tasks)

	// wait for the workers to finish
	wg.Wait()
}

// getFlacs reads the given directory and returns those entries that
// end in .flac.
func getFlacs(dirName string) []os.FileInfo {

	entries, err := ioutil.ReadDir(dirName)
	if err != nil {
		log.Fatal(err)
	}

	files := []os.FileInfo{}

	for _, entry := range entries {
		//if strings.HasSuffix(entry.Name(), ".flac") {
		if path.Ext(entry.Name()) == ".flac" {
			files = append(files, entry)
		}
	}

	return files
}
