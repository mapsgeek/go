package main

import (
	"log"
	"time"

	"github.com/gorilla/websocket"
	"github.com/mapsgeek/tolling/types"
)

var sendInterval = time.Second

const wsEndpoint = "ws://localhost:30000/ws"

func main() {
	bbox := &types.BoundingBox{
		MinLat: 30.0,
		MaxLat: 40.0,
		MinLng: -120,
		MaxLng: -110,
	}

	obuIDS := GenerateOBUIDS(20)

	conn, _, err := websocket.DefaultDialer.Dial(wsEndpoint, nil)
	if err != nil {
		log.Fatal("error can't connect ", err)
	}
	defer conn.Close()

	for {
		for i := 0; i < len(obuIDS); i++ {
			lat, lng := GenLocation(bbox)
			data := &types.OBUData{
				OBUID: obuIDS[i],
				Lat:   lat,
				Lng:   lng,
			}
			err := conn.WriteJSON(&data)
			if err != nil {
				log.Fatal("can't write data to server ", err)
			}
		}
		time.Sleep(sendInterval)
	}
}
