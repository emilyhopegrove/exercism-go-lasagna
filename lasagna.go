package lasagna

//create a constant variable for the amount of time it takes to cook a lasagna
const OvenTime = 40

func RemainingOvenTime(actual int) int {
	//the remaining time in the oven is figured out by taking in how much time has actually passed as an argument 
	//and subtracting it from the constant oven time
    return OvenTime - actual
}

func PreparationTime(numberOfLayers int) int {
	//the preparation time is figured out by multiplying the number of layers as passed in as an argument by 2
    return numberOfLayers * 2
}

func ElapsedTime(numberOfLayers, actualMinutesInOven int) int {
	/* This function figures out the total time that has elabsed
	 by adding the preparation time and the number of minutes 
	 the lasagna has been in the oven and returning the total as an integer */
	 
    preparationTime := PreparationTime(numberOfLayers)
    return preparationTime + actualMinutesInOven
}


