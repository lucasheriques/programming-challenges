package birdwatcher

// TotalBirdCount return the total bird count by summing
// the individual day's counts.
func TotalBirdCount(birdsPerDay []int) int {
	totalBirsdPerDay := 0

	for i := 0; i < len(birdsPerDay); i++ {
		totalBirsdPerDay += birdsPerDay[i]
	}

	return totalBirsdPerDay
}

// semana 1 -> 0, começar em 0, terminar no 6
// semana 2 -> 7, começar em 7, terminar no 13 [7, 8, 9, 10, 11, 12, 13]
// semana 3 -> 14

// BirdsInWeek returns the total bird count by summing
// only the items belonging to the given week.
func BirdsInWeek(birdsPerDay []int, week int) int {
	totalBirdsInWeek := 0

	i := 0
	if week > 1 {
		i = (week - 1) * 7
	}

	for ; i < week*7; i++ {
		totalBirdsInWeek += birdsPerDay[i]
	}

	return totalBirdsInWeek
}

/**
- a cada 2 dias o pássara volta
- ele chegou no primeiro dia

[]int{2, 5, 0, 7, 4, 1} -> len == 6

i = 2
{3, 6, 3}
*/

// FixBirdCountLog returns the bird counts after correcting
// the bird counts for alternate days.
func FixBirdCountLog(birdsPerDay []int) []int {

	for i := 0; i < len(birdsPerDay); i++ {
		if i%2 == 0 {
			birdsPerDay[i] += 1 // birdsPerDay[i] = birdsPerDay[i] + 1
		}
	}

	return birdsPerDay
}
