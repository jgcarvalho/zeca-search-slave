package db

import (
	"fmt"
	"testing"

	"github.com/jgcarvalho/zeca-search-slave/db"
)

func TestLoadProteinsFromMongo(t *testing.T) {
	// proteins := LoadProteinsFromMongo("proteindb_dev", "protein")
	// fmt.Println(len(proteins))
	proteins := db.LoadProteinsFromBoltDB("/home/jgcarvalho/sync/data/multissdb/", "chameleonic.db", "proteins")
	fmt.Println("Teste")
	fmt.Println(len(proteins))
}

func TestGetProteins(t *testing.T) {
	data := db.Config{
		Dir:    "/home/jgcarvalho/sync/data/multissdb/",
		Name:   "chameleonic.db",
		Bucket: "proteins",
		Init:   "Seq",
		Target: "All3",
	}
	start, end, e := db.GetProteins(data)
	fmt.Println(start)
	fmt.Println(end)
	fmt.Println(e)
}
