package util

import "errors"

var (
	validSchedule = map[string]bool{
		"nocturno":   true,
		"vespertino": true,
		"matutino":   true,
		"rotativo":   true,
	}

	contractType = map[string]bool{
		"medio tiempo":    true,
		"tiempo completo": true,
		"practicante":     true,
		"temporal":        true,
		"proyecto":        true,
		"freelance":       true,
	}
)

func VerifySchedule(schedule string) error {
	if !validSchedule[schedule] {
		return errors.New("invalid schedule")
	}
	return nil
}
func VerifyContractType(contract string) error {
	if !contractType[contract] {
		return errors.New("invalid contract")
	}
	return nil
}
