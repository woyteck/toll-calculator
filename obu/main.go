package main

import (
	"log"
	"math/rand"
	"time"

	"github.com/gorilla/websocket"
	"github.com/woyteck/toll-calculator/types"
)

const wsEndpoint = "ws://127.0.0.1:30000/ws"

var sendInterval = time.Second * 5

func main() {
	obuIDS := generateOBUIDS(20)
	conn, _, err := websocket.DefaultDialer.Dial(wsEndpoint, nil)
	if err != nil {
		log.Fatal(err)
	}

	for {
		for i := 0; i < len(obuIDS); i++ {
			data := types.OBUData{
				OBUID: obuIDS[i],
				Lat:   genCoord(),
				Long:  genCoord(),
			}
			// fmt.Printf("%+v\n", data)
			if err := conn.WriteJSON(data); err != nil {
				log.Fatal(err)
			}
		}

		time.Sleep(sendInterval)
	}
}

func generateOBUIDS(n int) []int {
	ids := make([]int, n)
	for i := 0; i < n; i++ {
		ids[i] = rand.Intn(999999)
	}

	return ids
}

func genCoord() float64 {
	n := float64(rand.Intn(100) + 1)
	f := rand.Float64()

	return n + f
}
