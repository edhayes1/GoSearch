package main

import (
		"fmt"
		"Utils"
		)

func getCitiesList(numCities int) []int {
    cities := make([]int, numCities)
    for i := range cities {
        cities[i] = i
    }
    return cities
}//no built in function :(

func heapPerm(tour []int, size int, permutations *[][]int){

	if (size == 1){
        t := make([]int, len(tour))       
        copy(t,tour)   
        fmt.Printf("address of tour %p  address of t %p \n", &tour[0], &t[0])
        *permutations = append(*permutations, t)
        return
    }

    for i := 0; i < size; i++{
        heapPerm(tour, size-1, permutations)
 
        if (size%2==1){
        	tour[i], tour[size-1] = tour[size-1], tour[i]
        } else{
            tour[0], tour[size-1] = tour[size-1], tour[0]
        }
    }
}

func main() {
	cities := Utils.ParseFile()
	n := len(cities)+1
	permutations := make([][]int ,0, 40320)

	heapPerm(getCitiesList(n), n, &permutations)

	fmt.Print("\n\n\n\n")
	fmt.Print(permutations)
}