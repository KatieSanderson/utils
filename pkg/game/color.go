package game

// Color : Int constant mapping
const (
	Red        = 0
	Blue       = 1
	Green      = 2
	Pink       = 3
	Orange     = 4
	Yellow     = 5
	Black      = 6
	White      = 7
	Purple     = 8
	Brown      = 9
	Cyan       = 10
	Lime       = 11
	Maroon     = 12
	Rose       = 13
	Banana     = 14
	Gray       = 15
	Tan        = 16
	Coral      = 17
	Watermelon = 18
	Chocolate  = 19
	SkyBlue    = 20
	Beige      = 21
	Magenta    = 22
	Turquoise  = 23
	Lilac      = 24
	Olive      = 25
	Azure      = 26
	Plum       = 27
	Jungle     = 28
	Mint       = 29
	Chartreuse = 30
	Macau      = 31
	Tawny      = 32
	Gold       = 33
)

// ColorStrings for lowercase, possibly for translation if needed
var ColorStrings = map[string]int{
	"red":        Red,
	"blue":       Blue,
	"green":      Green,
	"pink":       Pink,
	"orange":     Orange,
	"yellow":     Yellow,
	"black":      Black,
	"white":      White,
	"purple":     Purple,
	"brown":      Brown,
	"cyan":       Cyan,
	"lime":       Lime,
	"maroon":     Maroon,
	"rose":       Rose,
	"banana":     Banana,
	"gray":       Gray,
	"tan":        Tan,
	"coral":      Coral,
	"watermelon": Watermelon,
	"chocolate":  Chocolate,
	"skyblue":    SkyBlue,
	"beige":      Beige,
	"magenta":    Magenta,
	"turquoise":  Turquoise,
	"lilac":      Lilac,
	"olive":      Olive,
	"azure":      Azure,
	"plum":       Plum,
	"jungle":     Jungle,
	"mint":       Mint,
	"chartreuse": Chartreuse,
	"macau":      Macau,
	"tawny":      Tawny,
	"gold":       Gold,
}

// GetColorStringForInt does what it sounds like
func GetColorStringForInt(colorint int) string {
	for str, idx := range ColorStrings {
		if idx == colorint {
			return str
		}
	}
	return ""
}

// IsColorString determines if a string is actually one of our colors
func IsColorString(test string) bool {
	_, ok := ColorStrings[test]
	return ok
}
