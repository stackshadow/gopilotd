/*
Copyright (C) 2019 by Martin Langlotz aka stackshadow

This file is part of gopilot, an rewrite of the copilot-project in go

gopilot is free software: you can redistribute it and/or modify
it under the terms of the GNU Lesser General Public License as published by
the Free Software Foundation, version 3 of this License

gopilot is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU Lesser General Public License for more details.

You should have received a copy of the GNU Lesser General Public License
along with gopilot.  If not, see <http://www.gnu.org/licenses/>.
*/

package gopilotd

import (
	"gitlab.com/gopilot/lib/gbus"
	"gitlab.com/gopilot/lib/mynodename"
)

// Init the info-modules
func Init() {

	MessageBus.Init()
	MessageBus.Run()
	Sockets = gbus.SocketNew()

	MessageBus.Subscribe("gopilotd", mynodename.NodeName, "core", onCore)
	//MessageBus.Subscribe("core/ping", mynodename.NodeName, "core", onPing)
	//MessageBus.Subscribe("core/subscriberListGet", mynodename.NodeName, "core", onSubscriberListGet)

}

// Serve will start the copilotd-socket-server
func Serve(socketBusFileName string) {
	Sockets.Serve(socketBusFileName, gbus.SocketCallbacks{
		OnHandshakeFinished: func(socket *gbus.SocketConnection) {
			MessageBus.Subscribe(
				socket.ID(),
				socket.RemoteNodeName(),
				socket.RemoteNodeGroup(),
				func(message *gbus.Msg, group, command, payload string) {
					socket.SendMessage(*message)
				},
			)
		},
		OnDisconnect: func(socket *gbus.SocketConnection) {
			MessageBus.UnSubscribeID(socket.ID())
		},
		OnMessage: func(socket *gbus.SocketConnection, message gbus.Msg) {
			MessageBus.PublishMsg(message)
		},
	})
}

func onCore(message *gbus.Msg, group, command, payload string) {
	onPing(message, group, command, payload)
	onSubscriberListGet(message, group, command, payload)
}
