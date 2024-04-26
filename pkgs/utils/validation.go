package utils

import (
	"net/url"
	"regexp"
)

func ValidatePhoneNumber(mobileNo string) bool {
	pattern := `^\+\d{1,15}$`
	re := regexp.MustCompile(pattern)
	return re.MatchString(mobileNo)
}

func ValidateEmailAddress(email string) bool {
	pattern := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	re := regexp.MustCompile(pattern)
	return re.MatchString(email)
}

func ValidateWebstie(website string) bool {
	_, err := url.ParseRequestURI(website)
	return err == nil
}
