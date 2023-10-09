package lasagna

// TODO: define the 'PreparationTime()' function
func PreparationTime(layers []string, timePerLayer int) int {
	realTimePerLayer := timePerLayer

	if timePerLayer == 0 {
		realTimePerLayer = 2
	}

	return len(layers) * realTimePerLayer // len(layers) * timePerLayer > 0 ? timePerLayer : 2
}

// TODO: define the 'Quantities()' function
func Quantities(layers []string) (noodles int, sauce float64) {

	for i := 0; i < len(layers); i++ {
		if layers[i] == "sauce" {
			sauce += 0.2
		} else if layers[i] == "noodles" {
			noodles += 50
		}
	}

	return
}

// TODO: define the 'AddSecretIngredient()' function
func AddSecretIngredient(friendIngList []string, myIngList []string) {
	myIngList[len(myIngList)-1] = friendIngList[len(friendIngList)-1]
}

// TODO: define the 'ScaleRecipe()' function
func ScaleRecipe(quantities []float64, portionsToCook int) []float64 {
	quatitiesForPortions := []float64{}

	for i := 0; i < len(quantities); i++ {
		quatitiesForPortions = append(quatitiesForPortions, quantities[i]*(float64(portionsToCook)/2))
	}

	return quatitiesForPortions
}

// Your first steps could be to read through the tasks, and create
// these functions with their correct parameter lists and return types.
// The function body only needs to contain `panic("")`.
//
// This will make the tests compile, but they will fail.
// You can then implement the function logic one by one and see
// an increasing number of tests passing as you implement more
// functionality.
