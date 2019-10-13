/*
Define any integer to the noOfAirports constant
Provide same number of values in the canTravel array (defined in the main function)
Run !
*/

package main

import (
	"time"
	"math/rand"
	"fmt"
	"strconv"
	"strings"
	"math"
)

const (
	RadiusOfEarth float64 = 4000        // in miles
	noOfAirports int = 4                // no of random airports
)

//Initialize ample Airports
func initSampleAirports() ([noOfAirports]float64, [noOfAirports]float64) {
	var latitudes [noOfAirports]float64
	var longitudes [noOfAirports]float64

	for i := 0; i < noOfAirports; i++ {
		lat := (rand.Float64() * 180.0) - 90.0                 //generate random no between -90 to 90
		long := (rand.Float64() * 360.0) - 180.0               //generate random no between -180 to 180
		latitudes[i] = lat
		longitudes[i] = long
	}

	return latitudes, longitudes

}

//Print Sample generted Airport Coordinates
func PrintAirportCoordinates(latitudes, longitudes [noOfAirports]float64) {
	for i := 0; i < noOfAirports; i++ {
		fmt.Printf("latitude :%v , longitude :%v", latitudes[i], longitudes[i])
		fmt.Println()
	}
}

func main() {

	rand.Seed(time.Now().UTC().UnixNano())

	latitude, longitude := initSampleAirports()
	PrintAirportCoordinates(latitude, longitude)

	//Create a test case
	self := 0
	origin := 0
	destination := 1
	canTravel := []string{"2", "0 2", "0 3", "1"}            //length of this array should be equal to noOfAirports

	if len(canTravel) != noOfAirports {
		fmt.Println("Wrong Inputs !")
		return
	}

	//Find Shortest Trip
	shortestPath := shortestTrip(self, latitude, longitude, canTravel, origin, destination)
	if shortestPath == math.MaxFloat32 {
		shortestPath = -1
	}
	fmt.Println("Shortest Distance Travelled is :", shortestPath, "miles")
}

//Calculate shortest trip
func shortestTrip(self int, latitude [noOfAirports]float64, longitude[noOfAirports]float64, canTravel []string, origin int, destination int) float64 {
	fmt.Println("self :", self)
	fmt.Println("latitude :", latitude)
	fmt.Println("longitude :", longitude)
	fmt.Println("canTravel :", canTravel, len(canTravel))
	fmt.Println("origin :", origin)
	fmt.Println("destination :", destination)

	//handling base cases
	//if origin == destination {
	//	//no need to travel
	//	return 0.0
	//}
	//sourceDestSet := setOfAllPossibleDestinations(canTravel)
	//if _, ok := sourceDestSet[destination]; !ok {
	//	//the destination is unreachable
	//	return -1
	//}

	//Calculating the shortest distance
	var finalDistToTravel float64 = 0.0;

	//Check if direct travel to the destination is allowed
	isDirectTravelallowed := false
	allowedDestinationsFromOrigin := canTravel[origin]
	allowedDestinationFromThisIndex := strings.Split(allowedDestinationsFromOrigin, " ")
	for _, v := range allowedDestinationFromThisIndex {
		allowedDestination, err := strconv.Atoi(v)
		checkError(err)
		if allowedDestination == destination {
			isDirectTravelallowed = true
		}
	}

	//if direct travel is allowed , then calculate and return that distance else calculate minimum distance of all possible paths from source to destination
	if isDirectTravelallowed {
		finalDistToTravel = calculateDirectDist(latitude, longitude, origin, destination)
	} else {
		visitedDestinations := make(map[int]struct{})
		visitedDestinations[origin] = struct{}{}                                     //add origin to visited destination Set
		allowedDestinations := allowedDestinationsFromThisAirport(canTravel, origin) //calculate direct allowed destinations for the origin
		//store intermediate results to avoid redundant calculations
		memo := make(map[string]float64)                                             //key->string(origin+destination) , value->minimum distance between them
		finalDistToTravel = calculateDist(latitude, longitude, origin, destination, canTravel, visitedDestinations, allowedDestinations, memo)
	}

	return finalDistToTravel

}

//Calculating distance fro give origin to destination
func calculateDirectDist(latitude [noOfAirports]float64, longitude[noOfAirports]float64, origin int, destination int) float64 {
	lat1 := latitude[origin]
	lat2 := latitude[destination]
	long1 := longitude[origin]
	long2 := longitude[destination]

	result := RadiusOfEarth * math.Acos(math.Sin(lat1) * math.Sin(lat2) * math.Cos(lat1) * math.Cos(lat2) * math.Cos(long1 - long2))

	//Log to Check intermediate results
	fmt.Println("intermediate result from ", origin, " to ", destination, " : ", result)
	return result
}

//Calculate the shortest distance
func calculateDist(latitude [noOfAirports]float64, longitude[noOfAirports]float64, origin int, destination int, canTravel []string, visitedDestinations map[int]struct{}, allowedDestinations map[int]struct{}, memo map[string]float64) float64 {

	//var minDist float64
	if origin == destination {
		return 0.0
	}
	var minDist float64 = math.MaxFloat32

	//log to keep check of Allowed Destinations and Visited Destinations
	fmt.Println("allowed dest :", allowedDestinations)
	fmt.Println("visited dest :", visitedDestinations)

	for currDest := range allowedDestinations {
		if _, ok := visitedDestinations[currDest]; !ok {
			visitedDestinations[currDest] = struct{}{}
			allowedDestinations = allowedDestinationsFromThisAirport(canTravel, currDest)
			tempOrigin := strconv.Itoa(origin)
			tempDest := strconv.Itoa(origin)
			var newValue float64
			if val, ok := memo[tempOrigin + tempDest]; ok {
				newValue = val
			} else {
				newValue = calculateDirectDist(latitude, longitude, origin, currDest) + calculateDist(latitude, longitude, currDest, destination, canTravel, visitedDestinations, allowedDestinations, memo)
			}

			minDist = math.Min(minDist, newValue)
			memo[tempOrigin + tempDest] = minDist
		}
	}

	return minDist
}


/*
Helper functions
*/
//To get the destinations that can be travelled from this airport
func allowedDestinationsFromThisAirport(canTravel []string, airport int) map[int]struct{} {
	allowedDestinationsSet := make(map[int]struct{}, 0)

	allowedDestinationsString := canTravel[airport]
	allowedDestinationString := strings.Split(allowedDestinationsString, " ")
	for _, v := range allowedDestinationString {
		allowedDestinationVal, err := strconv.Atoi(v)
		checkError(err)
		allowedDestinationsSet[allowedDestinationVal] = struct{}{}
	}
	return allowedDestinationsSet
}

//To get the all the destinations that can be travelled (be it from any airport)
func setOfAllPossibleDestinations(canTravel []string) map[int]struct{} {
	sourceDestSet := make(map[int]struct{})
	for _, v := range canTravel {
		canTravelFromThisIndex := strings.Split(v, " ")
		for _, v1 := range canTravelFromThisIndex {
			intVal, err := strconv.Atoi(v1)
			checkError(err)
			sourceDestSet[intVal] = struct{}{}
		}
	}

	return sourceDestSet

}

//Error checking
func checkError(err error) {
	if err != nil {
		fmt.Println("Error occured :", err)
	}
}
