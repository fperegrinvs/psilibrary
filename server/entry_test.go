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

func Test_check_is_validation_is_called_on_insert(t *testing.T) {
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
