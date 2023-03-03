package main

import (
	"fmt"
)

func main() {

	var N, M uint32
	var Path []uint32
	Path = append(Path, 1)
	// var PathParityMode []bool

	// curAirport := 1
	// prevAirport
	// fmt.Scanln(&N, &M)
	// N, M = 3, 3 // -1
	N, M = 4, 6 // 3

	// var flights flights
	flights := make(Flights, M)
	airports := make(map[Name]Airport)

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

	// fmt.Println(39)
	airports = initAirports(flights, N)
	for name, v := range airports {
		fmt.Printf("\n%v\n", name)
		// fmt.Printf("\t%s\t%v", "EvenFl", v.evenFlights)
		// fmt.Printf("\t%s\t%v", "OddFl", v.oddFlights)

		fmt.Printf("\t%s\t%v", "oddNeighbors", v.oddNeighbors)
		fmt.Printf("\t%s\t%v", "evenNeighbors", v.evenNeighbors)

		fmt.Printf("\t%s\t%v", "destinatedFrom", v.destinatedFrom)

	}

	fmt.Printf("\nresult = %v\n", GetPath(airports))
	fmt.Printf("\nresultcosts = %v\n", resultCosts)

	TestGraph1 := Graph{
		1: {
			evenNeighbors: Neighbors{
				2: 1,
			},
			oddNeighbors: Neighbors{
				3: 1,
			},
			// evenFlights: {
			// 	1,
			// 	false,
			// },
			// oddFlights: {
			// 	1,
			// 	false,
			// },
		},

		2: {
			evenNeighbors: Neighbors{
				3: 1,
			},
			oddNeighbors: Neighbors{
				4: 1,
			},
		},

		3: {
			evenNeighbors: Neighbors{
				4: 1,
			},
			oddNeighbors: Neighbors{
				// 4 : 1,
			},
		},

		4: {
			evenNeighbors: Neighbors{
				5: 1,
			},
			oddNeighbors: Neighbors{
				7: 1,
			},
		},

		5: {
			evenNeighbors: Neighbors{
				6: 1,
			},
			oddNeighbors: Neighbors{
				// 4 : 1,
			},
		},

		6: {
			evenNeighbors: Neighbors{
				// 3 : 1,
			},
			oddNeighbors: Neighbors{
				7: 1,
			},
		},

		7: {},
	}

	fmt.Println(GetPath(TestGraph1))

}

type Flight struct {
	from   Name
	to     Name
	parity bool
}

type Name uint32

type Flights []Flight

type Cost uint32

type Graph map[Name]Airport

type Neighbors map[Name]Cost

type Result Neighbors

type Airport struct {
	evenNeighbors Neighbors
	oddNeighbors  Neighbors

	destinatedFrom map[Name]bool

	oddFlights  Flights
	evenFlights Flights

	parity bool
}

func initAirports(flights Flights, N uint32) Graph {
	airports := make(map[Name]Airport)

	for i := Name(1); uint32(i) <= N; i++ {
		fmt.Println("initAirports for1 iteration")
		if entry, _ := airports[i]; true {
			entry.evenNeighbors = make(Neighbors)
			entry.oddNeighbors = make(Neighbors)
			entry.destinatedFrom = make(map[Name]bool)
			airports[i] = entry
		}
	}

	for _, v := range flights {
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
	// count := 0
	for i := range airports {
		// count++
		// fmt.Println("initAirports for3 iteration")

		for j := range airports[i].evenFlights {
			airports[i].evenNeighbors[airports[i].evenFlights[j].to] = 1
		}
		for j := range airports[i].oddFlights {
			airports[i].oddNeighbors[airports[i].oddFlights[j].to] = 1
		}

		for j := range airports[i].evenNeighbors {
			airports[j].destinatedFrom[i] = true
		}

		for j := range airports[i].oddNeighbors {
			airports[j].destinatedFrom[i] = true
		}

	}
	// fmt.Println(count)

	// for i := range airports {
	// 	if entry, _ := airports[i]; true {
	// 		entry.Name = i
	// 		airports[i] = entry
	// 	}
	// }

	return airports

}

var (
	processed      []Name
	resultCosts    Result
	data           Graph
	highestCostTmp HighestCost
	startNode      Name
)

func GetPath(airports Graph) Result {
	startNode = 1
	initData(airports, startNode)
	processNodeRecursive(&startNode)

	return resultCosts
}

type HighestCost struct {
	Name Name
	Cost Cost
}

func processNodeRecursive(nodeName *Name) {
	if isNodeProcessed(nodeName) {
		return
	}
	processed = append(processed, *nodeName)
	setResultCosts(nodeName)
	// recheckResultCosts(nodeName)
	setHighestCostTmp()
	processNodeRecursive(&highestCostTmp.Name)

}

// func recheckResultCosts(nodeName *Name) {
// 	// c := getCostToCurrentNode(nodeName)
// 	for destFr := range data[*nodeName].destinatedFrom {

// 		fmt.Printf("\n||||%v\t||||%v||||", *nodeName, resultCosts[*nodeName])
// 		// for nameDestFr := range data[destFr].evenNeighbors {
// 		if resultCosts[destFr]+1 > resultCosts[*nodeName] {
// 			resultCosts[*nodeName] = resultCosts[destFr] + 1
// 		}
// 		// }
// 		// for j := range data[i].oddFlights {

// 		// }
// 	}
// }

func setResultCosts(nodeName *Name) {
	c := getCostToCurrentNode(nodeName)
	neighbors := data[*nodeName]
	// count := 0
	for neighborName, neighborCost := range neighbors.evenNeighbors {
		// count++
		// fmt.Println("setResultCosts for1 iteration")

		if !isNodeProcessed(&neighborName) {
			if (c + neighborCost) > resultCosts[neighborName] {
				resultCosts[neighborName] = c + neighborCost
				fmt.Printf("\n||%v||\t||||%v||||", *nodeName, resultCosts[*nodeName])
			}
		}
	}

	// fmt.Println(count)
	// count = 0
	for neighborName, neighborCost := range neighbors.oddNeighbors {
		// count++
		// fmt.Println("setResultCosts for2 iteration")

		if !isNodeProcessed(&neighborName) {
			if (c + neighborCost) > resultCosts[neighborName] {
				resultCosts[neighborName] = c + neighborCost
				fmt.Printf("\n||%v||\t||||%v||||", *nodeName, resultCosts[*nodeName])
			}
		}
	}
	// fmt.Println(count)
}

func getCostToCurrentNode(nodeName *Name) (c Cost) {
	if *nodeName == startNode {
		return 0
	}
	return resultCosts[*nodeName]
}

func isNodeProcessed(nodeName *Name) bool {
	for _, v := range processed {
		if v == *nodeName {
			return true
		}
	}
	return false
}

func setHighestCostTmp() {
	prevHighestCostTmp := highestCostTmp
	resetHighestCostTmp()
	count := 0
	for name, cost := range resultCosts {
		count++
		fmt.Println("setHighestCostTmp for iteration")

		if !isNodeProcessed(&name) {
			if cost > highestCostTmp.Cost {
				if prevHighestCostTmp.Cost != 0 {
					highestCostTmp.Cost = cost + prevHighestCostTmp.Cost
				} else {
					highestCostTmp.Cost = cost
				}
				highestCostTmp.Name = name
			}
		}
	}
	fmt.Println(count)
	return
}

func initData(d Graph, startNode Name) {
	setStartNode(startNode)
	setData(d)
	resetProcessed()
	resetResultCosts()
	resetHighestCostTmp()
	// fmt.Printf("\nResetting\n%v\n%v\n%v\n", processed, resultCosts, highestCostTmp)
}

func setStartNode(s Name) {
	startNode = s
}
func setData(d Graph) {
	data = d
}
func resetProcessed() {
	processed = make([]Name, 0)
}
func resetHighestCostTmp() {
	highestCostTmp.Cost = 0
	highestCostTmp.Name = 0
}
func resetResultCosts() {
	resultCosts = make(Result)
	for node := range data {
		resultCosts[node] = 0
	}
}
