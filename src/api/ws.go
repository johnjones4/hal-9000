package main

import (
	"crypto/sha1"
	"encoding/json"
	"errors"
	"fmt"
	"hal9000"
	"hal9000/types"
	"net/http"
	"strconv"

	"github.com/gorilla/websocket"
)

type InterfaceTypeWebsocket struct {
	Connection       *websocket.Conn
	Open             bool
	VisualsSupported bool
}

func (i *InterfaceTypeWebsocket) Type() string {
	return "websocket"
}

func (i *InterfaceTypeWebsocket) ID() string {
	h := sha1.New()
	h.Write([]byte(i.Connection.RemoteAddr().String()))
	bs := h.Sum(nil)
	return fmt.Sprintf("ws-%x", bs)
}

func (i *InterfaceTypeWebsocket) IsStillValid() bool {
	return i.Open
}

func (i *InterfaceTypeWebsocket) SupportsVisuals() bool {
	return i.VisualsSupported
}

func (i *InterfaceTypeWebsocket) SendMessage(m types.ResponseMessage) error {
	responseBytes, err := json.Marshal(m)
	if err != nil {
		i.Open = false
		return err
	}
	err = i.Connection.WriteMessage(websocket.TextMessage, responseBytes)
	if err != nil {
		i.Open = false
		return err
	}
	return nil
}

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func wsHandler(runtime *types.Runtime) func(w http.ResponseWriter, req *http.Request) {
	return func(w http.ResponseWriter, req *http.Request) {
		userId := req.URL.Query().Get("user")
		if userId == "" {
			errorResponse(w, errors.New("no user id provided"))
			return
		}

		visualsStr := req.URL.Query().Get("visuals")
		visuals := false
		if visualsStr != "" {
			visuals, _ = strconv.ParseBool(visualsStr)
		}

		person, err := (*(*runtime).People()).GetPersonByID(userId)
		if err != nil {
			errorResponse(w, err)
			return
		}

		c, err := upgrader.Upgrade(w, req, nil)
		if err != nil {
			errorResponse(w, err)
			return
		}

		defer c.Close()

		iface := InterfaceTypeWebsocket{c, true, visuals}
		var iface1 types.Interface = &iface
		(*(*runtime).InterfaceStore()).Register(person, &iface1)
		ses := hal9000.NewSession(runtime, person, &iface1)

		for {
			_, request, err := c.ReadMessage()
			if err != nil {
				(*(*runtime).Logger()).LogError(err)
				iface.Open = false
				return
			}

			var halReq types.RequestMessage
			err = json.Unmarshal(request, &halReq)
			if err != nil {
				(*(*runtime).Logger()).LogError(err)
				iface.Open = false
				return
			}

			response, err := hal9000.ProcessIncomingMessage(runtime, &ses, halReq)
			if err != nil {
				(*(*runtime).Logger()).LogError(err)
				iface.Open = false
				return
			}

			err = (*ses.Interface).SendMessage(response)
			if err != nil {
				(*(*runtime).Logger()).LogError(err)
				iface.Open = false
				return
			}

			(*(*runtime).SessionStore()).SaveSession(&ses)
			if err != nil {
				(*(*runtime).Logger()).LogError(err)
				iface.Open = false
				return
			}
		}
	}
}
