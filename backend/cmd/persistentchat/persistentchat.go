// main
// notes:
// on the frontend, send this, and it'll create a client
// {
//   "ContainsToken": false,
//   "Token": ""
// }
package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/bseto/arcade/backend/game"
	"github.com/bseto/arcade/backend/hub/hubmanager"
	"github.com/bseto/arcade/backend/log"
	"github.com/bseto/arcade/backend/websocket"
	"github.com/bseto/arcade/backend/websocket/identifier"
	"github.com/bseto/arcade/backend/websocket/registry"
	"github.com/bseto/blog2/backend/pkg/data"
	"github.com/bseto/blog2/backend/pkg/database"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

// PersistentChat implements both GameFactory interface and GameRouter interface
type PersistentChat struct{}

func (p *PersistentChat) RouteMessage(
	messageType int,
	message []byte,
	clientID identifier.Client,
	messageErr error,
	reg registry.Registry,
) {
	log.Infof("echoing the message: %v", string(message))
	reg.SendToSameHub(message) // just repeat the message to all in the same hub
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

type Presale struct {
	gorm.Model
	WalletPublicKey  data.NullString `gorm:"column:wallet_public_key"` // The public key
	DiscordID        data.NullString `gorm:"column:discord_id"`
	NumAllowedToMint data.NullInt64  `gorm:"column:num_allowed_to_mint"`
}

func main() {
	rootPassword := os.Getenv("MYSQL_ROOT_PASSWORD")
	password := os.Getenv("MYSQL_PASSWORD")
	port := os.Getenv("MYSQL_PORT")
	databaseName := os.Getenv("MYSQL_DATABASE")
	user := os.Getenv("MYSQL_USER")
	chatPort := os.Getenv("PERSISTENTCHAT_PORT")

	if port == "" {
		log.Fatalf("requires MYSQL_PORT defined")
	}
	if rootPassword == "" {
		log.Fatalf("requires MYSQL_ROOT_PASSWORD defined")
	}
	if password == "" {
		log.Fatalf("requires MYSQL_PASSWORD defined")
	}
	if databaseName == "" {
		log.Fatalf("requires MYSQL_DATABASE defined")
	}
	if user == "" {
		log.Fatalf("requires MYSQL_USER defined")
	}

	db, _, err := database.SetupDB(
		"root",
		rootPassword,
		user,
		password,
		"db", // hostname of the 'db' service on the dockerfile
		port,
		databaseName,
	)
	if err != nil {
		log.Fatalf("unable to setup db: %v", err)
	}
	db.AutoMigrate(&Presale{})

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
	address := fmt.Sprintf(":%v", chatPort)
	log.Infof("starting server on: %v", address)

	headersOk := handlers.AllowedHeaders([]string{"X-Requested-With"})
	originsOk := handlers.AllowedOrigins([]string{"*"}) // we need to add our domain name here one day
	methodsOk := handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "OPTIONS"})

	err = http.ListenAndServe(address, handlers.CORS(originsOk, headersOk, methodsOk)(r))
	log.Fatalf("unable to listen and serve: %v", err)
}
