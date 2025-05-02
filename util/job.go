package util

import "errors"

var validSchedule = map[string]bool{
	"nocturno":   true,
	"vespertino": true,
	"matutino":   true,
	"rotativo":   true,
}

func VerifySchedule(schedule string) error {
	if !validSchedule[schedule] {
		return errors.New("invalid schedule")
	}
	return nil
}
