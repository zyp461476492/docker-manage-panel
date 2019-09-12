package client

import (
	"context"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
	localType "github.com/zyp461476492/docker-app/types"
	"strconv"
)

func GetClient(asset localType.DockerAsset) (*client.Client, error) {
	// tcp://192.168.184.123:2376
	host := "tcp://" + asset.Ip + ":" + strconv.Itoa(asset.Port)
	cli, err := client.NewClient(host, asset.Version, nil, nil)
	return cli, err
}

func GetClientInfo(cli *client.Client) (types.Info, error) {
	ctx := context.Background()
	return cli.Info(ctx)
}
