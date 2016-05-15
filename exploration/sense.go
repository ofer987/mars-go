package exploration

import ()

type Sense int

func NewSense(symbol rune) *Sense {
	var sense Sense

	switch symbol {
	case 'N':
		sense = 0
	case 'E':
		sense = 1
	case 'S':
		sense = 2
	case 'W':
		sense = 3
	default:
		return nil
	}

	return &sense
}

func (se *Sense) Symbol() rune {
	var symbol rune
	numericValue := int(*se) % 4

	switch numericValue {
	case 0:
		symbol = 'N'
	case 1:
		symbol = 'E'
	case 2:
		symbol = 'S'
	case 3:
		symbol = 'W'
	}

	return symbol
}

func (se *Sense) String() string {
	return string(se.Symbol())
}
