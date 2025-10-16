package util

import (
	"errors"
	"regexp"
	"strings"
	"unicode"
)

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

var (
	nonAlnumDash = regexp.MustCompile(`[^a-z0-9-]`)
	multiDash    = regexp.MustCompile(`-+`)
)

func Slugify(s string) string {
	s = strings.ToLower(s)

	var b strings.Builder
	b.Grow(len(s))
	for _, r := range s {
		switch r {
		case 'á', 'à', 'ä', 'â', 'ã', 'å':
			r = 'a'
		case 'é', 'è', 'ë', 'ê':
			r = 'e'
		case 'í', 'ì', 'ï', 'î':
			r = 'i'
		case 'ó', 'ò', 'ö', 'ô', 'õ':
			r = 'o'
		case 'ú', 'ù', 'ü', 'û':
			r = 'u'
		case 'ñ':
			r = 'n'
		case 'ç':
			r = 'c'
		}
		// Reemplaza espacio y separadores Unicode por guion.
		if unicode.IsSpace(r) || unicode.Is(unicode.Zs, r) {
			b.WriteRune('-')
			continue
		}
		b.WriteRune(r)
	}
	s = b.String()

	// 3) Sustituye cualquier carácter no permitido por nada.
	s = nonAlnumDash.ReplaceAllString(s, "")

	// 4) Colapsa guiones múltiples en uno solo.
	s = multiDash.ReplaceAllString(s, "-")

	// 5) Recorta guiones al inicio y final.
	s = strings.Trim(s, "-")

	return s
}
