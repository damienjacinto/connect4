package player

import "fmt"

type iatype uint

const (
	UNKNOWNIA iatype = iota
	RANDOM
	MINMAX
)

const (
	RandomString = "Random"
	MINMAXString = "MinMax"
)

func (r iatype) String() string {
	switch r {
	case RANDOM:
		return RandomString
	case MINMAX:
		return MINMAXString
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
	default:
		return UNKNOWNIA, fmt.Errorf("unknown ia: %s", ia)
	}
}
