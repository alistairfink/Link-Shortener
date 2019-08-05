package main

import (
	"fmt"
	"github.com/alistairfink/Link-Shortener/Config"
	"github.com/alistairfink/Link-Shortener/Data/DatabaseConnection"
	"github.com/alistairfink/Link-Shortener/Domain/Controllers"
	"github.com/alistairfink/Link-Shortener/Domain/Middleware"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/render"
	"github.com/go-pg/pg"
	"log"
	"net"
	"net/http"
	"os"
	"strings"
)

func main() {
	// Read in Config
	var config *Config.Config
	if _, err := os.Stat("./Config.json"); err == nil {
		config = Config.GetConfig(".", "Config")
	} else {
		config = Config.GetConfig("/go/src/github.com/alistairfink/Link-Shortener/.", "Config")
	}

	// Open DB
	db := DatabaseConnection.Connect(config)
	defer DatabaseConnection.Close(db)

	// Router
	localAddrs, _ := net.InterfaceAddrs()
	ip, _ := localAddrs[1].(*net.IPNet)
	println("=============================== Serving On ===============================")
	fmt.Printf(" %-12s%-12s\n", "Local", "localhost:"+config.Port)
	fmt.Printf(" %-12s%-12s\n", "Network", ip.IP.String()+":"+config.Port)
	println("==========================================================================\n")
	router := Routes(db, config)

	walkFunc := func(method string, route string, handler http.Handler, middlewares ...func(http.Handler) http.Handler) error {
		fmt.Printf(" %-10s%-10s\n", method, strings.Replace(route, "/*", "", -1))
		return nil
	}

	if err := chi.Walk(router, walkFunc); err != nil {
		log.Panicf("Logging err: %s\n", err.Error())
	}

	log.Fatal(http.ListenAndServe(":"+config.Port, router))
}

func Routes(db *pg.DB, config *Config.Config) *chi.Mux {
	// Init Router
	router := chi.NewRouter()
	router.Use(
		render.SetContentType(render.ContentTypeJSON),
		middleware.Logger,
		middleware.DefaultCompress,
		middleware.RedirectSlashes,
		middleware.Recoverer,
		Middleware.CorsMiddleware,
	)

	// Init Controllers
	linkController := Controllers.NewLinkController(db, config)

	// Init Paths
	router.Route("/", func(routes chi.Router) {
		routes.Mount("/link", linkController.Routes())
	})

	return router
}
