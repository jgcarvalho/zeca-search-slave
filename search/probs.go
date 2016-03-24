package search

import "github.com/jgcarvalho/zeca-search-master/rules"

// type Pattern [3]string
type Probability map[string]float64
type ProbRule map[rules.Pattern]Probability

type Probabilities struct {
	PID        uint32
	Generation int
	Data       ProbRule
}
