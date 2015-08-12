package repositories

import (
	"testing"
	"log"
)

func TestEntryTypesList(t *testing.T) {
   	log.Printf("testando lista de entries")
   	types := ListEntryTypes()

	if len(types) == 0 {
		t.Error("Erro ao listar tipos de registros")
	}

    //log.Printf("%d registro(s)", len(types))
    //log.Printf("nome: %s", types[0].Name)
}
