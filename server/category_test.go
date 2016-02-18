package main
import  (
	"errors"
	"testing"
	"database/sql"
	. "psilibrary/server/models"
	"psilibrary/server/repositories"
    
    "github.com/DATA-DOG/go-sqlmock"
)


type fakeCategoryValidator struct{
}

var (
	cat1 Category
	cat2 Category
	invalidCat Category
	fakeValidator fakeCategoryValidator
	repo repositories.CategoryRepository
)

func init(){
	cat1 = Category {
		ID: 1,
		Name: "Parent",
		ParentId: 0,
	}	
	cat2 = Category {
		ID: 2,
		Name: "Cat2",
		ParentId: 1,
	}	
	invalidCat = Category {
		ID: 3,
		Name: "Cat3",
		ParentId: 900,
	}
}

/////////
// FAKES 
/////////

// fake para recuperar categoria por ID
func (fakeCategoryValidator) GetCategoryById(id int, mydb *sql.DB) (*Category, error) {
	switch id{
		case 1: return &cat1,nil
		case 2: return &cat2,nil
		case 3: return &invalidCat,nil
	}

	return nil,errors.New("Categoria não encontrada")
}

// fake para validar categoria
func (fakeCategoryValidator) ValidateCategory(category *Category, getter repositories.CategoryValidator) (bool, error){
	if category.ID == invalidCat.ID {
		return false, errors.New("Categoria inválida")
	}

	return true, nil
}

// fake para validar se uma categoria foi usada
func (fakeCategoryValidator) CheckForUsedCategory(id int, validator repositories.CategoryValidator) (repositories.CategoryCheckResult, error){
	return repo.CheckForUsedCategory(id, validator)
}

// get categories by parentID
func (fakeCategoryValidator) GetCategoriesByParentId(catid int, db *sql.DB)([]*Category, error){
	if catid == 1 {
		cats := []*Category{&cat2}
		return cats, nil
	}

	return nil, nil
}	

// get entries from some category
func (fakeCategoryValidator) GetEntriesByCategoryId(catid int, db *sql.DB)([]*Entry, error)	{
	if catid == 1 {
		entry := new(Entry)
		entries := []*Entry{entry}
		return entries, nil
	}

	return nil, nil
}


/////////
// TESTS
/////////

// insert new category (ok)
func TestCreatingNewCategory(t *testing.T) {
	db, mock, err := sqlmock.New()

	mock.ExpectExec("^insert into Category .+$").WithArgs(cat2.Name, cat2.ParentId).WillReturnResult(sqlmock.NewResult(0, 1))

	_, err = repo.CreateCategory(&cat2, db, fakeValidator)

	if err == nil{
		err =  mock.ExpectationsWereMet()
  	}

	if err != nil{
      t.Error("Erro ao inserir categoria: ", err.Error())
	}
}

// insert new category without parent
func TestCheckCategoryNoParent(t *testing.T){
	b,_ := repo.ValidateCategory(&cat1, fakeValidator)

	if b != true{
		t.Error("Erro ao validar categoria")
	}
}

// check category with valid parent
func TestCheckCategoryExistingParent(t *testing.T){
	b,_ := repo.ValidateCategory(&cat2, fakeValidator)

	if b != true{
		t.Error("Erro ao validar categoria")
	}
}

// check category with invalid parent
func TestCheckCategoryInvalidParent(t *testing.T){
	b,_ := repo.ValidateCategory(&invalidCat, fakeValidator)

	if b == true{
		t.Error("Erro ao validar categoria, deveria ser falso")
	}
}

// Test list all categories
func TestListingAllCategories(t *testing.T) {
	db, mock, err := sqlmock.New()

	rows := sqlmock.NewRows([]string{"ID", "Name", "ParentId"}).AddRow(cat1.ID, cat1.Name, cat1.ParentId)

	mock.ExpectQuery("^select .+$").WillReturnRows(rows)

	cats, err := repo.ListCategories(db)

	if err == nil{
		err =  mock.ExpectationsWereMet()
  	}

  	if cats == nil{
  		err = errors.New("Nenhuma categoria retornada")
  	}

  	if err == nil && cats[0].ID != 1{
  		err = errors.New("Erro ao retornar categoria")
  	}

	if err != nil{
      t.Error("Erro ao listar categorias: ", err.Error())
      return
	}
}

// Get existing category
func TestGettingACategory(t *testing.T) {
	db, mock, err := sqlmock.New()

	rows := sqlmock.NewRows([]string{"ID", "Name", "ParentId"}).AddRow(cat1.ID, cat1.Name, cat1.ParentId)

	mock.ExpectQuery("^select .+$").WithArgs(1).WillReturnRows(rows)

	cats, err := repo.GetCategoryById(1, db)

	if err == nil{
		err =  mock.ExpectationsWereMet()
  	}

  	if err == nil && cats == nil{
  		err = errors.New("Nenhuma categoria retornada")
  	}

  	if err == nil && cats.ID != 1{
  		err = errors.New("Dados da categoria diferente do esperado")
  	}

  	if err != nil{
      t.Error("Erro ao recuperar categoria: %s", err.Error())
      return
	}
}

// Get unknown category
func TestGetUnknownCategory(t *testing.T){
	db, mock, err := sqlmock.New()

	rows := sqlmock.NewRows([]string{"ID", "Name", "ParentId"})

	mock.ExpectQuery("^select .+$").WithArgs(900).WillReturnRows(rows)

	cats, err := repo.GetCategoryById(900, db)

	if err == nil{
		err =  mock.ExpectationsWereMet()
  	}

  	if err == nil && cats != nil{
  		err = errors.New("Categoria retornada, mas deveria retornar vazio")
  	}

  	if err != nil{
      t.Error("Erro ao recuperar categoria: ", err.Error())
      return
	}
}

// update category with valid values
func TestUpdatingACategory(t *testing.T) {
	db, mock, err := sqlmock.New()

	mock.ExpectExec("^update .+$").WithArgs(cat2.Name, cat2.ParentId, cat2.ID).WillReturnResult(sqlmock.NewResult(0, 1))

	err = repo.UpdateCategory(&cat2, db, fakeValidator)

	if err == nil{
		err =  mock.ExpectationsWereMet()
  	}

  	if err != nil{
      t.Error("Erro ao atualizar categoria: ", err.Error())
	}
}

// update unknown category
func TestUpdatingAnUnknownCategory(t *testing.T) {
	db, mock, err := sqlmock.New()

	mock.ExpectExec("^update .+$").WithArgs(cat2.Name, cat2.ParentId, cat2.ID).WillReturnResult(sqlmock.NewResult(0, 0))

	err = repo.UpdateCategory(&cat2, db, fakeValidator)

	if err == nil{
		err =  mock.ExpectationsWereMet()
  	}

  	if err == nil{
      t.Error("Erro era esperado ao atualizar categoria inexistente")
	}
}

// update category with invalid values
func TestUpdatingAnInvalidCategory(t *testing.T) {
	db, _, err := sqlmock.New()

	err = repo.UpdateCategory(&invalidCat, db, fakeValidator)

  	if err == nil{
      t.Error("Erro era esperado ao atualizar categoria inválida")
	}
}

// delete existing category
func TestDeletingACategory(t *testing.T) {
	db, mock, err := sqlmock.New()

	mock.ExpectExec("^delete .+$").WithArgs(cat2.ID).WillReturnResult(sqlmock.NewResult(0, 1))

	_,err = repo.DeleteCategory(cat2.ID, db, fakeValidator)

	if err == nil{
		err =  mock.ExpectationsWereMet()
  	}

  	if err != nil{
      t.Error("Erro ao apagar categoria: ", err.Error())
	}
}

// delete unknown category
func TestDeletingUnkownCategory(t *testing.T) {
	db, mock, err := sqlmock.New()

	mock.ExpectExec("^delete .+$").WithArgs(invalidCat.ID).WillReturnResult(sqlmock.NewResult(0, 0))

	_,err = repo.DeleteCategory(invalidCat.ID, db, fakeValidator)

	if err == nil{
		err =  mock.ExpectationsWereMet()
  	}

  	if err == nil{
      t.Error("Erro era esperado ao tentar apagar categoria inexistente: ")
	}
}

// check unused category is used
func TestCheckIfUnusedCategoryIsUsed(t *testing.T) {
	b,_ := repo.CheckForUsedCategory(invalidCat.ID, fakeValidator)

	if b.Existing == true{
		t.Error("Categoria não é utilizada em lugar nenhum")
	}
}

// check if used category (by another category) is used
func TestCheckIfUsedCategoryIsUsedCategory(t *testing.T) {
	b,_ := repo.CheckForUsedCategory(cat1.ID, fakeValidator)

	if b.Existing == false || b.Categories == nil {
		t.Error("Categoria é utilizada por outras categorias")
	}
}

// check if used (by entry) category is used
func TestCheckIfUsedCategoryIsUsedEntry(t *testing.T) {
	b,_ := repo.CheckForUsedCategory(cat1.ID, fakeValidator)

	if b.Existing == false || b.Entries == nil {
		t.Error("Categoria é utilizada por registros")
	}
}

// delete used category
func TestDeletingUsedCategory(t *testing.T) {
	db, _, _ := sqlmock.New()

	used,_ := repo.DeleteCategory(cat1.ID, db, fakeValidator)

  	if used.Existing == false{
      t.Error("Erro era esperado ao tentar apagar categoria usada em outros lugares")
	}
}

// check is routes are ok
func TestCheckCategoryMethodsRoutes(t *testing.T){
	router := NewRouter()

	if router.Get("CategoryIndex") == nil {
		t.Error("rota de lista de categorias não está registrada")
	}

	if router.Get("CategoryCreate") == nil {
		t.Error("rota de criação de categoria não está registrada")
	}

	if router.Get("CategoryUpdate") == nil {
		t.Error("rota de atualização de categoria não está registrada")
	}

	if router.Get("CategoryDelete") == nil {
		t.Error("rota de remoção de categoria não está registrada")
	}

	if router.Get("CategoryGet") == nil {
		t.Error("rota para recuperar dados de categoria não está registrada")
	}	

}
