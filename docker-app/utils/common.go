package utils

import (
	"encoding/json"
	"github.com/zyp461476492/docker-app/types"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
)

func GetConfig(filePath string) *types.Config {
	f, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
		return nil
	}

	jsonBytes, err := ioutil.ReadAll(f)
	if err != nil {
		log.Fatal(err)
		return nil
	}

	var config = types.Config{}
	err = json.Unmarshal(jsonBytes, &config)
	if err != nil {
		log.Fatal(err)
		return nil
	}
	return &config
}

func getIntValue(r *http.Request, key string) int {
	v := r.Form.Get(key)
	if v == "" {
		v = r.PostForm.Get(key)
	}

	res, err := strconv.Atoi(v)
	if err != nil {
		log.Fatalln(err)
		return -1
	}
	return res
}

func LogHttpError(err error, w *http.ResponseWriter) {
	http.Error(*w, err.Error(), http.StatusInternalServerError)
}

var Config = *GetConfig("config.json")
