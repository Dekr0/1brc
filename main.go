package main

/*

Format <string: station name>;<double: measurement (one fractional digit)>

Calculates the min, mean, and max temperature value per weather station.

Emits the results on standard out, sorted alphabetically by station name.
The result values per station in the format <min>/<mean>/<max>, rounded to one
fraction.

*/

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

type Station struct {
    sum float64
    min float32
    max float32
    count int
}

var bufferSize = 2049

func Solution(filename string) error {
    file, err := os.Open(filename)
    if err != nil { 
        return fmt.Errorf("Failed to open %s: %s", filename, err.Error())
    }

    buffer := make([]byte, bufferSize + 1, bufferSize + 1)

    var line string
    var splits []string
    var stationName string
    var station *Station
    var temp float64
    var in bool
    stations := make(map[string]*Station)
    scanner := bufio.NewScanner(file)
    scanner.Buffer(buffer, bufferSize + 1)
    for scanner.Scan() {
        line = scanner.Text()
        splits = strings.Split(line, ";")

        stationName = splits[0]

        temp, err = strconv.ParseFloat(splits[1], 32)
        if err != nil {
            return fmt.Errorf("Failed to parse temperature %s at line %s: %s", splits[1], line, err.Error())
        }

        station, in = stations[stationName]
        if !in {
            stations[stationName] = &Station{
                temp,
                float32(temp), 
                float32(temp),
                1,
            }
        } else {
            station.sum += temp
            station.count += 1
            if station.min > float32(temp) {
                station.min = float32(temp)
            } else if station.max < float32(temp) {
                station.max = float32(temp)
            }
        }
    }

    stationNames := make([]*string, 0, len(stations))

    slices.SortStableFunc(stationNames, func(a *string, b *string) int {
        return strings.Compare(*a, *b)
    })

    for _, stationName := range stationNames {
        station, in := stations[*stationName]
        if !in {
            return fmt.Errorf("Station %s appears in %s but its data wasn't collected", *stationName, filename)
        }
        fmt.Printf("%s: %.1f/%.1f/%.1f\n", 
            *stationName,
            station.min,
            station.sum / float64(station.count),
            station.max,
        )
    }
    return nil
}
