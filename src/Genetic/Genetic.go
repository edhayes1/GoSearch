package main

import (
	"fmt"
	"Utils"
	"math/rand"
	"math"
)

func genRandomTours(numTours int, size int)([][]int){
	randomTours := make([][]int, numTours)

	for i, tour := range randomTours{
		tour = rand.Perm(size)
		randomTours[i] = tour
	}

	return randomTours
}

func sortnElements(population [][]int, cities [][]int, n int) ([][]int){
	for i := 0; i < n; i++ {
	    for j := i+1; j < n; j++ {
	        if(Utils.GetTourLength(population[j], cities) < Utils.GetTourLength(population[i], cities)){
	        	population[i], population[j] = population[j], population[i]
	        }
	    }
	}

	return population
}

func chooseParents(population [][]int, populationSize int) ([]int,[]int){
	P1 := population[rand.Intn(populationSize/5)]
	P2 := population[rand.Intn(populationSize/5)]

	return P1, P2
}

func roulette2(population [][]int, cities [][]int) ([]int,[]int){
	lengths := make ([]int, len(population))
	maxLength := math.MaxInt8

	for i, tour := range population{
		length := Utils.GetTourLength(tour, cities)

		if (length > maxLength){
			maxLength = length
		}
		lengths[i] = length
	}

	totalCumulativeProb := 0.0
	for j, _ := range population {
		totalCumulativeProb += math.Pow(float64(float64(maxLength - lengths[j]) / float64(maxLength)),4)
	}

	points := []float64{rand.Float64() * totalCumulativeProb, rand.Float64() * totalCumulativeProb}
	cumulativeProb := 0.0

	parents := make([][]int, 2)

	for i := 0; i < 2; i++ {
		cumulativeProb = 0.0
		for j, _ := range population {
			cumulativeProb += math.Pow(float64(float64(maxLength - lengths[j]) / float64(maxLength)),4)
			if (cumulativeProb >= points[i]) {
				parents[i] = population[j]
				break
			}
		}
	}

	return parents[0], parents[1]
}

func roulette(population [][]int, cities [][]int) ([]int,[]int){
	totalLength := 0
	lengths := make ([]int, len(population))
	for i, tour := range population{
		length := Utils.GetTourLength(tour, cities)
		lengths[i] = length
		totalLength += length
	}

	wheel := make ([]float64, len(population))
	totalWheelLength := 0.0
	for i, _ := range wheel{
		wheel[i] = math.Pow((float64(totalLength)/float64(lengths[i])),4)
		totalWheelLength += wheel[i]
	}

	choice1 := float64(rand.Intn(int(totalWheelLength)))
	choice2 := float64(rand.Intn(int(totalWheelLength)))

	i := -1

	P1 := -1
	for choice1 >= 0{
		i++
		choice1 -= wheel[i]
	}
	P1 = i

	P2 := -1
	i = -1

	for choice2 >= 0{
		i++
		choice2 -= wheel[i]
	}
	P2 = i

	return population[P1], population[P2]
}

//func doReplace(child []int) ([]int){
//
//	temp := make ([]int, len(child)) 				//init array with -1s
//	for i, _ := range temp{temp[i] = -1}
//
//	duplicates := make([]int, 0, len(child)/2)
//	missingCities := make([]int, 0, len(child)/2)
//
//	for _, city := range child{
//
//		if temp[city] == -1{
//			temp[city] = city
//		} else{
//			duplicates = append(duplicates, city)
//		}
//	}
//
//	for index, city := range temp{
//		if city == -1{
//			missingCities = append(missingCities, index)
//		}
//	}
//
//	for i, duplicate := range duplicates{
//		for j, city := range child{
//			if city == duplicate{
//				child[j] = missingCities[i]
//				break
//			}
//		}
//	}
//
//	return child
//}

//func doCrossover(parent1 []int, parent2 []int) ([]int, []int){
//	i := rand.Intn(len(parent1))
//
//	child1 := append(parent1[:i], parent2[i:]...)
//	child2 := append(parent1[i:], parent2[:i]...)
//
//	temp1 := make ([]int, len(child1))
//	temp2 := make ([]int, len(child1))
//	copy(temp1, child1)
//	copy(temp2, child2)
//
//	child1 = doReplace(temp1)
//	child2 = doReplace(temp2)
//
//	return child1, child2
//}

func getEdgeMap(parent []int, edgeMap [][] int) ([][]int){

   	for index, city := range parent {

		next := index + 1
		last := index - 1

		if index == len(parent)-1 {
			next = 0
		}
		if index == 0 {
			last = len(parent) - 1
		}

		if !Utils.CityInSlice(parent[last], edgeMap[city]) {
			edgeMap[city] = append(edgeMap[city], parent[last])
		}
		if !Utils.CityInSlice(parent[next], edgeMap[city]){
			edgeMap[city] = append(edgeMap[city], parent[next])
		}
	}
	return edgeMap
}

func chooseRandom(a int, b int) int{
	sample := rand.Intn(100)
	if sample > 50{return a}
	return b
}

func removeElement(x int, edgeMap [][]int){
	for i, listOfNeighbours := range edgeMap{
		for j, city := range listOfNeighbours{
			if city == x{
				edgeMap[i][j] = -1
			}
		}
	}
}

func chooseCityNotInChild(child []int, filledUpPoint int) int{
	fullList := Utils.GetCitiesList(len(child))

	for i := 0; i <= filledUpPoint; i++{
		fullList[child[i]] = -1
	}
	randomChoice := -1
	for randomChoice == -1{
		randomChoice = fullList[rand.Intn(len(fullList))]
	}
	return randomChoice
}

func doERCrossover(parent1 []int, parent2 []int) []int{
	//choose initial city
	child := make ([]int, len(parent1))

	edgeMap := make ([][]int, len(parent1))
	edgeMap = getEdgeMap(parent1, edgeMap)
	edgeMap = getEdgeMap(parent2, edgeMap)

	currCity := chooseRandom(parent1[0], parent2[0])
	nextCity := -1

	for i, _ := range child{
		child[i] = currCity

		if i == len(child)-1{break}

		removeElement(currCity, edgeMap)

		minNoNeighbours := math.MaxInt8
		minNeighbour := -1

		for _, neighbour := range edgeMap[currCity]{
			if neighbour != -1 {
				if len(edgeMap[neighbour]) < minNoNeighbours{
					minNoNeighbours = len(edgeMap[neighbour])
					minNeighbour = neighbour
				}
			}
		}

		if minNeighbour == -1{
			nextCity = chooseCityNotInChild(child, i)
		} else{
			nextCity = minNeighbour
		}

		currCity = nextCity
	}

	return child
}

func doInvMutation(child []int, probability int) []int{
	sample := rand.Intn(100)

	if sample < probability{

		maxLengthOfSubTour := len(child)-1
		//if maxLengthOfSubTour == 0{maxLengthOfSubTour = 2}

		lengthOfSubTour := rand.Intn(maxLengthOfSubTour)
		startIndex := rand.Intn(len(child)-1 - lengthOfSubTour)
		endIndex := startIndex + lengthOfSubTour -1

		for startIndex <= endIndex{
			child[startIndex], child[endIndex] = child[endIndex], child[startIndex]
			startIndex ++
			endIndex --
		}
	}

	return child
}

func doMutation(child []int, probability int) ([]int){
	sample := rand.Intn(100)
	if sample < probability{
		city1 := rand.Intn(len(child))
		city2 := rand.Intn(len(child))

		for city1 == city2{
			city2 = rand.Intn(len(child))
		}

		child[city1], child[city2] = child[city2], child[city1]
	}

	return child
}

func doIt(populationSize int, targetFitness int, cities [][]int, size int) ([]int, int){


	population := genRandomTours(populationSize, size)
	bestTour, populationFitness := Utils.FindBestTour(population, cities)

	shortestLength := populationFitness

	probabilityOfMutation := 10		//out of 100


	for populationFitness > targetFitness{

		newPop := make([][]int, 0, populationSize)

		for i := 0; i < populationSize; i++{
			P1, P2 := roulette2(population, cities)

			child := doERCrossover(P1, P2)

			child = doInvMutation(child, probabilityOfMutation)

			newPop = append(newPop, child)
		}

		population = newPop
		bestTour, populationFitness = Utils.FindBestTour(population, cities)

		//if (shortestLength < 6000 && probabilityOfMutation < 100){probabilityOfMutation++}
		if (populationFitness < shortestLength){
			shortestLength = populationFitness
			writtenTour := make([]int, len(bestTour))
			copy(writtenTour, bestTour)

			Utils.WriteFile(size, shortestLength, writtenTour)

			fmt.Print("\n")
			fmt.Print(populationFitness)
		}
	}

	return bestTour, populationFitness
}

func main() {

	cities := Utils.ParseFile()
	size := len(cities)+1

	bestTour, shortestLength := doIt(200, 1473, cities, size)
	print("reached goal ")
	print(bestTour)
	print(shortestLength)
}