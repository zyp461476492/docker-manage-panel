package container

import (
	"context"
	"github.com/docker/docker/api/types"
	myClient "github.com/zyp461476492/docker-app/sdk/client"
	myType "github.com/zyp461476492/docker-app/types"
	"github.com/zyp461476492/docker-app/web/service"
	"io"
	"log"
)

func List(id int) myType.RetMsg {
	asset, err := service.GetAsset(id)
	if err != nil {
		return myType.RetMsg{Res: false, Info: err.Error(), Obj: nil}
	}

	cli, err := myClient.GetClient(asset)
	if err != nil {
		log.Printf("连接失败 %s", err.Error())
		return myType.RetMsg{Res: false, Info: err.Error(), Obj: nil}
	}

	containerList, err := cli.ContainerList(context.Background(), types.ContainerListOptions{All: true})
	if err != nil {
		log.Printf("查询失败 %s", err.Error())
		return myType.RetMsg{Res: false, Info: err.Error(), Obj: nil}
	}

	return myType.RetMsg{Res: true, Obj: containerList}
}

/**
container id or name
*/
func Logs(assetId int, containerId string) io.ReadCloser {
	asset, err := service.GetAsset(assetId)
	if err != nil {
		return nil
	}

	cli, err := myClient.GetClient(asset)
	if err != nil {
		log.Printf("连接失败 %s", err.Error())
		return nil
	}

	logs, err := cli.ContainerLogs(context.Background(), containerId, types.ContainerLogsOptions{
		ShowStderr: true,
		ShowStdout: true,
		Details:    true,
		Follow:     true,
		Tail:       "50",
	})
	if err != nil {
		log.Printf("容器 LOGS 查询失败 %s", err.Error())
		return nil
	}
	return logs
}

func Stats(assetId int, containerId string) types.ContainerStats {
	asset, err := service.GetAsset(assetId)
	if err != nil {
		return types.ContainerStats{}
	}

	cli, err := myClient.GetClient(asset)
	if err != nil {
		log.Printf("连接失败 %s", err.Error())
		return types.ContainerStats{}
	}

	stats, err := cli.ContainerStats(context.Background(), containerId, true)
	if err != nil {
		log.Printf("容器 STATS 查询失败 %s", err.Error())
		return types.ContainerStats{}
	}

	return stats
}
