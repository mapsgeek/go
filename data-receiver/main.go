package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
	"github.com/mapsgeek/tolling/types"
	"github.com/segmentio/kafka-go"
)

const wsPort = 30000

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1028,
	WriteBufferSize: 1028,
}

type DataReceiver struct {
	msgch chan types.OBUData
	conn  *websocket.Conn
	prod  *kafka.Conn
}

func NewDataReceiver() *DataReceiver {
	prod, err := kafka.DialLeader(context.Background(), "tcp", "localhost:29092", "obu-data", 0)
	if err != nil {
		log.Fatal("failed to dial leader:", err)
	}
	return &DataReceiver{
		msgch: make(chan types.OBUData, 128),
		prod:  prod,
	}
}

func (dr *DataReceiver) produceData(data types.OBUData) error {
	b, err := json.Marshal(data)
	if err != nil {
		return err
	}
	_, err = dr.prod.WriteMessages(
		kafka.Message{Value: b},
	)
	return err
}

func (dr *DataReceiver) wsHandler(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Fatal("can't recive data from clients ", err)
	}
	dr.conn = conn
	go dr.wsReceiveLoop()
}

func (dr *DataReceiver) wsReceiveLoop() {
	fmt.Println("New OBU connected")
	for {
		var data types.OBUData
		err := dr.conn.ReadJSON(&data)
		if err != nil {
			log.Println("read error", err)
			continue
		}
		fmt.Printf("recevied data [%d] :: <lat %.2f, lng %.2f> \n", data.OBUID, data.Lat, data.Lng)
		dr.produceData(data)
		// dr.msgch <- data
	}
}

func main() {
	reciver := NewDataReceiver()
	http.HandleFunc("/ws", reciver.wsHandler)
	http.ListenAndServe(":30000", nil)
	fmt.Println("Connected to port 30000")
}
