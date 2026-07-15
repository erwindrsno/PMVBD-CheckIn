package attendee

import "strings"

func parseGender(gender string) Gender {
	// Use strings.ToLower to make the check case-insensitive
	switch strings.ToLower(gender) {
	case "male":
		return Male
	case "female":
		return Female
	default:
		// Defaulting to Female or handling an error is up to your preference
		return Male
	}
}
