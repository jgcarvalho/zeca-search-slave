package search

import "github.com/jgcarvalho/zeca-search-slave/rules"

type Tournament []Individual

type Individual struct {
	PID        uint32
	Generation int
	Rule       *rules.Rule
	Fitness    float64
	Q3         float64
	Score      float64
}
