package main

import (
	"flag"
	"fmt"
	"net/http"
	_ "net/http/pprof"
	"runtime"

	"github.com/jgcarvalho/zeca-search/search"
)

// func run(fnconfig string) {
// 	var conf search.Config
// 	md, err := toml.DecodeFile(fnconfig, &conf)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	if len(md.Undecoded()) > 0 {
// 		fmt.Printf("Chaves desconhecidas no arquivo de configuração: %q\n", md.Undecoded())
// 		fmt.Printf("Chaves conhecidas: %q\n", md.Keys())
// 		fmt.Println("Configuration:", conf)
// 		return
// 	}
// 	fmt.Println("Configuration:", conf)
// 	search.RunClient(conf)
// }

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	// flags
	profile := flag.Bool("profile", true, "profile")
	flag.Parse()

	// Profile
	if *profile {
		go http.ListenAndServe(":8081", http.DefaultServeMux)
	}

	serverIP := flag.Arg(0)
	if serverIP == "" {
		fmt.Println("Please, tell me the zeca-server IP.")
		return
	}
	search.RunClient(serverIP)
}
