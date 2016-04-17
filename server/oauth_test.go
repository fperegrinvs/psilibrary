//+build integration

package main_test

import (
  "testing"
  "github.com/lstern/psilibrary/server"
  "github.com/lstern/psilibrary/server/repositories"
  "github.com/stretchr/gomniauth"  
  "github.com/stretchr/gomniauth/common"
 )

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
