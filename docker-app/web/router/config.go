package router

import "net/http"

func ConfigRouter() {
	http.HandleFunc("/image/list", list)
	http.HandleFunc("/image/search", search)
	http.HandleFunc("/image/history", history)
	http.HandleFunc("/image/pull", imagePull)
	http.HandleFunc("/image/del", imageDel)
	http.HandleFunc("/container/list", containerList)
	http.HandleFunc("/container/logs", containerLogs)
	http.HandleFunc("/container/create", containerCreate)
	http.HandleFunc("/container/stats", containerStats)
	http.HandleFunc("/container/start", containerStart)
	http.HandleFunc("/container/pause", containerPause)
	http.HandleFunc("/container/unpause", containerUnpause)
	http.HandleFunc("/container/stop", containerStop)
	http.HandleFunc("/asset/add", addAsset)
	http.HandleFunc("/asset/update", updateAsset)
	http.HandleFunc("/asset/delete", deleteAsset)
	http.HandleFunc("/asset/list", listAsset)
	http.HandleFunc("/asset/info", dockerInfo)
}
