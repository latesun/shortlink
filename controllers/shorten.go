package controllers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/justinas/alice"
	"github.com/latesun/shortlink/db"
	"github.com/latesun/shortlink/models"
)

type shortenReq struct {
	LongURL string `json:"url" valid:"nonzero"`
	Secret  string `json:"secret" valid:"regexp=[bindolabs|sVdbhgpC]"`
	Length  int    `json:"length" valid:"min=6,max=32"`
	Expires int    `json:"expires" valid:"nonzero"`
}

type shortLinkResp struct {
	LongURL    string `json:"url"`
	ShortURL   string `json:"short_url"`
	ExpireDate string `json:"expire_date"`
	Length     string `json:"length"`
	Expires    string `json:"expires"`
}
type App struct {
	Router     *mux.Router
	Middleware *Middleware
}

func (a *App) Initialize() {
	// set log formatter
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	a.Router = mux.NewRouter()
	a.Middleware = &Middleware{}
	a.initializeRoutes()
}

func (a *App) initializeRoutes() {
	m := alice.New(a.Middleware.LoggingHandler, a.Middleware.RecoverHandler)
	a.Router.Handle("/api/shorten", m.ThenFunc(a.createShortlink)).Methods("POST")
	a.Router.Handle("/api/info", m.ThenFunc(a.getShortLinkInfo)).Methods("GET")
	a.Router.Handle("/{shortlink:[a-zA-Z0-9]{6,32}}", m.ThenFunc(a.redirect)).Methods("GET")
}

func (a *App) createShortlink(w http.ResponseWriter, r *http.Request) {
	var req shortenReq
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		responseWithError(w, StatusError{http.StatusBadRequest, fmt.Errorf("parse parameters failed %v", r.Body)})
		return
	}

	if err := validator.Validate(req); err != nil {
		responseWithError(w, StatusError{http.StatusBadRequest, fmt.Errorf("validate parameters failed %v", req)})
		return
	}

	defer r.Body.Close()

	log.Printf("%#v\n", req)
}

func (a *App) getShortLinkInfo(w http.ResponseWriter, r *http.Request) {
	var (
		shortener models.Shortener
		resp      shortLinkResp
	)

	vals := r.URL.Query()
	s := vals.Get("shortlink")

	if err := db.DB.Model(shortener).Where("shorten_key = ?", s).Find(&shortener).Error; err != nil {
		log.Printf("Get shortener failed, %s\n", err.Error())
		responseWithError(w, StatusError{http.StatusNotFound, fmt.Errorf("query shortlink failed, %v", err.Error())})
	}

	resp.LongURL = shortener.LongURL
	resp.ShortURL = shortener.ShortenKey

	responseWithJSON(w, http.StatusOK, resp)

	log.Printf("%s\n", s)
}

func (a *App) redirect(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	log.Printf("%s\n", vars["shortlink"])
}

func (a *App) Run(addr string) {
	log.Fatal(http.ListenAndServe(addr, a.Router))
}
