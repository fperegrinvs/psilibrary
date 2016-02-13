package main
import  (
	"errors"
	"testing"
	"database/sql"
	. "psilibrary/server/models"
	"psilibrary/server/repositories"
    
    "github.com/DATA-DOG/go-sqlmock"
)


type fakeCategoryRepository struct{}

type fakeCategoryValidator struct{
	validatorCalled bool
}

var (
	cat1 Category
	cat2 Category
	cat3 Category
	fakeRepo fakeCategoryRepository
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
	cat3 = Category {
		ID: 3,
		Name: "Cat3",
		ParentId: 900,
	}

	fakeValidator.validatorCalled = false	
}

func (fakeCategoryRepository) GetCategoryById(id int, mydb *sql.DB) (*Category, error) {
	switch id{
		case 1: return &cat1,nil
		case 2: return &cat2,nil
		case 3: return &cat3,nil
	}

	return nil,errors.New("Categoria não encontrada")
}

func (fakeCategoryValidator) ValidateCategory(category *Category, getter repositories.CategoryGetter) (bool, error){
	fakeValidator.validatorCalled = true
	return true, nil
}


func TestCreatingNewCategory(t *testing.T) {
	db, mock, err := sqlmock.New()

	mock.ExpectExec("^insert into Category .+$").WithArgs(cat2.Name, cat2.ParentId).WillReturnResult(sqlmock.NewResult(0, 1))

	_, err = repo.CreateCategory(&cat2, db, fakeValidator)

	if err == nil{
		err =  mock.ExpectationsWereMet()
  	}

	if err != nil{
      t.Error("Erro ao inserir categoria: %s", err.Error())
      return
	}
}

func TestCreateValidatorCalled(t *testing.T){
	db, mock, _ := sqlmock.New()

	mock.ExpectExec("^.+$").WithArgs(cat2.Name, cat2.ParentId).WillReturnResult(sqlmock.NewResult(0, 1))

	repo.CreateCategory(&cat2, db, fakeValidator)

	if fakeValidator.validatorCalled == false{
      t.Error("Validação não foi ativada")
	}
}

func TestCheckCategoryNoParent(t *testing.T){
	b,_ := repo.ValidateCategory(&cat1, fakeRepo)

	if b != true{
		t.Error("Erro ao validar categoria")
	}
}

func TestCheckCategoryExistingParent(t *testing.T){
	b,_ := repo.ValidateCategory(&cat2, fakeRepo)

	if b != true{
		t.Error("Erro ao validar categoria")
	}
}

func TestCheckCategoryInvalidParent(t *testing.T){
	b,_ := repo.ValidateCategory(&cat3, fakeRepo)

	if b == true{
		t.Error("Erro ao validar categoria, deveria ser falso")
	}
}

func TestListingAllCategories(t *testing.T) {
	db, mock, err := sqlmock.New()

	rows := sqlmock.NewRows([]string{"ID", "Name", "ParentId"}).AddRow(cat1.ID, cat2.Name, cat3.ParentId)

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
      t.Error("Erro ao listar categorias: %s", err.Error())
      return
	}
}

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

func TestGetInvalidCategory(t *testing.T){
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
      t.Error("Erro ao recuperar categoria: %s", err.Error())
      return
	}
}


func TestUpdatingACategory(t *testing.T) {
	t.Error("Need to implement Test")
}

func TestDeletingACategory(t *testing.T) {
	t.Error("Need to implement Test")
}

func TestCheckIfCategoryIsUsed(t *testing.T) {
	t.Error("Need to implement Test")
}

func TestCheckIfCategoryIsUsedInvalid(t *testing.T) {
	t.Error("Need to implement Test")
}


func TestCheckCategoryMethodsRoutes(t *testing.T){
	t.Error("Need to implement Test")
}
