package repositories

import (
  "psilibrary/server/models"
  "testing"
)

func TestEntryTypeCrud(t *testing.T) {
    if testing.Short() {
      t.Skip("skipping test in short mode.")
      return
    }

    e := new(models.EntryType)
    e.Name = "Testing Add"
    i, err := CreateEntryType(e)

    if (err != nil){
      t.Error("Erro ao inserir tipo de registro: %s", err.Error())
    }

    addedType, err := GetEntryTypeById(i)

    if (err != nil){
      t.Error("Erro ao recuperar registro inserido: %s", err.Error())
    }

    if (addedType.Name != e.Name) {
      t.Error("Dados do registro não correspondem ao registro inserido")
    }

    types, err := ListEntryTypes()
    if err != nil {
      t.Error("Erro ao listar tipos de registros: %s", err.Error())
    }

    lastType := types[len(types) - 1]

    if (lastType.ID != i) {
      t.Error("Último registro inserido não encontrado")
    }

    err = DeleteEntryType(i)

    if (err != nil) {
      t.Error("Erro ao deletar registro")
    }
}
