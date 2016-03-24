package rules

import "github.com/jgcarvalho/zeca-search-slave/search"

type Config struct {
	Input  string `toml:"input"`
	Output string `toml:"output"`
}

type Pattern [3]string
type Rule map[Pattern]string

func GenRule(prob search.Probabilities) Rule {
	rule := make(Rule, len(prob.Data))
	for k, v := range prob.Data {
		// TODO create rule
	}
}
