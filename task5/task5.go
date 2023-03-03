/*
Я не успел написать работающее решение, намудрил с кодом.
Поэтому просто опишу алгоритм текстом, вдруг кто на него взглянет, вроде рабочий.

1. Сначала ищем первый аэропорт с вылетами одной четности (или без вылетов), запрещаем оттуда вылететь, это будет тупик
2. Проверяем ведущие в него маршруты и запрещаем вылеты в дни другой четности, чтобы вылететь можно было в тупик.
3. Если из этого аэропорта (который ведет в тупик) нет других вылетов в день с этой четностью, тогда идем дальше и проверяем уже ведущие в этот другие аэропорты в рекурсии с тем же правилом
4. Если дошли до первого аэропорта и он ведет в тупик, выводим -1, получилось не дать Ксюшей доехать до N
5. Если есть обходные пути, ищем другой тупик подальше от старта и повторяем попытку завести в него из 1 аэропорта
6. Если все тупики можно будет обойти, или их нет, то посчитать длинный маршрут можно будет жадным алгоритмом, просто выбирая в качестве следующего узла ближайший аэропорт, таким образом максимизирую путь
7. Выводим длину максимального пути, с учетом обхода тупиков (Ксюша может и обойдет их)
*/

package main

import (
	"fmt"
)

func main() {

	var N, M Name
	// var Path []Name
	// Path = append(Path, Name(1))
	// var PathParityMode []bool

	// curAirport := 1
	// prevAirport
	// fmt.Scanln(&N, &M)
	N, M = 3, 3 // -1
	// N, M = 4, 6 // 3

	// var flights flights
	flights := make(Flights, M)

	flights[0] = Flight{1, 2, false}
	flights[1] = Flight{1, 3, true}
	flights[2] = Flight{2, 3, false}

	// flights[0] = Flight{1, 3, false}
	// flights[1] = Flight{3, 4, false}
	// flights[2] = Flight{3, 4, true}
	// flights[3] = Flight{1, 2, true}
	// flights[4] = Flight{2, 3, true}
	// flights[5] = Flight{2, 4, false}

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
		fmt.Printf("\t%s\t%v", "parity", v.parity)

	}
	fmt.Printf("\n\n\n%v\n\n\n", findMaxPath(airports, N))

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
		fmt.Printf("\t%s\t%v", "parity", v.parity)

	}
	fmt.Printf("\n\n\n%v\n\n\n", findMaxPath(airports2, N))

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

var (
	Path     []Name
	deadEnds []Name
	g        Graph
)

type Graph map[Name]Airport
type Flights []Flight
type Name uint32
type Cost uint32
type Flight struct {
	from   Name
	to     Name
	parity bool
}

func initAirports(flights Flights, N Name) Graph {
	airports := make(Graph)
	Path = make([]Name, 0)
	for i := Name(1); i <= N; i++ {
		// fmt.Println("initAirports for1 iteration")
		if entry, _ := airports[i]; true {
			entry.destinatedFrom = make(map[Name]bool)
			airports[i] = entry
		}
	}

	for _, v := range flights {

		if entry, _ := airports[v.to]; true {
			entry.destinatedFrom[v.from] = v.parity
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

func findDeadEnds(N Name) bool {

	for i := Name(1); i < N; i++ {
		var AirportParityMode bool
		if (len(g[i].evenFlights) == 0) || (len(g[i].oddFlights) == 0) {
			deadEnds = append(deadEnds, i)
			if len(g[i].evenFlights) == 0 {
				AirportParityMode = true
			}
			if len(g[i].oddFlights) == 0 {
				AirportParityMode = false
			}
			if entry, _ := g[i]; true {
				entry.parity = AirportParityMode
				g[i] = entry
			}
			if closeTheFlights(i) {
				return true
			}
		}
	}
	return false
}

func closeTheFlights(Node Name) bool {
	var AirportParityMode bool
	if Node == 1 {
		return true
	}

	for dest, flightParity := range g[Node].destinatedFrom {

		AirportParityMode = flightParity
		if entry, _ := g[dest]; true {
			entry.parity = AirportParityMode
			g[dest] = entry
		}
		if AirportParityMode && len(g[dest].oddFlights) > 1 {
			return false
		} else if !AirportParityMode && len(g[dest].evenFlights) > 1 {
			return false
		}
		return closeTheFlights(dest)
	}
	return false
}

func findMaxPath(airports Graph, N Name) []Name {

	g = airports

	Path := make([]Name, 0)
	Path = append(Path, 1)

	if findDeadEnds(N) {
		Path = append(Path, 666666666)
		return Path
		fmt.Println("66666666")
	}

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
			// if entry, _ := airports[nextAirport]; true {
			// 	entry.parity = false
			// 	airports[1] = entry
			// }
		} else {
			if entry, _ := airports[1]; true {
				entry.parity = true
				airports[1] = entry
			}
			nextAirport = nearestEven
			// if entry, _ := airports[nextAirport]; true {
			// 	entry.parity = true
			// 	airports[1] = entry
			// }
		}
		Path = append(Path, nextAirport)
	} else {
		nextAirport = 0
		return Path
	}
	findMaxPathRecurcive(&airports, nextAirport, N)
	return Path
}

func findMaxPathRecurcive(g *Graph, nextAirport Name, N Name) []Name {
	if nextAirport == N {
		return Path
	}
	i := nextAirport
	airports := *g

	nearestEven := Name(2147483647)
	nearestOdd := Name(2147483647)

	if entry, _ := airports[i]; true {
		entry.costToReach = 0
		airports[i] = entry
	}

	for _, v := range airports[i].evenFlights {
		if v.to < nearestEven {
			nearestEven = v.to
		}
	}
	for _, v := range airports[i].oddFlights {
		if v.to < nearestOdd {
			nearestOdd = v.to
		}
	}
	if (nearestEven != Name(2147483647)) || (nearestOdd != Name(2147483647)) {
		if nearestEven > nearestOdd {
			if entry, _ := airports[i]; true {
				entry.parity = false
				airports[i] = entry
			}
			nextAirport = nearestOdd
			// if entry, _ := airports[nextAirport]; true {
			// 	entry.parity = false
			// 	airports[i] = entry
			// }
		} else {
			if entry, _ := airports[1]; true {
				entry.parity = true
				airports[i] = entry
			}
			nextAirport = nearestEven
			// if entry, _ := airports[nextAirport]; true {
			// 	entry.parity = true
			// 	airports[i] = entry
			// }
		}
		Path = append(Path, nextAirport)
	} else {
		nextAirport = 0
		return Path
	}

	// nearest := Name(2147483647)

	// if airports[i].parity {
	// 	for _, v := range airports[Name(i)].evenFlights {
	// 		if v.to < nearest {
	// 			nearest = v.to
	// 		}
	// 	}
	// } else {
	// 	for _, v := range airports[Name(i)].oddFlights {
	// 		if v.to < nearest {
	// 			nearest = v.to
	// 		}
	// 	}
	// }

	// if nearest != Name(2147483647) {
	// 	if entry, _ := airports[1]; true {
	// 		entry.parity = false
	// 		airports[1] = entry
	// 	}
	// 	nextAirport = nearest
	// 	Path = append(Path, nextAirport)

	// } else {
	// 	nextAirport = 0
	// 	return Path
	// }

	findMaxPathRecurcive(&airports, nextAirport, N)
	return Path
}
