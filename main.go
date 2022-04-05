package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func enhance(out string, filename string, data [][]string) {
	currentTime := time.Now()
	dateTime := currentTime.Format("2006-01-02")

	s := strings.Split(filename, "/")
	a := s[len(s)-1]
	f, _ := os.Create(out + "/" + a)
	defer f.Close()

	w := csv.NewWriter(f)
	defer w.Flush()

	for i, line := range data {
		if i == 0 {
			//	// append headers
			line = append(line, "DateTime")
			line = append(line, "FileName")
		}
		if i > 0 {
			line = append(line, dateTime)
			line = append(line, filename)
		}
		w.Write(line)

	}
}

func enhanceFile(out string, filename string) {
	f, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	csvReader := csv.NewReader(f)
	data, err := csvReader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	enhance(out, filename, data)

}

func enhanceDir(out string, path string) {
	fmt.Printf("From folder < %s, to folder > %s", path, out)

	files, err := ioutil.ReadDir(path)
	if err != nil {
		log.Fatal(err)
	}

	for _, f := range files {
		if filepath.Ext(f.Name()) == ".csv" {
			fmt.Println(f.Name())
			enhanceFile(out, path+"/"+f.Name())
		}
	}
}

func main() {
	folder := flag.String("in", ".", "Input folder")
	out := flag.String("out", "out", "Output folder")
	// dateFormat := flag.String("df", "", "Format for the current date")

	flag.Parse()

	if len(flag.Args()) == 0 {
		flag.PrintDefaults()
	}

	if _, err := os.Stat(*out); os.IsNotExist(err) {
		if err := os.MkdirAll(*out, os.ModePerm); err != nil {
			log.Fatal(err)
		}
	}

	enhanceDir(*out, *folder)
}
