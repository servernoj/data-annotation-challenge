package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

type Entry struct {
	Index int
	Text  string
}

type Entries []Entry

func (a Entries) Len() int           { return len(a) }
func (a Entries) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a Entries) Less(i, j int) bool { return a[i].Index < a[j].Index }

func main() {
	if len(os.Args) < 2 {
		log.Fatalf("Usage: %s <file-to-check>", os.Args[0])
	}
	f, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)
	entries := Entries{}
	for scanner.Scan() {
		line := scanner.Text()
		words := regexp.MustCompile(`\s+`).Split(line, 2)
		index, err := strconv.Atoi(words[0])
		if err != nil {
			log.Fatal(err)
		}
		entries = append(entries, Entry{
			Index: index,
			Text:  words[1],
		})
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	sort.Sort(entries)
	next := 0
	skipSize := 2
	result := strings.Builder{}
	for idx, entry := range entries {
		if idx == next {
			result.WriteString(entry.Text)
			result.WriteString(" ")
			next = next + skipSize
			skipSize++
		}
	}
	fmt.Printf("%s\n", strings.TrimSpace(result.String()))
}
