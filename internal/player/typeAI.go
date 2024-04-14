package player

import "fmt"

type iatype uint

const (
	UNKNOWNIA iatype = iota
	RANDOM
)

const (
	RandomString = "Random"
)

func (r iatype) String() string {
	switch r {
	case RANDOM:
		return RandomString
	default:
		panic("unhandled default case")
	}
}

func ParseAIType(ia string) (iatype, error) {
	switch ia {
	case RandomString:
		return RANDOM, nil
	default:
		return UNKNOWNIA, fmt.Errorf("unknown ia: %s", ia)
	}
}
