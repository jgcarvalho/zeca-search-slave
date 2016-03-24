package search

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"time"

	"github.com/jgcarvalho/zeca-search-slave/ca"
	"github.com/jgcarvalho/zeca-search-slave/db"
	"github.com/jgcarvalho/zeca-search-slave/rules"
	zmq "github.com/pebbe/zmq4"
)

type Config struct {
	Title string
	EDA   edaConfig
	Rules rules.Config
	DB    db.Config
	CA    ca.Config
	Dist  distConfig
}

type edaConfig struct {
	Generations int
	Population  int
	Tournament  int
	OutputProbs string `toml:"output-probabilities"`
	SaveSteps   int    `toml:"save-steps"`
}

type distConfig struct {
	MasterURL string `toml:"master-url"`
	PortA     string `toml:"port-a"`
	PortB     string `toml:"port-b"`
}

func Run(conf Config) {

	// Cria o receptor que recebe a probabilidade emitida pelo master na porta A
	receiver, _ := zmq.NewSocket(zmq.PULL)
	defer receiver.Close()
	receiver.Connect("tcp://" + conf.Dist.MasterURL + ":" + conf.Dist.PortA)

	// Cria o emissor que envia o individuo vencedor do torneio na rede pela
	// porta B
	sender, _ := zmq.NewSocket(zmq.PUSH)
	defer sender.Close()
	sender.Connect("tcp://" + conf.Dist.MasterURL + ":" + conf.Dist.PortB)

	// semente randomica
	rand.Seed(time.Now().UTC().UnixNano())

	// Le os dados das proteinas no DB
	fmt.Println("Loading proteins...")
	start, end, err := db.GetProteins(conf.DB)
	if err != nil {
		fmt.Println("Erro no banco de DADOS")
		panic(err)
	}
	fmt.Println("Done")

	var prob Probabilities

	var tourn Tournament
	tourn = make([]Individual, conf.EDA.Tournament)

	var (
		ind    Individual
		b      []byte
		m      string
		conerr error
	)

	for {
		// m é a mensagem com as probabilidades
		m, conerr = receiver.Recv(0)
		if conerr == nil {
			json.Unmarshal([]byte(m), &prob)
			fmt.Printf("PID: %d, Geracacao: %d\n", prob.PID, prob.Generation)

			for i := 0; i < conf.EDA.Tournament; i++ {

			}

		}
	}

	// para cada probabilidade recebida
	//criar t individuos do torneio
	// contruir e rodar automato celular
	// computar o Score

}
