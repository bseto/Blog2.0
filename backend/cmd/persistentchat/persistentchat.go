// main
// notes:
// on the frontend, send this, and it'll create a client
// {
//   "ContainsToken": false,
//   "Token": ""
// }
package main

import (
	"flag"
	"fmt"
	"net/http"

	"github.com/bseto/arcade/backend/game"
	"github.com/bseto/arcade/backend/hub/hubmanager"
	"github.com/bseto/arcade/backend/log"
	"github.com/bseto/arcade/backend/websocket"
	"github.com/bseto/arcade/backend/websocket/identifier"
	"github.com/bseto/arcade/backend/websocket/registry"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

var port *int = flag.Int("port", 8081, "defines the port to listen and serve on")

// PersistentChat implements both GameFactory interface and GameRouter interface
type PersistentChat struct{}

func (p *PersistentChat) RouteMessage(
	messageType int,
	message []byte,
	clientID identifier.Client,
	messageErr error,
	reg registry.Registry,
) {
	log.Infof("got some message: %v", message)
}
func (p *PersistentChat) NewClient(
	clientID identifier.Client,
	reg registry.Registry,
) {
	log.Infof("new client connected: %v", clientID)
}

func (p *PersistentChat) ClientQuit(
	clientID identifier.Client,
	reg registry.Registry,
) {
	log.Infof("client quit: %v", clientID)
}

func (p *PersistentChat) RouterName() string {
	return "Chat"
}

func (p *PersistentChat) GetAvailableGames() []string {
	return []string{"Chat"}
}

func (p *PersistentChat) GetGame(string, registry.Registry) game.GameRouter {
	return p
}

func main() {
	r := mux.NewRouter()
	hubManager := hubmanager.GetHubManager()
	hubManager.SetupRoutes(r)
	var persistentChat PersistentChat
	r.PathPrefix("/ws/{hubID}").HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		hubInstance, err := hubManager.GetHub(req, &persistentChat)
		if err != nil {
			log.Errorf("unable to create hubInstance: %v", err)
			return
		}

		wsClient := websocket.GetClient(hubInstance)
		err = wsClient.Upgrade(w, req)
		if err != nil {
			log.Errorf("unable to upgrade websocket: %v", err)
			return
		}
		wsClient.RegisterCloseListener(hubManager)
	})
	address := fmt.Sprintf(":%v", *port)
	log.Infof("starting server on: %v", address)

	headersOk := handlers.AllowedHeaders([]string{"X-Requested-With"})
	originsOk := handlers.AllowedOrigins([]string{"*"}) // we need to add our domain name here one day
	methodsOk := handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "OPTIONS"})

	err := http.ListenAndServe(address, handlers.CORS(originsOk, headersOk, methodsOk)(r))
	log.Fatalf("unable to listen and serve: %v", err)
}
