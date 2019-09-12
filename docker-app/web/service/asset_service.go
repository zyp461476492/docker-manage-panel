package service

import (
	"fmt"
	"github.com/asdine/storm"
	"github.com/zyp461476492/docker-app/database"
	"github.com/zyp461476492/docker-app/sdk/client"
	"github.com/zyp461476492/docker-app/types"
	"github.com/zyp461476492/docker-app/utils"
	"log"
	"time"
)

func AddAsset(asset *types.DockerAsset) types.RetMsg {
	db, err := database.GetStorm(utils.Config)
	if err != nil {
		database.CloseStorm(db)
		return types.RetMsg{Res: false, Info: types.DatabaseFail}
	}

	asset.CreateTime = time.Now().Format("2006-01-02 15:04:05")
	asset.Status = "0"
	err = db.Save(asset)

	if err != nil {
		return types.RetMsg{Res: false, Info: err.Error()}
	}

	return types.RetMsg{Res: true}
}

func UpdateAsset(asset *types.DockerAsset) types.RetMsg {
	db, err := database.GetStorm(utils.Config)
	if err != nil {
		database.CloseStorm(db)
		return types.RetMsg{Res: false, Info: types.DatabaseFail}
	}

	err = db.Update(asset)

	if err != nil {
		return types.RetMsg{Res: false, Info: err.Error()}
	}

	return types.RetMsg{Res: true}
}

func DeleteAsset(assetList []types.DockerAsset) types.RetMsg {
	db, err := database.GetStorm(utils.Config)
	if err != nil {
		database.CloseStorm(db)
		return types.RetMsg{Res: false, Info: types.DatabaseFail}
	}

	count := 0
	for _, asset := range assetList {
		err := db.DeleteStruct(&asset)
		if err != nil {
			count++
		}
	}

	info := fmt.Sprintf("成功数量 %d, 失败数量: %d", len(assetList)-count, count)

	return types.RetMsg{Res: true, Info: info}
}

func ListAsset(page, pageSize int) types.RetMsg {
	db, err := database.GetStorm(utils.Config)
	if err != nil {
		database.CloseStorm(db)
		return types.RetMsg{Res: false, Info: err.Error(), Obj: nil}
	}

	var assetList []types.DockerAsset
	skip := 0
	if (page - 1) > 0 {
		skip = (page - 1) * pageSize
	}
	err = db.All(&assetList, storm.Limit(pageSize), storm.Skip(skip))
	if err != nil {
		return types.RetMsg{Res: false, Info: err.Error(), Obj: nil}
	}

	total, err := db.Count(&types.DockerAsset{})
	if err != nil {
		database.CloseStorm(db)
		return types.RetMsg{Res: false, Info: err.Error(), Obj: nil}
	}

	obj := map[string]interface{}{}
	obj["list"] = assetList
	obj["total"] = total
	return types.RetMsg{Res: true, Info: types.SUCCESS, Obj: obj}
}

func GetAsset(id int) (types.DockerAsset, error) {
	db, err := database.GetStorm(utils.Config)
	asset := types.DockerAsset{}
	if err != nil {
		database.CloseStorm(db)
		return asset, err
	}
	err = db.One("Id", id, &asset)
	if err != nil {
		return asset, err
	}

	return asset, nil
}

func DockerInfo(id int) types.RetMsg {
	db, err := database.GetStorm(utils.Config)
	if err != nil {
		log.Print(err.Error())
		return types.RetMsg{Res: false, Info: err.Error(), Obj: nil}
	}

	asset := types.DockerAsset{}
	err = db.One("Id", id, &asset)
	if err != nil {
		return types.RetMsg{Res: false, Info: err.Error(), Obj: nil}
	}

	cli, err := client.GetClient(asset)
	if err != nil {
		log.Printf("连接失败 %s", err.Error())
		return types.RetMsg{Res: false, Info: err.Error(), Obj: nil}
	}

	info, err := client.GetClientInfo(cli)
	if err != nil {
		return types.RetMsg{Res: false, Info: err.Error(), Obj: nil}
	}
	return types.RetMsg{Res: true, Info: types.SUCCESS, Obj: info}
}
