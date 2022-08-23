package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
)

type weather struct {
	minTemp     int
	location    string
	time        string
	description string
}

func parseData(columns map[string]int, record []string) (*weather, error) {
	minTemp, err := strconv.Atoi(record[columns["MINtemp"]])
	if err != nil {
		return nil, err
	}
	location := record[columns["Location"]]
	time := record[columns["Time"]]
	description := record[columns["Description"]]

	return &weather{
		minTemp:     minTemp,
		location:    location,
		time:        time,
		description: description,
	}, nil

}

func main() {
	// open file
	f, err := os.Open("weather.csv")
	if err != nil {
		log.Fatalln(err)
	}
	defer f.Close()

	weatherLookup := map[string]*weather{}

	// parse csv file
	csvReader := csv.NewReader(f)
	columns := make(map[string]int)

	for rowCount := 0; ; rowCount++ {
		record, err := csvReader.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			log.Fatal(err)
		}

		if rowCount == 0 {
			for idx, column := range record {
				columns[column] = idx
			}
		} else {
			weather, err := parseData(columns, record)
			if err != nil {
				log.Fatal(err)
			}
			log.Println(weather)
			weatherLookup[weather.description] = weather
		}
	}

	if len(os.Args) < 2 {
		log.Fatalln("expected weather description")
	}

	wd := os.Args[1]
	d, ok := weatherLookup[wd]
	if !ok {
		log.Fatalln("invalid description")
	}

	fmt.Println(`
	<html>
		<head></head>
		<body>
			<table>
				<tr>
					<th>Location</th>
					<th>Weather description</th>
				</tr>
	`)
	fmt.Println(`
		<tr>
			<td>` + d.location + `</td>
			<td>` + d.description + `</td>
		</tr>
	`)

	fmt.Println(`
			</table>
		</body>
	</html>
	`)

}
