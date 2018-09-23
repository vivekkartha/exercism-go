package space

//Planet represents a planet's name
type Planet string

//Age returns the age in years for the given seconds and planet
func Age(secs float64, planet Planet) float64 {
	var age = 0.0
	switch planet {
	case "Earth":
		age = getAge(secs, 1)
	case "Mercury":
		age = getAge(secs, 0.2408467)
	case "Venus":
		age = getAge(secs, 0.61519726)
	case "Mars":
		age = getAge(secs, 1.8808158)
	case "Jupiter":
		age = getAge(secs, 11.862615)
	case "Saturn":
		age = getAge(secs, 29.447498)
	case "Uranus":
		age = getAge(secs, 84.016846)
	case "Neptune":
		age = getAge(secs, 164.79132)
	}
	return age
}

//getSeconds returns the seconds in the given orbitalTime(in EarthYears)
func getAge(ageInSecs float64, orbitalTime float64) float64 {
	var secsInYear = 365.25 * 24 * 60 * 60
	return ageInSecs / (orbitalTime * secsInYear)
}
