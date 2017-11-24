package Utils

import (
		"fmt"
		"io/ioutil"
		"regexp"
		"strconv"
		"os"
		"strings"
		"math/rand"
)


func ParseFile(filename string) ([][]int){
    b, err := ioutil.ReadFile(filename) // just pass the file name
    if err != nil {
        fmt.Print(err)
    }

    str := string(b) // convert content to a 'string'
	str = strings.Replace(str, "\n", "", -1)
	str = strings.Replace(str, "\r", "", -1)
	str = strings.Replace(str, " ", "", -1)

	regex := regexp.MustCompile(`[0-9]+`)
	cities_UnForm := regex.FindAllString(str, -1)

	cities_UnFormInts := make([]int, len(cities_UnForm))

	for index, i := range cities_UnForm{
		cities_UnFormInts[index], err = strconv.Atoi(i)
	}

	x := cities_UnFormInts[0]

	cities := make([][]int, x-1)
	index := 0
	if (filename == "/home/ed/Documents/SoftwareMethodologies/AISearch/Search/CityFiles/AISearchtestcase.txt"){
		index = 1
	} else {
		index = 2
	}				//huge hack im actually disgusted in myself


	for city := 0; city < cities_UnFormInts[0]-1; city++{
		x--
		zeros := make([]int, cities_UnFormInts[0]-x)
		cities[city] = append(zeros, cities_UnFormInts[index:index+x]...)
		index += x
	}

    return cities
}

func fileError(e error) {
	if e != nil {
		panic(e)
	}
}

func WriteFile(tourSize int, length int, tour []int){
	for i, _ := range tour{
		tour[i]++
	}

	tourSizeString := "TOURSIZE = " + strconv.Itoa(tourSize) + ",\n"
	tourSizeStringZero := strconv.Itoa(tourSize)
	if (len(tourSizeStringZero) < 3) {tourSizeStringZero = strconv.Itoa(0) + tourSizeStringZero}


	nameString := 	"NAME = AISearchfile" + tourSizeStringZero + ",\n"
	lengthString := "LENGTH = " + strconv.Itoa(length) + ",\n"

	f, err := os.Create("/home/ed/Documents/SoftwareMethodologies/AISearch/GoSearch/main/answers/TourfileB/tourAISearchfile" + tourSizeStringZero + ".txt")
	fileError(err)

	fmt.Fprint(f, nameString)
	fmt.Fprint(f, tourSizeString)
	fmt.Fprint(f, lengthString)

	for i := 0; i < len(tour)-1; i++{
		fmt.Fprint(f, tour[i])
		fmt.Fprint(f, ",")
	}
	fmt.Fprint(f, tour[len(tour)-1])

	f.Sync()
	f.Close()
}

func GetTourLength(tour []int, cities [][]int) (int){
	distance := 0

	for i := 0; i < len(tour)-1; i++{

		x := tour[i]
		y := tour[i+1]

		if y > x{
			x, y = y ,x
		}
		
		distance += cities[y][x]
	}

	x := tour[0]
	y := tour[len(tour)-1]

	if y > x{
		x, y = y ,x
	}

	distance += cities[y][x]

	return distance
}

func FindBestTour(tours [][]int, cities [][]int) ([]int, int){

	shortestLength := GetTourLength(tours[0], cities)
	bestTour := tours[0]

	for _, tour := range tours{
		length := GetTourLength(tour, cities)

		if length < shortestLength{
			shortestLength = length
			bestTour = tour
		}
	}

	return bestTour, shortestLength
}

func GetCitiesList(numCities int) []int {
    cities := make([]int, numCities)
    for i := range cities {
        cities[i] = i
    }
    return cities
}//no built in function :(

func CityInSlice(city1 int, slice []int) bool {
	for _, b := range slice {
		if b == city1 {
			return true
		}
	}
	return false
}

func GetIndexOfValue(currCity int, parent []int) int{

	for j, city := range parent{
		if city == currCity{
			return j
		}
	}
	return -1
}

func GenRandomTours(numTours int, size int)([][]int){
	randomTours := make([][]int, numTours)

	for i, tour := range randomTours{
		tour = rand.Perm(size)
		randomTours[i] = tour
	}

	return randomTours
}