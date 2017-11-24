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

func main(){
	filename := "/home/ed/Documents/SoftwareMethodologies/AISearch/Search/CityFiles/AISearchfile535.txt"
	cities := Utils.ParseFile(filename)
	size := len(cities)+1

	currentTourTemp := Utils.GenRandomTours(1, size)
	currentTour := currentTourTemp[0]

	temperature := 100000000.0

	bestLength := math.MaxInt32

	for temperature > 0.00000000000000000000000000000000000000000000001 {
		temp := make([]int, len(currentTour))
		copy(temp, currentTour)

		successorTour := invertSubTour(temp)

		lenSuccessorTour := Utils.GetTourLength(successorTour, cities)
		lenCurrentTour := Utils.GetTourLength(currentTour, cities)

		deltaE := lenSuccessorTour - lenCurrentTour

		if (deltaE <= 0 || probFunc(deltaE, temperature)){
			currentTour = successorTour
			lenCurrentTour = lenSuccessorTour
		}

		if (lenCurrentTour < bestLength){
			bestLength = lenCurrentTour
			fmt.Print(lenCurrentTour)
			fmt.Print(" temp = ")
			fmt.Print(temperature)
			fmt.Print("\n\n")

			writtenTour := make([]int, len(currentTour))
			copy(writtenTour, currentTour)

			Utils.WriteFile(size, lenCurrentTour, writtenTour)
		}

		temperature *= 0.9999
	}
}