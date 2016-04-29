//+build oauth

// como páginas são desacopladas a autenticação deve ficar no angularjs
package main_test

import (
  "testing"
  "time"
  "github.com/lstern/psilibrary/server"
  "github.com/lstern/psilibrary/server/models"
  "github.com/lstern/psilibrary/server/repositories"
  "github.com/stretchr/gomniauth"  
  "github.com/stretchr/gomniauth/common"
 )

func createSession() models.Token {
	return models.Token{
		ID: "1234",
		Expiration: time.Now(),
	}
}

func Test_validate_security_key(t *testing.T){
	main.Setup();
	key := common.GetSecurityKey();
	if (key == ""){
		t.Error("Chave de segurança não definida")
	}
}

func Test_provider_config(t *testing.T) {
	main.Setup();
	l := gomniauth.SharedProviderList;
	if (l == nil) {
		t.Error("Nenhum provider configurado");
	}
}


func Test_check_callbackRoute(t *testing.T) {
	router := main.NewRouter()

	if router.Get("FacebookCallback") == nil {
		t.Error("callback de facebook não está registrado")
	}
}

func Test_create_session(t *testing.T) {
	repo := repositories.MakeAuthRepository(nil)
	email := "email@email.com"

	token := repo.CreateSession(email)

	if token.Login != email {
		t.Error("Erro ao criar token")
	}

	if token.ID == "" {
		t.Error("Session id não gerada")
	}

	if token.Expiration .IsZero() {
		t.Error("Nenhuma data de expiração definida")
	}
}

func Test_save_session(t *testing.T) {
	repo := repositories.MakeAuthRepository(nil)
	email := "email@email.com"

	token := repo.CreateSession(email)
	err := repo.SaveSession(token)

	if err != nil {
		t.Error("Ocorreu um erro ao salvar um sessão: ", err)
	}
}

func Test_get_token(t *testing.T) {
	repo := repositories.MakeAuthRepository(nil)
	email := "email@email.com"

	session := repo.CreateSession(email)
	repo.SaveSession(session)

	token, err := repo.GetToken(session.ID)

	if (err != nil || token == nil || token.ID != session.ID) {
		t.Error("Erro ao recuperar token: ", err)
	}
}

func Test_check_if_token_is_valid_valid(t *testing.T) {
	repo := repositories.MakeAuthRepository(nil)
	session := createSession()
	session.Login = "teste@teste.com"
	isValid := repo.CheckSession(session.ID, session.Login, &session)

	if !isValid {
		t.Error("Erro ao validar session")
	}
}

func Test_check_if_token_is_valid_invalid(t *testing.T) {
	repo := repositories.MakeAuthRepository(nil)
	session := createSession()
	session.Login = "teste@teste.com"
	session.Expiration = session.Expiration.AddDate(0, 0, -10)
	isValid := repo.CheckSession(session.ID, session.Login, &session)

	if isValid {
		t.Error("Erro ao validar session inválida")
	}
}

func Test_create_redirect_url(t *testing.T) {
	main.Setup();
	repo := repositories.MakeAuthRepository(nil)
	url, err := repo.GetFacebookUrl()

	if (err != nil){
		t.Error("Erro ao gerar url: ", err)
		return
	}

	if (url == ""){
		t.Error("Erro ao gerar url")
	}
}

/* 
func Test_callback_url(t *testing.T){
	main.Setup();
	repo := repositories.MakeAuthRepository(nil)
	callback := "http://psi-library.azurewebsites.net/auth/facebook/callback?code=AQDEgqkLU6x4IgUoXWoq7yK6SLpD95bBIsEVTsqlZHiZyu-HsG5fN61Kqm_nvqAuZ1yZr2q5Piy9FEU6xkt2Mv-P9O8HNgLz3_PyqNZbmi60aUjJOSViImSOK097NnIE1RMOy1czd54LQ06PR1hyAO6vJfdN9fHG7kG6-UhRpYijAuETy_tEgf-E6W_5P6y4LkOdTkoiC2HjEwamEXesy1SuQcC8jNPD2nsxSwI1fn71m-63NRTq5trI3clJrK2zVdknLDvwQIdzBWruPlF2exJ5nK724ocrnY_uZR1RSSaZqgGc8RYr-khsVTF3T_439HgQsgcd_V4jsVL0IulOEA-j&state=eyJhZnRlciI6InN1Y2Nlc3MifQ%3D%3D_963bfcb2821371a0afd3f15bf1d79085c512b7aa#_=_"
	creds, err := repo.ProcessFacebookCallback(callback)
	if (err != nil){
		t.Error("Erro ao processar callback: ", err)
		return
	}

	if (creds == nil){
		t.Error("Erro ao processar callback")
	}
}
*/