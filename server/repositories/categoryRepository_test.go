package repositories

import (
  "psilibrary/server/models"
  "testing"
)

func TestCategoryCrud(t *testing.T) {
    if testing.Short() {
      t.Skip("skipping test in short mode.")
      return
    }

    e := new(models.Category)
    e.Name = "Testing Add"

    var repository CategoryRepository
    i, err := repository.CreateCategory(e, nil, repository)

    if (err != nil){
      t.Error("Erro ao inserir categoria: %s", err.Error())
    }

    addedType, err := repository.GetCategoryById(i, nil)

    if (err != nil){
      t.Error("Erro ao recuperar categoria inserida: %s", err.Error())
    }

    if (addedType.Name != e.Name) {
      t.Error("Dados da categoria não correspondem à categoria inserida")
    } 

    types, err := repository.ListCategories(nil)
    if err != nil {
      t.Error("Erro ao listar categorias: %s", err.Error())
    }

    lastType := types[len(types) - 1]

    if (lastType.ID != i) {
      t.Error("Último categoria inserida não encontrada")
    }

    _,err = repository.DeleteCategory(i, nil, repository)

    if (err != nil) {
      t.Error("Erro ao deletar categoria")
    }
}
