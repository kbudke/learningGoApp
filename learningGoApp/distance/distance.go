package distance

import "fmt"

const metersToYards float64 = 1.09361

func Meters() {
	var meters float64
	fmt.Print("Eneter the number of meters: ")
	fmt.Scan(&meters)
	yards := meters * metersToYards
	fmt.Println(meters, " meters is ", yards, " yards.")
}
