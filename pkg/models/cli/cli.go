package cli

import (
	"flag"

	"github.com/vault-thirteen/auxie/errors"
)

const (
	Err_NullPointer = "null pointer"

	Errf_ServerListenHost       = "host name is not valid: '%v'"
	Errf_ServerListenPort       = "port is not valid: '%v'"
	Errf_ServerSharedFolderPath = "shared folder is not valid: '%v'"
)

type Arguments struct {
	ServerListenHost string
	ServerListenPort uint16
	SharedFolderPath string
}

func (a *Arguments) IsValid() (ok bool, err error) {
	if len(a.ServerListenHost) < 1 {
		return false, errors.ErrorFromFormattedText(Errf_ServerListenHost, a.ServerListenHost)
	}

	if (a.ServerListenPort) < 1 {
		return false, errors.ErrorFromFormattedText(Errf_ServerListenPort, a.ServerListenPort)
	}

	if len(a.SharedFolderPath) < 1 {
		return false, errors.ErrorFromFormattedText(Errf_ServerSharedFolderPath, a.SharedFolderPath)
	}

	return true, nil
}

func NewArgumentsFromOs(
	listenServerHostName string,
	listenServerPortNumber string,
	sharedFolderPath string,
) (args *Arguments, err error) {
	serverListenHostPtr := flag.String(listenServerHostName, "", "Host name of the server")
	serverListenPortPtr := flag.Uint(listenServerPortNumber, 0, "Port number of the server")
	sharedFolderPathPtr := flag.String(sharedFolderPath, "", "Path to the shared folder")

	flag.Parse()

	if (serverListenHostPtr == nil) ||
		(serverListenPortPtr == nil) ||
		(sharedFolderPathPtr == nil) {
		return nil, errors.ErrorFromText(Err_NullPointer)
	}

	args = &Arguments{
		ServerListenHost: *serverListenHostPtr,
		ServerListenPort: uint16(*serverListenPortPtr),
		SharedFolderPath: *sharedFolderPathPtr,
	}

	_, err = args.IsValid()
	if err != nil {
		return nil, err
	}

	return args, nil
}
