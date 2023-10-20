package cal

const (
	undefinedString = ""
)

var (
	superScript = map[int32]string{
		'0': "⁰",
		'1': "¹",
		'2': "²",
		'3': "³",
		'4': "⁴",
		'5': "⁵",
		'6': "⁶",
		'7': "⁷",
		'8': "⁸",
		'9': "⁹",
	}
	subscript = map[int32]string{
		'0': "₀",
		'1': "₁",
		'2': "₂",
		'3': "₃",
		'4': "₄",
		'5': "₅",
		'6': "₆",
		'7': "₇",
		'8': "₈",
		'9': "₉",
	}
)

func getSuperScript(ss string) string {
	destSS := undefinedString
	for _, ssByte := range ss {
		destSS += superScript[ssByte]
	}
	return destSS
}

func getSubScript(ss string) string {
	destSS := undefinedString
	for _, ssByte := range ss {
		destSS += subscript[ssByte]
	}
	return destSS
}

func GetExponent(exp string) string {
	return getSuperScript(exp)
}
