package cmd

import (
	"github.com/urfave/cli/v2"
)

// Variables declared for CLI
var (
	IpAddress string
	Ports     string
	Mode      string
	GPU         bool
	ListServers bool
)

var AppConfigFlags = []cli.Flag{
	// Deprecated to be implemented using GRPC
	&cli.StringFlag{
		Name:        "Mode",
		Value:       "client",
		Usage:       "Specifies mode of running",
		EnvVars: []string{"P2P_MODE"},
		Destination: &Mode,
	},
	&cli.BoolFlag{
		Name:        "ListServers",
		Usage:       "List servers which can render tasks",
		EnvVars: []string{"LIST_SERVERS"},
		Destination: &ListServers,
	},
	&cli.StringFlag{
		Name:        "CreateVM",
		Usage:       "Creates Docker container on the selected server",
		EnvVars: []string{"CREATE_VM"},
		Destination: &IpAddress,
	},
	&cli.StringFlag{
		Name:        "Ports",
		Usage:       "Number of ports to open for the Docker Container",
		EnvVars: []string{"NUM_PORTS"},
		Destination: &Ports,
	},
	&cli.BoolFlag{
		Name:        "GPU",
		Usage:       "Create Docker Containers to access GPU",
		EnvVars: []string{"USE_GPU"},
		Destination: &GPU,
	},
}