package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
	"github.com/vault-thirteen/VSFS/pkg/models/cli"
	"github.com/vault-thirteen/auxie/Versioneer"
)

// Command Line Interface Arguments.
const (
	ArgumentNameServerListenHost = "host"
	ArgumentNameServerListenPort = "port"
	ArgumentNameSharedFolderPath = "folder"
)

func main() {
	showIntro()

	cliArguments, err := cli.NewArgumentsFromOs(
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

func showIntro() {
	versioneer, err := ver.New()
	mustBeNoError(err)
	versioneer.ShowIntroText("Server")
	versioneer.ShowComponentsInfoText()
	fmt.Println()
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
