package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	_ "net/http/pprof"
	"os"
	"runtime"

	"github.com/BurntSushi/toml"
	"github.com/jgcarvalho/zeca-search/search"
)

func run(fnconfig string) {
	var conf search.Config
	md, err := toml.DecodeFile(fnconfig, &conf)
	if err != nil {
		log.Fatal(err)
	}
	if len(md.Undecoded()) > 0 {
		fmt.Printf("Chaves desconhecidas no arquivo de configuração: %q\n", md.Undecoded())
		fmt.Printf("Chaves conhecidas: %q\n", md.Keys())
		fmt.Println("Configuration:", conf)
		return
	}
	fmt.Println("Configuration:", conf)
	search.RunClient(conf)
}

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	// flags
	fnconfig := flag.String("config", "default", "Configuration file")
	profile := flag.Bool("profile", false, "profile")
	flag.Parse()

	if *profile {
		go http.ListenAndServe(":8081", http.DefaultServeMux)
	}

	if *fnconfig == "default" {
		run(os.Getenv("GOPATH") + "/src/github.com/jgcarvalho/zeca-search-master/config.toml")
	} else {
		run(*fnconfig)
	}

}
