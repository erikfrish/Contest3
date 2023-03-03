package main

import (
	"fmt"
	"math"
)

// type flights []Flight

func main() {
	// fmt.Println("kek")

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
	flights := make(flights, M)
	airports := make([]Airport, N)

	// flights[0] = Flight{1, 2, false}
	// flights[1] = Flight{1, 3, true}
	// flights[2] = Flight{2, 3, false}

	flights[0] = Flight{1, 3, false}
	flights[1] = Flight{3, 4, false}
	flights[2] = Flight{3, 4, true}
	flights[3] = Flight{1, 2, true}
	flights[4] = Flight{2, 3, true}
	flights[5] = Flight{2, 4, false}

	// fmt.Println(39)
	airports = initAirports(flights, N)
	for _, v := range airports {
		fmt.Println(v.destinationalAirports)
	}

	Graph := make(map[Name]Neighbors)
	for i := range airports {
		Graph[Name(i+1)] = airports[i].destinationalAirports
	}

	fmt.Println(GetPath(Graph, 1))

	TestGraph := make(map[Name]Neighbors)

	TestGraph = map[Name]Neighbors{
		1: {
			2: {
				1,
				false,
			},
			3: {
				1,
				false,
			},
		},

		2: {
			3: {
				1,
				false,
			},
			4: {
				1,
				false,
			},
		},

		3: {
			4: {
				1,
				false,
			},
		},

		4: {
			5: {
				1,
				false,
			},
			7: {
				1,
				false,
			},
		},

		5: {
			6: {
				1,
				false,
			},
		},

		6: {
			7: {
				1,
				false,
			},
		},

		7: {},
	}
	fmt.Println(GetPath(TestGraph, 1))

	// fmt.Println(airports)

	//

	//

	// S := make(map[uint32]bool)

	// for {

	// 	notEffFlight := airports[curAirport-1].flights.findLessEffectiveWay()
	// 	Path = append(Path, notEffFlight.to)
	// 	PathParityMode = append(PathParityMode, notEffFlight.parity)
	// 	// fmt.Println(airports[curAirport-1].flights.findLessEffectiveWay())
	// 	// fmt.Println(Path)
	// 	// fmt.Println(PathParityMode)

	// 	curAirport = int(notEffFlight.to)
	// 	if curAirport >= int(N) {
	// 		break
	// 	}

	// }

}

type Flight struct {
	from   Name
	to     Name
	parity bool
}

type flights []Flight

// func (f *flights) findLessEffectiveWay() Flight {
// 	var lessEffectiveWayTo uint32 = 4294967295
// 	var lessEffectiveWay Flight

// 	for _, v := range *f {
// 		if lessEffectiveWayTo > v.to {
// 			lessEffectiveWayTo = v.to
// 			lessEffectiveWay = v
// 		}
// 	}
// 	return lessEffectiveWay
// }

type Airport struct {
	flights               flights
	destinationalAirports Neighbors
	// parity    bool
	// wayLength uint32
}
type Name uint32

type Cost uint32

type Edge struct {
	Cost   Cost
	parity bool
}

type Neighbors map[Name]Edge

func initAirports(flights flights, N uint32) []Airport {
	airports := make([]Airport, N)
	for _, v := range flights {
		airports[v.from-1].flights = append(airports[v.from-1].flights, v)
	}
	var Cost Cost = 1
	for i := range airports {
		airports[i].destinationalAirports = make(Neighbors)
		// fmt.Println(v, Cost)
		for _, fl := range airports[i].flights {
			airports[i].destinationalAirports[fl.to] = Edge{Cost: Cost, parity: fl.parity}
		}
	}

	return airports

}

var (
	processed      []Name
	resultCosts    Result
	data           map[Name]Neighbors
	highestCostTmp HighestCost
	startNode      Name
)

type Result map[Name]Edge

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
	setHighestCostTmp()
	processNodeRecursive(&highestCostTmp.Name)

}

func getCostToCurrentNode(nodeName *Name) (c Cost) {
	if *nodeName == startNode {
		return 0
	}
	return resultCosts[*nodeName].Cost
}

func setHighestCostTmp() {
	prevHighestCostTmp := highestCostTmp
	resetHighestCostTmp()
	for name, edge := range resultCosts {
		cost := edge.Cost
		if !isNodeProcessed(&name) {
			if edge.Cost > highestCostTmp.Cost {
				if prevHighestCostTmp.Cost != getZeroCost() {
					highestCostTmp.Cost = cost + prevHighestCostTmp.Cost
				} else {
					highestCostTmp.Cost = cost
				}
				highestCostTmp.Name = name
			}
		}
	}
	return
}

func resetHighestCostTmp() {
	highestCostTmp.Cost = getZeroCost()
	highestCostTmp.Name = 0
}

func getInf() Cost {
	return math.MaxUint8
}
func getZeroCost() Cost {
	return 0
}

func isNodeProcessed(nodeName *Name) bool {
	for _, v := range processed {
		if v == *nodeName {
			return true
		}
	}
	return false
}
func setResultCosts(nodeName *Name) {
	c := getCostToCurrentNode(nodeName)
	neighbors := data[*nodeName]
	for neighborName, neighborEdge := range neighbors {
		if !isNodeProcessed(&neighborName) {
			if (c + neighborEdge.Cost) > resultCosts[neighborName].Cost {
				resultCosts[neighborName] = Edge{Cost: c + neighborEdge.Cost, parity: resultCosts[neighborName].parity}
			}
		}
	}
}

func GetPath(d map[Name]Neighbors, startNode Name) Result {
	initData(d, startNode)
	processNodeRecursive(&startNode)
	return resultCosts
}

func initData(d map[Name]Neighbors, startNode Name) {
	setStartNode(startNode)
	setData(d)
	resetProcessed()
	resetResultCosts()
	resetHighestCostTmp()
}

func setStartNode(s Name) {
	startNode = s
}

func resetProcessed() {
	processed = make([]Name, 0)
}

func setData(d map[Name]Neighbors) {
	data = d
}
func resetResultCosts() {
	resultCosts = make(map[Name]Edge)
	for node := range data {
		resultCosts[node] = Edge{getZeroCost(), false}
	}
}
