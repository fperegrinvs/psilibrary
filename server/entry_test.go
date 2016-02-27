//+build integration

package main_test

import (
  "github.com/lstern/psilibrary/server/models"
  "github.com/lstern/psilibrary/server/repositories"
  "testing"
  "time"
  "errors"
)

var (
	entry4 models.Entry
	entryRepository repositories.EntryRepository
	validator fakeValidator
)

type fakeValidator struct{}

// prepara dados que serão apagados
func init() {
}


//fakes
func (fakeValidator) ValidateEntry(*models.Entry) (error) {
	return errors.New("fake")
}

func createObject() models.Entry {
	return models.Entry{
		Abstract: "A abstract",
		Author:"Leonardo Stern",
		Title:"Testing",
		Content:"Dummy content",
		EntryType: models.EntryType{ID: 2},
		Journal:"Some journal",
		PublishDate:time.Now(),
	}
}

// teste de um insert simples
func Test_create_ok(t *testing.T) {
	obj := createObject()
	id, err := entryRepository.Create(&obj)

	if err != nil || id < 1 {
		t.Error("Registro não foi inserido" )
	}
}

func Test_validate_without_title_fail(t *testing.T){
	obj := createObject()
	obj.Title = ""

	err := entryRepository.ValidateEntry(&obj)

	if err == nil {
		t.Error("Registro não deveria ser inserido")
	}
}

func Test_validate_without_abstract_fail(t *testing.T){
	obj := createObject()
	obj.Abstract = ""

	err := entryRepository.ValidateEntry(&obj)

	if err == nil {
		t.Error("Registro não deveria ser inserido")
	}	
}

func Test_check_if_validation_is_called_on_insert(t *testing.T) {
	obj := createObject()
	repo := repositories.MakeEntryRepository(validator)

	_, err := repo.Create(&obj)

	if err == nil || err.Error() != "fake" {
		t.Error("Método de validação não foi acionado")
	} 
}

func Test_check_invalid_entrytype(t *testing.T) {
	obj := createObject()
	obj.EntryType.ID = -9

	err := entryRepository.ValidateEntry(&obj)

	if err == nil {
		t.Error("Validação deveria falhar")
	}
}

func Test_get_entryType(t *testing.T){
	repo := repositories.MakeEntryTypeRepository(nil)
	obj := createObject()
	entry_type, err := repo.GetById(obj.EntryType.ID)

	if err != nil {
		t.Error(err)
	}

	if entry_type.ID != obj.EntryType.ID{
		t.Error("Erro ao buscar tipo de registro")
	}
}

func Test_check_invalid_category(t *testing.T) {
	obj := createObject()
	obj.Categories = []models.Category{
		models.Category{ID: 2},
		models.Category {ID: -1},
	}

	err := entryRepository.ValidateEntry(&obj)

	if err == nil {
		t.Error("Validação deveria falhar")
	}
}


func Test_validation_is_called_on_update(t *testing.T){
	obj := createObject()
	repo := repositories.MakeEntryRepository(validator)

	err := repo.Update(&obj)

	if err == nil || err.Error() != "fake" {
		t.Error("Método de validação não foi acionado")
	} 
}

func Test_update_ok(t *testing.T){
	obj := createObject()
	repo := repositories.MakeEntryRepository(nil)

	id, _ := repo.Create(&obj)
	obj.EntryId = id
	obj.Title = "Updated"

	err := repo.Update(&obj)

	if err != nil {
		t.Error("Erro atualizando registro: " + err.Error())
	}
}

func Test_update_invalid_id(t *testing.T) {
	obj := createObject()
	obj.EntryId = 999999

	repo := repositories.MakeEntryRepository(nil)
	err := repo.Update(&obj)	

	if err == nil {
		t.Error("Deveria dar erro ao atualizar o registro")
	}
}

func Test_test_select_ok(t *testing.T) {
	obj := createObject()
	repo := repositories.MakeEntryRepository(nil)

	id, _ := repo.Create(&obj)

	selected, err := repo.GetById(id)

	if err != nil || selected.Title != obj.Title {
		t.Error("Erro ao recuperar registro" + err.Error())
	}
}

//do a complete object compare
// insert /update categories 