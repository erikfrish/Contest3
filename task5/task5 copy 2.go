package main

import (
	"fmt"
)

func main() {

	var N, M uint32
	// var Path []Name
	// Path = append(Path, Name(1))
	// var PathParityMode []bool

	// curAirport := 1
	// prevAirport
	// fmt.Scanln(&N, &M)
	// N, M = 3, 3 // -1
	N, M = 4, 6 // 3

	// var flights flights
	flights := make(Flights, M)

	// flights[0] = Flight{1, 2, false}
	// flights[1] = Flight{1, 3, true}
	// flights[2] = Flight{2, 3, false}

	flights[0] = Flight{1, 3, false}
	flights[1] = Flight{3, 4, false}
	flights[2] = Flight{3, 4, true}
	flights[3] = Flight{1, 2, true}
	flights[4] = Flight{2, 3, true}
	flights[5] = Flight{2, 4, false}

	// flights[0] = Flight{1, 3, false}
	// flights[1] = Flight{3, 4, false}
	// flights[2] = Flight{3, 4, true}
	// flights[3] = Flight{1, 2, true}
	// flights[4] = Flight{2, 3, true}
	// flights[5] = Flight{2, 4, false}

	airports := initAirports(flights, N)
	for name, v := range airports {
		fmt.Printf("\nairports %v\n", name)
		fmt.Printf("\t%s\t%v", "EvenFl", v.evenFlights)
		fmt.Printf("\t%s\t%v", "OddFl", v.oddFlights)
		fmt.Printf("\t%s\t%v", "destinated", v.destinatedFrom)

	}

	N, M = 7, 10
	flights2 := make(Flights, M)

	flights2[0] = Flight{1, 2, true}
	flights2[1] = Flight{1, 3, false}
	flights2[2] = Flight{2, 3, true}
	flights2[3] = Flight{2, 4, false}
	flights2[4] = Flight{3, 4, true}
	flights2[5] = Flight{3, 4, false}
	flights2[6] = Flight{4, 5, true}
	flights2[7] = Flight{4, 7, false}
	flights2[8] = Flight{5, 6, false}
	flights2[9] = Flight{6, 7, false}

	airports2 := initAirports(flights, N)
	for name, v := range airports2 {
		fmt.Printf("\nairports2 %v\n", name)
		fmt.Printf("\t%s\t%v", "EvenFl", v.evenFlights)
		fmt.Printf("\t%s\t%v", "OddFl", v.oddFlights)
		fmt.Printf("\t%s\t%v", "destinated", v.destinatedFrom)

	}

}

type Airport struct {

	// evenNeighbors    []Name
	// oddNeighbors     []Name

	oddFlights  Flights //массив
	evenFlights Flights //массив

	costToReach    Cost
	destinatedFrom map[Name]bool

	parity bool
}

type Graph map[Name]Airport
type Flights []Flight
type Name uint32
type Cost uint32
type Flight struct {
	from   Name
	to     Name
	parity bool
}

func initAirports(flights Flights, N uint32) Graph {
	airports := make(Graph)

	for i := Name(1); uint32(i) <= N; i++ {
		// fmt.Println("initAirports for1 iteration")
		if entry, _ := airports[i]; true {
			entry.destinatedFrom = make(map[Name]bool)
			airports[i] = entry
		}
	}

	for _, v := range flights {

		if entry, _ := airports[v.to]; true {
			entry.destinatedFrom[v.from] = true
			airports[v.to] = entry
		}
		// fmt.Println("initAirports for2 iteration")
		if v.parity {
			if entry, _ := airports[v.from]; true {
				entry.evenFlights = append(entry.evenFlights, v)
				airports[v.from] = entry
			}
		}
		if !v.parity {
			if entry, _ := airports[v.from]; true {
				entry.oddFlights = append(entry.oddFlights, v)
				airports[v.from] = entry
			}
		}

	}

	return airports

}

func findMaxPath(airports Graph, N Name) {
	Path := make([]Name, 0)
	Path = append(Path, 1)

	var nextAirport Name
	nearestEven := Name(2147483647)
	nearestOdd := Name(2147483647)

	if entry, _ := airports[1]; true {
		entry.costToReach = 0
		airports[1] = entry
	}

	for _, v := range airports[1].evenFlights {
		if v.to < nearestEven {
			nearestEven = v.to
		}
	}
	for _, v := range airports[1].oddFlights {
		if v.to < nearestOdd {
			nearestOdd = v.to
		}
	}
	if (nearestEven != Name(2147483647)) || (nearestOdd != Name(2147483647)) {
		if nearestEven > nearestOdd {
			if entry, _ := airports[1]; true {
				entry.parity = false
				airports[1] = entry
			}
			nextAirport = nearestOdd
			if entry, _ := airports[nextAirport]; true {
				entry.parity = false
				airports[1] = entry
			}
		} else {
			if entry, _ := airports[1]; true {
				entry.parity = true
				airports[1] = entry
			}
			nextAirport = nearestEven
			if entry, _ := airports[nextAirport]; true {
				entry.parity = true
				airports[1] = entry
			}
		}
	} else {
		nextAirport = 0
		return
	}

	for i := nextAirport; Name(i) <= N; i++ {
		nearestEven = Name(2147483647)
		nearestOdd = Name(2147483647)

		for _, v := range airports[Name(i)].evenFlights {
			if v.to < nearestEven {
				nearestEven = v.to
			}
		}
		for _, v := range airports[Name(i)].oddFlights {
			if v.to < nearestOdd {
				nearestOdd = v.to
			}
		}
		if (nearestEven != Name(2147483647)) || (nearestOdd != Name(2147483647)) {
			if nearestEven > nearestOdd {
				if entry, _ := airports[1]; true {
					entry.parity = false
					airports[1] = entry
				}
				nextAirport = nearestOdd
			} else {
				if entry, _ := airports[1]; true {
					entry.parity = true
					airports[1] = entry
				}
				nextAirport = nearestEven
			}
		} else {
			nextAirport = 0
			return
		}
	}

}
