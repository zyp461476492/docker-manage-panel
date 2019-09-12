package container

import (
	"context"
	"github.com/docker/docker/api/types"
	myClient "github.com/zyp461476492/docker-app/sdk/client"
	myType "github.com/zyp461476492/docker-app/types"
	"github.com/zyp461476492/docker-app/web/service"
	"log"
	"time"
)

func Start(assetId int, containerId string) myType.RetMsg {
	asset, err := service.GetAsset(assetId)
	if err != nil {
		return myType.RetMsg{Res: false, Info: err.Error(), Obj: nil}
	}

	cli, err := myClient.GetClient(asset)
	if err != nil {
		log.Printf("连接失败 %s", err.Error())
		return myType.RetMsg{Res: false, Info: err.Error(), Obj: nil}
	}

	err = cli.ContainerStart(context.Background(), containerId, types.ContainerStartOptions{})
	if err != nil {
		log.Printf("启动失败 %s", err.Error())
		return myType.RetMsg{Res: false, Info: err.Error(), Obj: nil}
	}

	return myType.RetMsg{Res: true, Info: "启动成功"}
}

func Pause(assetId int, containerId string) myType.RetMsg {
	asset, err := service.GetAsset(assetId)
	if err != nil {
		return myType.RetMsg{Res: false, Info: err.Error(), Obj: nil}
	}

	cli, err := myClient.GetClient(asset)
	if err != nil {
		log.Printf("连接失败 %s", err.Error())
		return myType.RetMsg{Res: false, Info: err.Error(), Obj: nil}
	}

	err = cli.ContainerPause(context.Background(), containerId)
	if err != nil {
		log.Printf("暂停失败 %s", err.Error())
		return myType.RetMsg{Res: false, Info: err.Error(), Obj: nil}
	}

	return myType.RetMsg{Res: true}
}

func Unpause(assetId int, containerId string) myType.RetMsg {
	asset, err := service.GetAsset(assetId)
	if err != nil {
		return myType.RetMsg{Res: false, Info: err.Error(), Obj: nil}
	}

	cli, err := myClient.GetClient(asset)
	if err != nil {
		log.Printf("连接失败 %s", err.Error())
		return myType.RetMsg{Res: false, Info: err.Error(), Obj: nil}
	}

	err = cli.ContainerUnpause(context.Background(), containerId)
	if err != nil {
		log.Printf("恢复失败 %s", err.Error())
		return myType.RetMsg{Res: false, Info: err.Error(), Obj: nil}
	}

	return myType.RetMsg{Res: true}
}

func Stop(assetId int, containerId string) myType.RetMsg {
	asset, err := service.GetAsset(assetId)
	if err != nil {
		return myType.RetMsg{Res: false, Info: err.Error(), Obj: nil}
	}

	cli, err := myClient.GetClient(asset)
	if err != nil {
		log.Printf("连接失败 %s", err.Error())
		return myType.RetMsg{Res: false, Info: err.Error(), Obj: nil}
	}

	timeout := 2 * time.Second
	err = cli.ContainerStop(context.Background(), containerId, &timeout)
	if err != nil {
		log.Printf("停止失败 %s", err.Error())
		return myType.RetMsg{Res: false, Info: err.Error(), Obj: nil}
	}

	return myType.RetMsg{Res: true}
}
