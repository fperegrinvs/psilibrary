package main
import  (
	"errors"
	"testing"
	. "psilibrary/server/models"
	"psilibrary/server/repositories"
    
    "github.com/DATA-DOG/go-sqlmock"
)


type fakeCategoryRepository struct{}

type fakeCategoryValidator struct{}

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
}

func (fakeCategoryRepository) GetCategoryById(id int) (*Category, error) {
	switch id{
		case 1: return &cat1,nil
		case 2: return &cat2,nil
		case 3: return &cat3,nil
	}

	return nil,errors.New("Categoria não encontrada")
}

func (fakeCategoryValidator) ValidateCategory(category *Category, getter repositories.CategoryGetter) (bool, error){
	panic("Not implemented")
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

func TestCheckCategory(t *testing.T){
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
	t.Error("Need to implement Test")
}

func TestGettingACategory(t *testing.T) {
	t.Error("Need to implement Test")
}

func TestGetInvalidCategory(t *testing.T){
	t.Error("Need to implement Test")
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
