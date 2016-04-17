package main_test

import (
  "testing"
  "github.com/lstern/psilibrary/server"
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
