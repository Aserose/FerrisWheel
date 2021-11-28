package server

import (
	"encoding/json"
	"fmt"
	"github.com/Aserose/ferrisWheel/internal/config"
	"github.com/Aserose/ferrisWheel/internal/logger"
	"github.com/Aserose/ferrisWheel/internal/web/tg"
	"io/ioutil"
	"net/http"
)

type Server interface {
	Start() error
}

type (
	authServer struct {
		Server  *http.Server
		Log     logger.Logger
		tgClient tg.TgApi
		Cfg     *config.ServerConfig
	}

	token struct {
		AccessToken string `json:"access_token"`
	}
)

func NewAuthServer(log logger.Logger, cfg *config.ServerConfig, tgClient tg.TgApi) Server {
	return &authServer{
		Log:     log,
		Cfg:     cfg,
		tgClient: tgClient,
	}
}

func (s *authServer) Start() error {

	s.Server = &http.Server{
		Addr:    ":"+s.Cfg.Port,
		Handler: s,
	}

	return s.Server.ListenAndServe()
}

func (s *authServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.createAccessTokenVK(w, r)
}

func (s *authServer) createAccessTokenVK(w http.ResponseWriter, r *http.Request) {
	keys := r.URL.Query()

	resp, err := http.Get(fmt.Sprintf("https://oauth.vk.com/access_token?client_id=%s&client_secret=%s&redirect_uri=%s&code=%s", s.Cfg.ClientID, s.Cfg.ClientSecret, s.Cfg.AuthServerURL, keys.Get("code")))
	if err != nil {
		s.Log.Errorf("server: failed with method get: %s", err.Error())
	}

	var t token

	if err := json.Unmarshal(func() []byte { a, _ := ioutil.ReadAll(resp.Body); return a }(), &t); err != nil {
		s.Log.Errorf("server: failed with unmarshal: %s", err.Error())
	}

	s.tgClient.CreateAccessKey(t.AccessToken)

	http.Redirect(w, r, s.Cfg.Redirect, 301)
}
