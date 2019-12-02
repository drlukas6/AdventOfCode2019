package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func calculateFuel(module int) int {
	return (module / 3) - 2
}

func calculateTotalFuel(module int) int {
	fuelForModule := calculateFuel(module)
	if fuelForModule <= 0 {
		return 0
	}
	return fuelForModule + calculateTotalFuel(fuelForModule)
}

func main() {
	file, err := os.Open("./input.txt")

	if err != nil {
		log.Fatal(err)
	}

	massSum := 0
	massSumWithAdded := 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		mass, err := strconv.Atoi(scanner.Text())

		if err != nil {
			log.Fatal(err)
		}

		massSum += calculateFuel(mass)
		massSumWithAdded += calculateTotalFuel(mass)
	}

	fmt.Println("Total fuel requirement:", massSum)
	fmt.Println("Total fuel requirement considering added", massSumWithAdded)
}
