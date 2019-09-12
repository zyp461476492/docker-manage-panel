package image

import (
	"context"
	"github.com/docker/docker/api/types"
	myClient "github.com/zyp461476492/docker-app/sdk/client"
	myType "github.com/zyp461476492/docker-app/types"
	"github.com/zyp461476492/docker-app/web/service"
	"log"
)

/**
返回 client 客户端查询的所有 image 信息
*/
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

	imageList, err := cli.ImageList(context.Background(), types.ImageListOptions{})
	if err != nil {
		log.Printf("查询失败 %s", err.Error())
		return myType.RetMsg{Res: false, Info: err.Error(), Obj: nil}
	}

	return myType.RetMsg{Res: true, Obj: imageList}
}

func Search(assetId int, term string) myType.RetMsg {
	asset, err := service.GetAsset(assetId)
	if err != nil {
		return myType.RetMsg{Res: false, Info: err.Error(), Obj: nil}
	}

	cli, err := myClient.GetClient(asset)
	if err != nil {
		log.Printf("连接失败 %s", err.Error())
		return myType.RetMsg{Res: false, Info: err.Error(), Obj: nil}
	}

	searchRes, err := cli.ImageSearch(context.Background(), term, types.ImageSearchOptions{Limit: 100})
	if err != nil {
		log.Printf("搜索失败 %s", err.Error())
		return myType.RetMsg{Res: false, Info: err.Error(), Obj: nil}
	}

	return myType.RetMsg{Res: true, Obj: searchRes}
}

/**
返回 id 对应 image 的历史信息
*/
func History(assetId int, imageId string) myType.RetMsg {
	asset, err := service.GetAsset(assetId)
	if err != nil {
		return myType.RetMsg{Res: false, Info: err.Error(), Obj: nil}
	}

	cli, err := myClient.GetClient(asset)
	if err != nil {
		log.Printf("连接失败 %s", err.Error())
		return myType.RetMsg{Res: false, Info: err.Error(), Obj: nil}
	}

	historyInfo, err := cli.ImageHistory(context.Background(), imageId)
	if err != nil {
		log.Printf("查询失败 %s", err.Error())
		return myType.RetMsg{Res: false, Info: err.Error(), Obj: nil}
	}

	return myType.RetMsg{Res: true, Obj: historyInfo}
}
