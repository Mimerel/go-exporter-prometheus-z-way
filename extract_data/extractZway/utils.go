package extractZway

import (
	"golang.org/x/text/transform"
	"golang.org/x/text/unicode/norm"
	"strings"
	"unicode"
)

func isMn(r rune) bool {
	return unicode.Is(unicode.Mn, r) // Mn: nonspacing marks
}

func RemoveAccents(value string) string{
	b := make([]byte, len(value))

	t := transform.Chain(norm.NFD, transform.RemoveFunc(isMn), norm.NFC)
	nDst, _, e := t.Transform(b, []byte(value), true)
	if e != nil {
		panic(e)
	}
	return string(b[:nDst])
}

func Trim(value string) (string) {
	return strings.TrimSpace(strings.ToLower(RemoveAccents(value)))
}

func (data *Data) validTypes(value string) (bool) {
	for _, x := range data.Conf.DeviceTypes {
		if strings.ToLower(x) == Trim(value) {
			return true
		}
	}
	return false
}

func BoolToIntensity(value bool) (float64) {
	if value == true {
		return 255.0
	} else {
		return 0.0
	}
}