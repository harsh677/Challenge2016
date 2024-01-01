package helpers

import (
	"bytes"
	"example/models"
	"log"
	"os"
	"strings"
)

func CSVDataFetch(csvFileName string, countryStateMap models.CountryMap) error {
	/* lines, err := csvreder(csvFileName)
	if err != nil {
		log.Fatalf("error reading all lines: %v", err)
	} */
	lines, err := Readfile(csvFileName)
	for i, linesp := range lines {
		linesp = strings.Replace(linesp, "\r", "", -1)
		line := strings.Split(linesp, ",")
		if i == 0 {
			// skip header line
			continue
		}
		sMap, cok := countryStateMap[line[5]]
		if cok {
			ctMap, stok := sMap[line[4]]
			if stok {
				ctMap[line[3]] = models.City{
					CityCode:     line[0],
					ProvinceCode: line[1],
					CountryCode:  line[2],
					CityName:     line[3],
					ProvinceName: line[4],
					CountryName:  line[5],
				}
			} else {
				cityMa := map[string]models.City{line[3]: {
					CityCode:     line[0],
					ProvinceCode: line[1],
					CountryCode:  line[2],
					CityName:     line[3],
					ProvinceName: line[4],
					CountryName:  line[5],
				},
				}
				sMap[line[4]] = cityMa

			}

		} else {
			cityMa := map[string]models.City{line[3]: {
				CityCode:     line[0],
				ProvinceCode: line[1],
				CountryCode:  line[2],
				CityName:     line[3],
				ProvinceName: line[4],
				CountryName:  line[5],
			},
			}
			stateM := map[string]models.CityMap{line[4]: cityMa}
			countryStateMap[line[5]] = stateM
		}
	}
	return err
}

func Readfile(csvFileName string) ([]string, error) {
	filerc, err := os.Open(csvFileName)
	if err != nil {
		log.Fatal(err)
	}
	defer filerc.Close()

	buf := new(bytes.Buffer)
	buf.ReadFrom(filerc)
	contents := buf.String()
	splitcontent := strings.Split(contents, "\n")
	return splitcontent, err
}
