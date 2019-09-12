package image

import (
	"context"
	"github.com/docker/docker/api/types"
	"github.com/zyp461476492/docker-app/sdk/client"
	myType "github.com/zyp461476492/docker-app/types"
	"github.com/zyp461476492/docker-app/web/service"
	"io"
	"log"
)

func PullImage(assetId int, ref string) (myType.RetMsg, io.ReadCloser) {
	asset, err := service.GetAsset(assetId)
	if err != nil {
		return myType.RetMsg{Res: false, Info: err.Error(), Obj: nil}, nil
	}

	cli, err := client.GetClient(asset)
	if err != nil {
		log.Printf("连接失败 %s", err.Error())
		return myType.RetMsg{Res: false, Info: err.Error(), Obj: nil}, nil
	}

	out, err := cli.ImagePull(context.Background(), ref, types.ImagePullOptions{})
	if err != nil {
		log.Printf("拉取失败 %s", err.Error())
		return myType.RetMsg{Res: false, Info: err.Error(), Obj: nil}, nil
	}

	return myType.RetMsg{Res: true}, out
}

func DelImage(assetId int, imageId string) myType.RetMsg {
	asset, err := service.GetAsset(assetId)
	if err != nil {
		return myType.RetMsg{Res: false, Info: err.Error(), Obj: nil}
	}

	cli, err := client.GetClient(asset)
	if err != nil {
		log.Printf("连接失败 %s", err.Error())
		return myType.RetMsg{Res: false, Info: err.Error(), Obj: nil}
	}

	removeList, err := cli.ImageRemove(context.Background(), imageId, types.ImageRemoveOptions{})
	if err != nil {
		log.Printf("image delete failed %s", err.Error())
		return myType.RetMsg{Res: false, Info: err.Error(), Obj: nil}
	}

	return myType.RetMsg{Res: true, Obj: removeList}
}
