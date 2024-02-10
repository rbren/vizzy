package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"sync"

	"github.com/rbren/vizzy/test/pkg/tcase"
)

const testDir = "test/e2e/cases/"

func main() {
	var wg sync.WaitGroup

	dirSize := len(strings.Split(testDir, "/"))
	err := filepath.Walk(testDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() || path == testDir {
			return nil
		}
		if len(strings.Split(path, "/")) != dirSize {
			// We want dirs like ./test/e2e/cases/foo, but not ./test/e2e/cases/foo/bar
			return nil
		}
		name := strings.Split(path, "/")[dirSize-1]

		wg.Add(1)
		go func(path string) {
			defer wg.Done()
			err := tcase.RunTestCase(name)
			if err != nil {
				panic(err)
			}

		}(path)
		return nil
	})

	if err != nil {
		fmt.Printf("Error walking through the directory: %v\n", err)
		return
	}

	wg.Wait()
	fmt.Println("Completed processing all directories.")
}
