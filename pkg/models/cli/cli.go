package cli

import (
	"errors"
	"flag"
	"fmt"
)

const (
	ErrFServerListenHost       = "host name is not valid: '%v'"
	ErrFServerListenPort       = "port is not valid: '%v'"
	ErrFServerSharedFolderPath = "shared folder is not valid: '%v'"
	ErrNullPointer             = "null pointer"
)

type Arguments struct {
	ServerListenHost string
	ServerListenPort uint16
	SharedFolderPath string
}

func (a *Arguments) IsValid() (ok bool, err error) {
	if len(a.ServerListenHost) < 1 {
		return false, fmt.Errorf(ErrFServerListenHost, a.ServerListenHost)
	}

	if (a.ServerListenPort) < 1 {
		return false, fmt.Errorf(ErrFServerListenPort, a.ServerListenPort)
	}

	if len(a.SharedFolderPath) < 1 {
		return false, fmt.Errorf(ErrFServerSharedFolderPath, a.SharedFolderPath)
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
		return nil, errors.New(ErrNullPointer)
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
