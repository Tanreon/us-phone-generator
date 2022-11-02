package us_phone_generator

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"
)

// CSV file from https://nationalnanpa.com/enas/geoAreaCodeNumberReport.do
func generateFromNpa() map[string][]string {
	workDir, _ := os.Getwd()

	csvFile, err := os.Open(filepath.Join(workDir, "generate", "npa.csv"))
	if err != nil {
		log.Fatal(err)
	}
	defer csvFile.Close()

	reader := bufio.NewReader(csvFile)
	csvReader := csv.NewReader(reader)
	//csvReader.Comma = ','
	csvReader.LazyQuotes = true

	//

	codes := make(map[string][]string, 0)

	for {
		row, err := csvReader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Println(err)
			continue
		}

		if len(row[1]) > 2 {
			continue
		}

		if _, present := codes[row[1]]; !present {
			codes[row[1]] = make([]string, 0)
		}

		codes[row[1]] = append(codes[row[1]], row[0])
	}

	fmt.Printf(`areaCodes := make(map[string][]int, 0)`)
	for state, codes := range codes {
		fmt.Printf(`areaCodes["%s"] = []int{%s}`+"\n", state, strings.Join(codes, ","))
		//fmt.Printf(`case "%s":`+"\n"+`code = randomFromSlice([]int{%s})`+"\n", state, strings.Join(codes, ","))
	}

	return codes
}
