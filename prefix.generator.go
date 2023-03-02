package us_phone_generator

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/PuerkitoBio/goquery"

	HttpRunner "github.com/Tanreon/go-http-runner"
	NetworkRunner "github.com/Tanreon/go-network-runner"
)

// https://www.fonefinder.net/findome.php?npa=XXX
func generateFromFoneFinder(npaCodes map[string][]string) {
	workDir, _ := os.Getwd()

	directDialer, err := NetworkRunner.NewDirectDialer()
	if err != nil {
		log.Fatal(err)
	}

	httpRunner, err := HttpRunner.NewDirectHttpRunner(directDialer)
	if err != nil {
		log.Fatal(err)
	}

	prefixesMobile := make(map[string][]string, 0)
	prefixesOther := make(map[string][]string, 0)

	for _, codes := range npaCodes {
		//mobilePrefixes := make([]string, 0)
		//otherPrefixes := make([]string, 0)

		for _, code := range codes {
			cachedFileName := filepath.Join(workDir, "generate", "cache", hashMD5(code)+".html")

			if _, err := os.Stat(cachedFileName); os.IsNotExist(err) {
				htmlRequestOptions := HttpRunner.NewHtmlRequestOptions(fmt.Sprintf("https://www.fonefinder.net/findome.php?npa=%s", code))
				response, err := httpRunner.GetHtml(htmlRequestOptions)
				if err != nil {
					log.Fatal(err)
				}

				ioutil.WriteFile(cachedFileName, response.Body(), os.ModePerm)
			}

			cachedFile, err := os.OpenFile(cachedFileName, os.O_RDONLY, os.ModePerm)
			if err != nil {
				log.Fatal(err)
			}

			cachedReponse, err := ioutil.ReadAll(cachedFile)
			if err != nil {
				log.Fatal(err)
			}

			// Load the HTML document
			doc, err := goquery.NewDocumentFromReader(bytes.NewReader(cachedReponse))
			if err != nil {
				log.Fatal(err)
			}

			doc.Find("body table:nth-child(7) tr").Each(func(i int, rowSelection *goquery.Selection) {
				if rowSelection.Find("td").Length() > 0 {
					prefixSelection := rowSelection.Find("td:nth-child(2)")
					prefix := prefixSelection.Text()

					if len(prefix) <= 0 {
						return
					}

					telcoTypeSelection := rowSelection.Find("td:nth-child(6)")
					telcoType := telcoTypeSelection.Text()

					//if len(telcoType) <= 0 {
					//	return
					//}

					if strings.EqualFold(telcoType, "WIRELESS PROV") {
						if _, present := prefixesMobile[code]; !present {
							prefixesMobile[code] = make([]string, 0)
						}

						prefixesMobile[code] = append(prefixesMobile[code], prefix)
						//mobilePrefixes = append(mobilePrefixes, prefix)
					} else {
						if _, present := prefixesOther[code]; !present {
							prefixesOther[code] = make([]string, 0)
						}

						prefixesOther[code] = append(prefixesOther[code], prefix)
						//otherPrefixes = append(otherPrefixes, prefix)
					}
				}
			})

			//prefixesMobile[code] = mobilePrefixes
			//prefixesOther[code] = otherPrefixes

			//log.Println(prefixesMobile)
		}
	}

	fmt.Printf(`mobilePrefixCodes := make(map[int][]int, 0)` + "\n")
	for npaCode, prefixes := range prefixesMobile {
		fmt.Printf(`mobilePrefixCodes[%s] = []int{%s}`+"\n", npaCode, strings.Join(prefixes, ","))
	}

	fmt.Printf(`otherPrefixCodes := make(map[int][]int, 0)` + "\n")
	for npaCode, prefixes := range prefixesOther {
		fmt.Printf(`otherPrefixCodes[%s] = []int{%s}`+"\n", npaCode, strings.Join(prefixes, ","))
	}
}
