package cal

import "bfw/wheel/adt"

type Group struct {
	opt rune
	set map[rune]Elem
}

type Elem adt.Pair[rune, int]
