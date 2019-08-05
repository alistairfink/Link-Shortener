package Controllers

import (
	"github.com/alistairfink/Link-Shortener/Domain/Managers"
	"github.com/go-chi/chi"
	"github.com/go-pg/pg"
	"net/http"
)

type LinkController struct {
	DB          *pg.DB
	LinkManager *Managers.LinkManager
}

func NewLinkController(db *pg.DB) *LinkController {
	controller := &LinkController{
		DB:          db,
		LinkManager: Managers.NewLinkManager(db),
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
	println("Create")
}

func (this *LinkController) GetLink(w http.ResponseWriter, r *http.Request) {
	println("Get")
}
