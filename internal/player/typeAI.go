package player

import "fmt"

type iatype uint

const (
	UNKNOWNIA iatype = iota
	RANDOM
	MINMAX
	ALPHABETA
)

const (
	RandomString    = "Random"
	MINMAXString    = "MinMax"
	ALPHABETAString = "AlphaBeta"
)

func (r iatype) String() string {
	switch r {
	case RANDOM:
		return RandomString
	case MINMAX:
		return MINMAXString
	case ALPHABETA:
		return ALPHABETAString
	default:
		panic("unhandled default case")
	}
}

func ParseAIType(ia string) (iatype, error) {
	switch ia {
	case RandomString:
		return RANDOM, nil
	case MINMAXString:
		return MINMAX, nil
	case ALPHABETAString:
		return ALPHABETA, nil
	default:
		return UNKNOWNIA, fmt.Errorf("unknown ia: %s", ia)
	}
}
