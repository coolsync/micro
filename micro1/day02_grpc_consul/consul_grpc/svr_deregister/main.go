package main

import (
	"implconsul/consulexe"
	"log"
)

func main() {
	service_name := "bj38"
	err := consulexe.ConsulDeregister(service_name)
	if err != nil {
		log.Fatal(err)
	}
}
