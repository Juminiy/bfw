package char

// char Stand for character

const (
	undefinedString = ""
	defaultAes      = "x"
	lambdaAes       = "λ"
)

func getSuperScript(ss string) string {
	destSS := undefinedString
	for _, ssByte := range ss {
		destSS += superScript[string(ssByte)]
	}
	return destSS
}

func getSubScript(ss string) string {
	destSS := undefinedString
	for _, ssByte := range ss {
		destSS += subScript[string(ssByte)]
	}
	return destSS
}

func GetExponent(exp string) string {
	return getSuperScript(exp)
}

func GetEquationAes(ss string, charType ...string) string {
	destAes := defaultAes
	if len(charType) > 0 {
		switch charType[0] {
		case undefinedString:
			{

			}
		case "lambda":
			{
				destAes = lambdaAes
			}
		default:
			{
				if charType[0] != undefinedString {
					destAes = charType[0]
				}
			}
		}
	}
	return destAes + getSubScript(ss)
}

func GetSubScript(ss string) string {
	return getSubScript(ss)
}

func GetUpperCaseGeekAlphabetBySpell(spell string) string {
	return upperCaseGeekAlphabet[spell]
}

func GetLowerCaseGeekAlphabetBySpell(spell string) string {
	return lowerCaseGeekAlphabet[spell]
}

// GetFraction
// example: 2,5 -> ²⁄₅
func GetFraction(numerator, denominator string) string {
	return getSuperScript(numerator) + superScript["Frac"] + getSubScript(denominator)
}
