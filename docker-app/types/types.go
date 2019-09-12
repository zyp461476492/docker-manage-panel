package types

import (
	"time"
)

type DockerAsset struct {
	Id         int    `storm:"id,increment" json:"id"`
	Ip         string `json:"ip"`
	AssetName  string `storm:"index" json:"assetName"`
	Port       int    `json:"port"`
	Version    string `json:"version"`
	CreateTime string `storm:"index"`
	Status     string `json:"status"`
}

type Config struct {
	FileLocation string
	Timeout      time.Duration
}

type RetMsg struct {
	Res  bool
	Info string
	Obj  interface{}
}
