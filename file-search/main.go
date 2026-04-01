package main

import (
	"flag"
	"fmt"
	"os"
	"regexp"
	"strings"
	"sync"
)

var EXTENSIONS_TO_READ = []string{
	"ipynb",
	"md",
	"txt",
	"py",
	"Dockerfile",
	"java",
	"js",
	"tf",
	"ts",
	"go",
	"xml",
	"yaml",
	"yml",
	"json",
	"sh",
}

var (
	repoPath     *string = flag.String("repo", "/tmp/go-git-test", "Path to the git repository")
	searchString *string = flag.String("search", "string-to-search-for", "String to search for")
	outputFile   *string = flag.String("output", "output.txt", "Output file")
	spaceRe      string  = `([-_]|\s*\n\s*(#|\/\/)\s*|\s*",[\s\n]*"\s*|\s*)`
	mut          sync.Mutex
)

func main() {
	flag.Parse()

	s := *searchString
	if strings.Contains(s, ",") {
		st := strings.Split(s, ",")
		for _, rss := range st {
			err := doSearch(*repoPath, rss)
			if err != nil {
				fmt.Println(err)
				continue
			}
		}
	} else {
		err := doSearch(*repoPath, *searchString)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	}

	os.Exit(0)
}

func doSearch(path, search string) error {
	fmt.Printf("Searching %s for '%s'...\n", path, search)
	err := os.Chdir(*repoPath)
	if err != nil {
		return err
	}

	es, err := os.ReadDir(".")
	if err != nil {
		fmt.Println(err)
		return err
	}

	var wg sync.WaitGroup
	wg.Add(len(es))

	for _, entry := range es {
		go func(e os.DirEntry) {
			walkFiles(e, ".", search)
			wg.Done()
		}(entry)
	}
	wg.Wait()
	return nil
}

func walkFiles(entry os.DirEntry, path, search string) error {
	sp := fmt.Sprintf("%s/%s", path, entry.Name())
	if entry.IsDir() {
		subEntries, err := os.ReadDir(sp)
		if err != nil {
			return err
		}
		for _, subEntry := range subEntries {
			walkFiles(subEntry, sp, search)
		}
		return nil
	}

	if entry.Name() == ".DS_Store" {
		return nil
	}

	err := readFile(sp, search)
	if err != nil {
		return err
	}
	return nil
}

func readFile(path, search string) (err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("panic: %v", r)
			fmt.Println(err)
		}
	}()

	exs := strings.Split(path, ".")
	if len(exs) < 2 {
		return nil
	}

	ext := exs[len(exs)-1]

	if !contains(EXTENSIONS_TO_READ, ext) {
		return nil
	}

	data, err := os.ReadFile(path)
	if err != nil {
		return err
	}

	f := string(data)

	ts := strings.ReplaceAll(search, " ", spaceRe)
	re := regexp.MustCompile(ts)

	if re.MatchString(f) {
		printResult(search, path)
	}

	return err
}

func printResult(search, path string) {
	fmt.Printf("!,%s,%s\n", search, path)
	if *outputFile == "" {
		return
	}

	mut.Lock()
	defer mut.Unlock()

	f, err := os.OpenFile(*outputFile,
		os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		fmt.Println(err)
	}
	defer f.Close()
	if _, err := f.WriteString(fmt.Sprintf("!,%s,%s\n", search, path)); err != nil {
		fmt.Println(err)
	}
}

func contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}
