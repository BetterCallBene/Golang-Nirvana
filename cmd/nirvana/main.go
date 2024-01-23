package main

import (
	"fmt"

	"github.com/BetterCallBene/Golang-Nirvana/internal/signal"
	"github.com/pion/webrtc/v4"
)



func main() {
	fmt.Println("Start webrtc server")

	config := webrtc.Configuration{
		ICEServers: []webrtc.ICEServer{
			{
				URLs: []string{"stun:stun.stunprotocol.org:3478"},
			},
		},
	}

	peer_connection, err := webrtc.NewPeerConnection(config)
	if err != nil {
		panic(err)
	}

	defer func() {
		if err = peer_connection.Close(); err != nil {
			panic(err)
		}
	}()

	peer_connection.OnConnectionStateChange(
		func(connection_state webrtc.PeerConnectionState) {
			fmt.Println("Connection State has changed", connection_state.String())
			
			if connection_state == webrtc.PeerConnectionStateConnected {
				fmt.Println("Connected")
			} else if connection_state == webrtc.PeerConnectionStateDisconnected {
				fmt.Println("Disconnected")
			} else if connection_state == webrtc.PeerConnectionStateFailed {
				fmt.Println("Connection failed")
			} 
	})

	peer_connection.OnDataChannel(
		func(dc *webrtc.DataChannel) {
			fmt.Println("New DataChannel %s %d\n", dc.Label(), dc.ID())

			

			dc.OnMessage(
				func(msg webrtc.DataChannelMessage) {
					fmt.Printf("Message from DataChannel '%s': '%s'\n", dc.Label(), string(msg.Data))
					
				},
			)

			dc.OnOpen(
				func() {
					fmt.Printf("Data channel '%s'-'%d' open.\n", dc.Label(), dc.ID())
				})
		})

	

}
