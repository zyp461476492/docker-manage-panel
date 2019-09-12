package router

import (
	"bufio"
	"encoding/json"
	"github.com/gorilla/websocket"
	"github.com/zyp461476492/docker-app/sdk/container"
	"log"
	"net/http"
	"strconv"
	"time"
)

func containerList(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	id, err := strconv.Atoi(r.Form.Get("index"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	msg := container.List(id)
	jsonByte, err := json.Marshal(msg)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	value, err := w.Write(jsonByte)

	if err != nil {
		log.Fatalf("return value %v, err %v", value, err)
	}
}

func containerStart(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	id, err := strconv.Atoi(r.Form.Get("assetId"))
	containerId := r.Form.Get("containerId")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	msg := container.Start(id, containerId)
	jsonByte, err := json.Marshal(msg)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	value, err := w.Write(jsonByte)

	if err != nil {
		log.Fatalf("return value %v, err %v", value, err)
	}
}

func containerPause(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	id, err := strconv.Atoi(r.Form.Get("assetId"))
	containerId := r.Form.Get("containerId")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	msg := container.Pause(id, containerId)
	jsonByte, err := json.Marshal(msg)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	value, err := w.Write(jsonByte)

	if err != nil {
		log.Fatalf("return value %v, err %v", value, err)
	}
}

func containerUnpause(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	id, err := strconv.Atoi(r.Form.Get("assetId"))
	containerId := r.Form.Get("containerId")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	msg := container.Unpause(id, containerId)
	jsonByte, err := json.Marshal(msg)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	value, err := w.Write(jsonByte)

	if err != nil {
		log.Fatalf("return value %v, err %v", value, err)
	}
}

func containerStop(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	id, err := strconv.Atoi(r.Form.Get("assetId"))
	containerId := r.Form.Get("containerId")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	msg := container.Stop(id, containerId)
	jsonByte, err := json.Marshal(msg)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	value, err := w.Write(jsonByte)

	if err != nil {
		log.Fatalf("return value %v, err %v", value, err)
	}
}

func containerStats(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	c, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Printf("err %s", err.Error())
	}
	assetId, err := strconv.Atoi(r.Form.Get("assetId"))

	if err != nil {
		log.Fatalln(err)
	}
	containerId := r.Form.Get("containerId")

	stats := container.Stats(assetId, containerId)

	if stats.Body == nil {
		c.Close()
	}

	reader := bufio.NewReader(stats.Body)

	for {
		str, err := reader.ReadString('\n')
		if err != nil {
			log.Printf("Read Error: %s", err)
			err = stats.Body.Close()
			if err != nil {
				log.Printf("stream close error: %s", err)
			}
			c.Close()
			return
		}
		// 去掉八个字节的头部信息
		err = c.WriteMessage(websocket.TextMessage, []byte(str))
		if err != nil {
			log.Printf("websocket WriteMessage error : %s", err.Error())
			err = stats.Body.Close()
			if err != nil {
				log.Printf("stream close error: %s", err)
			}
			c.Close()
			return
		}

		time.Sleep(time.Duration(2) * time.Second)
	}
}

func containerLogs(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	c, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Printf("err %s", err.Error())
	}
	assetId, err := strconv.Atoi(r.Form.Get("assetId"))

	if err != nil {
		log.Fatalln(err)
	}
	containerId := r.Form.Get("containerId")

	stream := container.Logs(assetId, containerId)

	if stream == nil {
		err = c.WriteMessage(1, []byte("{\"status\":"+"空"+"}"))
		if err != nil {
			log.Printf("websocket WriteMessage error : %s", err.Error())
		}

		c.Close()
	}

	reader := bufio.NewReader(stream)

	for {
		str, err := reader.ReadString('\n')
		if err != nil {
			log.Printf("Read Error: %s", err)
			err = stream.Close()
			if err != nil {
				log.Printf("stream close error: %s", err)
			}
			c.Close()
			return
		}
		// 去掉八个字节的头部信息
		err = c.WriteMessage(websocket.TextMessage, []byte(str)[8:])
		if err != nil {
			log.Printf("websocket WriteMessage error : %s", err.Error())
			err = stream.Close()
			if err != nil {
				log.Printf("stream close error: %s", err)
			}
			c.Close()
			return
		}
	}
}
