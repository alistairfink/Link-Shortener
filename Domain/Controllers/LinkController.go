package Controllers

import (
	"encoding/json"
	"github.com/alistairfink/Link-Shortener/Config"
	"github.com/alistairfink/Link-Shortener/Domain/Managers"
	"github.com/alistairfink/Link-Shortener/Domain/ViewModels"
	"github.com/go-chi/chi"
	"github.com/go-pg/pg"
	"io/ioutil"
	"log"
	"net/http"
)

type LinkController struct {
	DB          *pg.DB
	Conifg      *Config.Config
	LinkManager *Managers.LinkManager
}

func NewLinkController(db *pg.DB, config *Config.Config) *LinkController {
	controller := &LinkController{
		DB:          db,
		Conifg:      config,
		LinkManager: Managers.NewLinkManager(db, config),
	}

	return controller
}

func (this *LinkController) Routes() *chi.Mux {
	router := chi.NewRouter()
	router.Post("/", this.CreateLink)
	router.Get("/", this.GetLink)
	return router
}

func (this *LinkController) CreateLink(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		log.Print(err.Error())
		http.Error(w, "Invalid Request Object", http.StatusBadRequest)
		return
	}

	var request ViewModels.LinkRequestModel
	err = json.Unmarshal(body, &request)
	if err != nil {
		log.Print(err.Error())
		http.Error(w, "Invalid Request Object", http.StatusBadRequest)
		return
	}

	this.LinkManager.CreateLink(&request)
}

func (this *LinkController) GetLink(w http.ResponseWriter, r *http.Request) {
	println("Get")
}
