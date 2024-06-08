package util

import (
	"regexp"
	"strconv"
)

func StrToInt(str string) int {
	value, err := strconv.Atoi(str)
	if err != nil {
		return 0
	}

	return value
}

func ValidPhoneNumber(phone string) bool {
	re := regexp.MustCompile(`^0[0-9]{9,12}$`)
	return re.MatchString(phone)
}

func ValidEmailAddress(email string) bool {
	emailRegex := regexp.MustCompile(`^[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,4}$`)
	return emailRegex.MatchString(email)
}

func ValidSwipeType(swipe int) bool {
	swipes := make(map[SwipeTypes]bool)
	swipes[Like] = true
	swipes[Pass] = true

	if _, ok := swipes[SwipeTypes(swipe)]; !ok {
		return false
	}
	return true
}
