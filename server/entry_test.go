//+build integration

package main_test

import (
  "github.com/lstern/psilibrary/server"
  "github.com/lstern/psilibrary/server/models"
  "github.com/lstern/psilibrary/server/repositories"
  "errors"
  "testing"
  "time"
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

func CompareEntry(obj1 *models.Entry, obj2 *models.Entry) (bool) {
	return !(obj1.Title != obj2.Title || obj1.Abstract != obj2.Abstract || obj1.Author != obj2.Author ||
		obj1.Content != obj2.Content || obj1.EntryId != obj2.EntryId ||	obj1.Journal != obj2.Journal)
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

func Test_check_duplicated_category(t *testing.T) {
	obj := createObject()
	obj.Categories = []models.Category{
		models.Category{ID: 2},
		models.Category {ID: 2},
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

func Test_select_ok(t *testing.T) {
	obj := createObject()
	repo := repositories.MakeEntryRepository(nil)

	id, _ := repo.Create(&obj)
	obj.EntryId = id

	selected, err := repo.GetById(id)

	if err != nil || !CompareEntry(selected, &obj) {
		t.Error("Erro ao recuperar registro" + err.Error())
	}
}

func Test_insert_category_relation(t *testing.T){
	obj := createObject()
	entryRepo := repositories.MakeEntryRepository(nil)
	entry_id, _ := entryRepo.Create(&obj)

	cat := models.Category{Name: "Just testing"}
	catRepo := repositories.MakeCategoryRepository(nil, nil)
	cat_id, _ := catRepo.Create(&cat)

	err := entryRepository.InsertEntryCategory(entry_id, cat_id)

	if err != nil {
		t.Error("Categoria não foi inserida")
	}
}

func Test_select_entry_categories(t *testing.T) {
	entryRepo := repositories.MakeEntryRepository(nil)
	cats, err := entryRepo.GetEntryCategories(1)

	if err != nil || len(cats) != 2  || cats[0] != 2 || cats[1] != 3 {
		t.Error("Erro ao recuperar categorias")
	}
}

func Test_insert_with_categories(t *testing.T) {
	obj := createObject()
	obj.Categories = []models.Category{models.Category{ID:2},models.Category{ID:3}}

	entryRepo := repositories.MakeEntryRepository(nil)
	entry_id, _ := entryRepo.Create(&obj)

	cats, err := entryRepo.GetEntryCategories(entry_id)

	if err != nil || len(cats) != 2 || cats[0] != 2 || cats[1] != 3{
		t.Error("Erro ao inserir registro")
	}
}

func Test_update_with_categories(t *testing.T) {
	obj := createObject()

	entryRepo := repositories.MakeEntryRepository(nil)
	entry_id, _ := entryRepo.Create(&obj)

	obj.EntryId = entry_id
	obj.Categories = []models.Category{models.Category{ID:2},models.Category{ID:3}}

	err := entryRepo.Update(&obj)

	cats, err := entryRepo.GetEntryCategories(entry_id)

	if err != nil || len(cats) != 2 || cats[0] != 2 || cats[1] != 3{
		t.Error("Erro ao atualizar registro")
	}
}

func Test_update_replacing_categories(t *testing.T) {
	obj := createObject()
	obj.Categories = []models.Category{models.Category{ID:2},models.Category{ID:3}}

	entryRepo := repositories.MakeEntryRepository(nil)
	entry_id, _ := entryRepo.Create(&obj)

	obj.EntryId = entry_id
	obj.Categories = []models.Category{models.Category{ID:4}}

	err := entryRepo.Update(&obj)

	cats, err := entryRepo.GetEntryCategories(entry_id)

	if err != nil || len(cats) != 1 || cats[0] != 4{
		t.Error("Erro ao atualizar registro")
	}
}

func Test_delete_entry(t *testing.T) {
	obj := createObject()

	entryRepo := repositories.MakeEntryRepository(nil)
	entry_id, _ := entryRepo.Create(&obj)

	entryRepo.Delete(entry_id)
	entry, err := entryRepo.GetById(entry_id)

	if err == nil || (entry != nil && entry.EntryId > 0) {
		t.Error("Falha ao apagar registro")
	}
}

func Test_delete_invalid_entry(t *testing.T) {
	entryRepo := repositories.MakeEntryRepository(nil)
	err := 	entryRepo.Delete(394949)

	if err == nil {
		t.Error("Era esperado erro ao tentar apagar um registro inexistente")
	}

}

func Test_select_including_cats(t *testing.T){
	obj := createObject()
	obj.Categories = []models.Category{models.Category{ID:2},models.Category{ID:3}}

	repo := repositories.MakeEntryRepository(nil)
	obj.EntryId, _ = repo.Create(&obj)

	selected, err := repo.GetById(obj.EntryId)	

	if selected.Categories == nil || len(selected.Categories) != len(obj.Categories) || selected.Categories[1].Name != "ESP"  {
		t.Error("Erro ao selecionar registro com categorias" + err.Error())
	}
}

func Test_select_include_entryType(t *testing.T){
	obj := createObject()

	repo := repositories.MakeEntryRepository(nil)
	obj.EntryId, _ = repo.Create(&obj)

	selected, err := repo.GetById(obj.EntryId)	

	if selected == nil || selected.EntryType.Name != "Livro"   {
		t.Error("Erro ao inserir registro com tipo de registro", err)
	}
}

func Test_check_publishDate(t *testing.T) {
	obj := createObject()

	repo := repositories.MakeEntryRepository(nil)
	obj.EntryId, _ = repo.Create(&obj)

	selected, err := repo.GetById(obj.EntryId)	
	
	if selected.PublishDate.Format("2006-01-02") != obj.PublishDate.Format("2006-01-02")   {
		t.Error("Erro ao inserir registro com tipo de data de publicação", err)
	}
}

func Test_list(t *testing.T) {
	repo := repositories.MakeEntryRepository(nil)
	list, err := repo.List()

	if err != nil || list == nil || len(*list) < 2 || (*list)[0].Title != "Artigo dummy" {
		t.Error("Falha ao listar registros", err)
	}
}

// check is routes are ok
func Test_CheckEntryMethodsRoutes(t *testing.T){
	router := main.NewRouter()

	if router.Get("EntryIndex") == nil {
		t.Error("rota de lista de registros não está registrada")
	}

	if router.Get("EntryCreate") == nil {
		t.Error("rota de criação de registros não está registrada")
	}

	if router.Get("EntryUpdate") == nil {
		t.Error("rota de atualização de registro não está registrada")
	}

	if router.Get("EntryGet") == nil {
		t.Error("rota para recuperar dados de registro não está registrada")
	}	
}

// list