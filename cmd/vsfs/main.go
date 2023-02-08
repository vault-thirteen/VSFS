package main

import (
	"log"
	"net"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
	"github.com/vault-thirteen/vsfs/pkg/models/cli"
)

// Command Line Interface Arguments.
const (
	ArgumentNameServerListenHost = "host"
	ArgumentNameServerListenPort = "port"
	ArgumentNameSharedFolderPath = "folder"
)

func main() {
	var err error
	var cliArguments *cli.Arguments
	cliArguments, err = cli.NewArgumentsFromOs(
		ArgumentNameServerListenHost,
		ArgumentNameServerListenPort,
		ArgumentNameSharedFolderPath,
	)
	mustBeNoError(err)

	log.Printf("Settings: %+v.\r\n", cliArguments)

	err = listen(cliArguments)
	mustBeNoError(err)
}

func mustBeNoError(err error) {
	if err != nil {
		panic(err)
	}
}

func listen(cliArguments *cli.Arguments) (err error) {
	router := httprouter.New()

	router.ServeFiles("/*filepath", http.Dir(cliArguments.SharedFolderPath))

	httpServer := http.Server{
		Addr: net.JoinHostPort(
			cliArguments.ServerListenHost,
			strconv.FormatUint(uint64(cliArguments.ServerListenPort), 10),
		),
		Handler: router,
	}

	err = httpServer.ListenAndServe()
	if err != nil {
		log.Println(err)
	}

	return nil
}
