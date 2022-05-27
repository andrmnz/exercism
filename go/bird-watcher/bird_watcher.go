package birdwatcher

// TotalBirdCount return the total bird count by summing
// the individual day's counts.
func TotalBirdCount(birdsPerDay []int) int {
	var totalBirdCount int
	for i := 0; i < len(birdsPerDay); i++ {
		totalBirdCount += birdsPerDay[i]
	}

	return totalBirdCount
}

// BirdsInWeek returns the total bird count by summing
// only the items belonging to the given week.
func BirdsInWeek(birdsPerDay []int, week int) int {
	var birdsInWeek int
	lastDayInWeek := week * 7
	firstDayInWeek := lastDayInWeek - 7
	birdsPerWeek := birdsPerDay[firstDayInWeek:lastDayInWeek]
	for i := 0; i < len(birdsPerWeek); i++ {
		birdsInWeek += birdsPerWeek[i]
	}
	return birdsInWeek
}

// FixBirdCountLog returns the bird counts after correcting
// the bird counts for alternate days.
func FixBirdCountLog(birdsPerDay []int) []int {
	for i := 0; i < len(birdsPerDay); i += 2 {
		birdsPerDay[i]++
	}
	return birdsPerDay
}
