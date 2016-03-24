package ca

type Config struct {
	// 	InitStates     []string `toml:"initial-states"`
	// 	TransStates    []string `toml:"transition-states"`
	// 	Hydrophobicity string   `toml:"hydrophobicity"`
	// 	R              int      `toml:"r"`
	Steps     int `toml:"steps"`
	Consensus int `toml:"consensus"`
}
