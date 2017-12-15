package main

import (
	"fmt"
	"Utils"
	"math/rand"
	"math"
)

func invertSubTour (tour []int) []int{
	maxLengthOfSubTour := len(tour)-1
	//if maxLengthOfSubTour == 0{maxLengthOfSubTour = 2}

	lengthOfSubTour := rand.Intn(maxLengthOfSubTour)
	startIndex := rand.Intn(len(tour)-1 - lengthOfSubTour)
	endIndex := startIndex + lengthOfSubTour -1

	for startIndex <= endIndex{
		tour[startIndex], tour[endIndex] = tour[endIndex], tour[startIndex]
		startIndex ++
		endIndex --
	}

	return tour
}

func probFunc(deltaE int, temp float64) bool{
	sample := rand.Float64()
	if(sample < math.Exp(-math.Abs(float64(deltaE))/temp)) {
		return true
	} else{
		return false
	}
}

func createNullTour(size int) []int {
	tour := make([]int, size)
	for i,_ := range tour{
		tour[i] = -1
	}
	return tour
}

func nearestNeighbour(cities [][]int, size int) []int{
	tour := createNullTour(size)
	currentCity := rand.Intn(size)

	for i := 0; i < size; i++{
		minDistance := math.MaxInt32
		closestCity := -1

		for j := 0; j < size; j++ {

			if (j != currentCity && !Utils.CityInSlice(j, tour)){
				x := currentCity
				y := j

				if y > x {
					x, y = y, x
				}

				distance := cities[y][x]
				if (distance < minDistance){
					minDistance = distance
					closestCity = j
				}
			}
		}
		currentCity = closestCity
		tour[i] = currentCity
	}
	return tour
}

func expCool(numIterations int, T0 float64, a float64) float64{
	return T0*math.Pow(a, float64(numIterations))
}

func logCool(numIterations int, T0 float64, a float64) float64{
	return T0/(1 + a*math.Log10(float64(1+numIterations)))
}

func linCool(numIterations int, T0 float64, a float64) float64{
	return T0/(1.0 + a*float64(numIterations))
}

func quadCool(numIterations int, T0 float64, a float64) float64{
	return T0/(1.0+ a*math.Pow(float64(numIterations), 2))
}

func getNeighbourhood(tour []int) [][]int{

	neighbourhood := make([][]int, 5)
	for i := 0; i < 5; i++ {
		temp := make([]int, len(tour))
		copy(temp, tour)
		neighbourhood[i] = invertSubTour(temp)
	}

	return neighbourhood
}

func swapCities(tour []int) []int{
	city1 := rand.Intn(len(tour))
	city2 := rand.Intn(len(tour))

	for city1 == city2{
		city2 = rand.Intn(len(tour))
	}

	tour[city1], tour[city2] = tour[city2], tour[city1]

	return tour
}


func main(){
	filename := "/home/ed/Documents/SoftwareMethodologies/AISearch/Search/CityFiles/AISearchfile180.txt"
	cities := Utils.ParseFile(filename)
	size := len(cities)+1

	temperature := 200.0
	a := 0.999999

	currentTour := nearestNeighbour(cities, size)

	numIterations := 0

	bestTour := make([]int, size)
	bestLength := math.MaxInt32

	for temperature > 0.0000001 && bestLength > 1950{

		//temp := make([]int, len(currentTour))
		//copy(temp, currentTour)
		//successorTour := swapCities(temp)
		neighbourhood := getNeighbourhood(currentTour)
		successorTour, lenSuccessorTour := Utils.FindBestTour(neighbourhood, cities)

		//lenSuccessorTour := Utils.GetTourLength(successorTour, cities)
		lenCurrentTour := Utils.GetTourLength(currentTour, cities)

		deltaE := lenSuccessorTour - lenCurrentTour

		if (deltaE <= 0 || probFunc(deltaE, temperature)){
			currentTour = successorTour
			lenCurrentTour = lenSuccessorTour
		}

		if (lenCurrentTour < bestLength){
			bestLength = lenCurrentTour
			copy(bestTour, currentTour)
			fmt.Print(bestLength)
			fmt.Print(" temp = ")
			fmt.Print(temperature)
			fmt.Print("\n\n")
		}

		numIterations++
		temperature = a*temperature//expCool(numIterations, temperature, a)

	}

	Utils.WriteFile(size, Utils.GetTourLength(bestTour, cities), bestTour)
}