package main
import  (
	"errors"
	"testing"
	//"strings"
	//"database/sql"
	. "github.com/lstern/psilibrary/server/models"
	"github.com/lstern/psilibrary/server/repositories"
    
    "github.com/DATA-DOG/go-sqlmock"
)


type fakeEntryValidator struct{
}

var (
	entry1 Entry
	entry2 Entry
	entry3 Entry
	entryValidator fakeEntryValidator
	entryRepo repositories.EntryRepository
)

func init(){
	entry1 = Entry {
		ID: 1,
		Title: "Parent",
		Abstract: "Fake abstract",
	}	
	entry2 = Entry {
		ID: 2,
		Title: "Entry 2",
		Abstract: "another abstract",
	}	
}

/////////
// FAKES 
/////////
/*
// fake para recuperar categoria por ID
func (fakeCategoryValidator) GetCategoryById(id int, mydb *sql.DB) (*Category, error) {
	switch id{
		case 1: return &cat1,nil
		case 2: return &cat2,nil
		case 3: return &invalidCat,nil
	}

	return nil,errors.New("Categoria não encontrada")
}
*/
// fake para validar categoria
func (fakeEntryValidator) ValidateEntry(entry *Entry) (bool, string, error){
	if entry.ID == entry1.ID {
		return false, "", errors.New("Registro inválido")
	}

	return true, "", nil
}

/////////
// TESTS
/////////

// insert new entry (ok)
func TestCreatingNewEntry(t *testing.T) {
	db, mock, err := sqlmock.New()
	entryRepo.DB = db
	entryRepo.Validator = entryValidator

	mock.ExpectExec("^insert into Entry .+$").WithArgs(entry1.Abstract, entry1.Author, entry1.Content,
	entry1.EntryType.ID, entry1.Journal, entry1.PublishDate, entry1.Title).WillReturnResult(sqlmock.NewResult(0, 1))

	_, err = entryRepo.Create(&entry1)

	if err == nil{
		err =  mock.ExpectationsWereMet()
  	}

	if err != nil{
      t.Error("Erro ao inserir Registro: ", err.Error())
	}
}

// valida insert e edit
func TestCheckOkEntry(t *testing.T){
	entryRepo.Validator = entryValidator
	b,_,_ := entryRepo.ValidateEntry(&entry1)


	if b!= true {
		t.Error("Erro ao validar registro. Ele deveria ser ok")
	}
}

/*
func TestCheckInvalidCategory(t *testing.T){
	b,msg,_ := entryRepo.ValidateEntry(&entry2, entryValidator)

	if b == true || !strings.Contains(msg, "categoria") { 
		t.Error("Erro ao validar registro. Ele deveria ser inválido")
	}
}

func TestCheckInvalidEntryType(t *testing.T){
	b, _, msg := entryRepo.ValidateEntry(&entry3, entryValidator)

	if b == true || !strings.Contains(msg, "tipo de registro") { 
		t.Error("Erro ao validar registro. Ele deveria ser inválido")
	}
}

func TestUpdateValidEntry(t *testing.T){
	db, mock, err := sqlmock.New()

	mock.ExpectExec("^update Entry .+$").WithArgs(entry1.Abstract, entry1.Author, entry1.Content,
	entry1.EntryTypeId, entry1.Journal, entry1.PublishDate, entry1.Title, entry1.ID).WillReturnResult(sqlmock.NewResult(0, 1))

	err = entryRepo.UpdateEntry(&entry1, db, entryValidator)

	if (err != nil){
		t.Error("Erro ao atualizar um registro válido")
	}
}

func TestUpdateInvalidEntry(t *testing.T){
	db, _, _ := sqlmock.New()

	err := entryRepo.UpdateEntry(&entry2, db, entryValidator)

	if (err == nil){
		t.Error("Erro era esperado ao atualizar um registro inválido")
	}
}

func TestUpdateUnknownEntry(t *testing.T){
	db, mock, _ := sqlmock.New()

	mock.ExpectExec("^update Entry .+$").WithArgs(entry3.Abstract, entry3.Author, entry3.Content,
	entry3.EntryTypeId, entry3.Journal, entry3.PublishDate, entry3.Title, entry3.ID).WillReturnResult(sqlmock.NewResult(0, 0))

	err := entryRepo.UpdateEntry(&entry3, db, entryValidator)

	if (err == nil){
		t.Error("Erro era esperado ao atualizar um registro inexistente")
	}
}

func TestListingAllEntries(t *testing.T) {
	db, mock, err := sqlmock.New()

	rows := sqlmock.NewRows([]string{"ID", "Abstract", "Author", "Content", "EntryTypeId", "Journal", 
		"PublishDate", "Title"}).AddRow(entry3.ID, entry3.Abstract, entry3.Author, entry3.Content, 
			entry3.EntryTypeId, entry3.Journal, entry3.PublishDate, entry3.Title).AddRow(entry1.ID, 
			entry1.Abstract, entry1.Author, entry1.Content,	entry1.EntryTypeId, entry1.Journal,
			entry1.PublishDate, entry1.Title)

	mock.ExpectQuery("^select .+$").WillReturnRows(rows)

	entries, err := entryRepo.ListEntries(db)

	if err == nil{
		err =  mock.ExpectationsWereMet()
  	}

  	if entries == nil{
  		err = errors.New("Nenhuma publicação retornada")
  	}

  	if err == nil && entries[0].ID != 1{
  		err = errors.New("Erro ao retornar publicações")
  	}

	if err != nil{
      t.Error("Erro ao listar publicações: ", err.Error())
      return
	}

}

// Get existing category
func TestGettingAnEntry(t *testing.T) {
	db, mock, err := sqlmock.New()

	rows := sqlmock.NewRows([]string{"ID", "Abstract", "Author", "Content", "EntryTypeId", "Journal", 
		"PublishDate", "Title"}).AddRow(entry3.ID, entry3.Abstract, entry3.Author, entry3.Content, 
			entry3.EntryTypeId, entry3.Journal, entry3.PublishDate, entry3.Title)

	mock.ExpectQuery("^select .+$").WithArgs(3).WillReturnRows(rows)

	entry, err := entryRepo.GetEntryById(3, db)

	if err == nil{
		err =  mock.ExpectationsWereMet()
  	}

  	if err == nil && entry == nil{
  		err = errors.New("Nenhum registro retornado")
  	}

  	if err == nil && entry.ID != 3{
  		err = errors.New("Dados do registro diferente do esperado")
  	}

  	if err != nil{
      t.Error("Erro ao recuperar registro: ", err.Error())
      return
	}
}

// Get existing category
func TestGettingAnUnknownEntry(t *testing.T) {
	db, mock, err := sqlmock.New()

	rows := sqlmock.NewRows([]string{"ID", "Abstract", "Author", "Content", "EntryTypeId", "Journal", 
		"PublishDate", "Title"})

	mock.ExpectQuery("^select .+$").WithArgs(3).WillReturnRows(rows)

	entry, err := entryRepo.GetEntryById(900, db)

	if err == nil{
		err =  mock.ExpectationsWereMet()
  	}

  	if err == nil && entry != nil{
  		err = errors.New("Registro retornado, mas deveria retornar vazio")
  	}

  	if err != nil{
      t.Error("Erro ao recuperar registro: ", err.Error())
      return
	}
}


// delete existing entry
func TestDeletingAnEntry(t *testing.T) {
	db, mock, err := sqlmock.New()

	mock.ExpectExec("^delete .+$").WithArgs(entry2.ID).WillReturnResult(sqlmock.NewResult(0, 1))

	err = entryRepo.DeleteEntry(entry2.ID, db)

	if err == nil{
		err =  mock.ExpectationsWereMet()
  	}

  	if err != nil{
      t.Error("Erro ao apagar registro: ", err.Error())
	}
}

// delete unknown entry
func TestDeletingAnInvalidEntry(t *testing.T) {
	db, mock, err := sqlmock.New()

	mock.ExpectExec("^delete .+$").WithArgs(900).WillReturnResult(sqlmock.NewResult(0, 1))

	err = entryRepo.DeleteEntry(900, db)

	if err == nil{
		err =  mock.ExpectationsWereMet()
  	}

  	if err == nil{
      t.Error("Era esperado erro ao tentar apagar um registro inexistente")
	}
}

// check is routes are ok
func TestCheckEntryMethodsRoutes(t *testing.T){
	router := NewRouter()

	if router.Get("EntryIndex") == nil {
		t.Error("rota de lista de publicação não está registrada")
	}

	if router.Get("EntryCreate") == nil {
		t.Error("rota de criação de publicação não está registrada")
	}

	if router.Get("EntryUpdate") == nil {
		t.Error("rota de atualização de publicação não está registrada")
	}

	if router.Get("EntryDelete") == nil {
		t.Error("rota de remoção de publicação não está registrada")
	}

	if router.Get("EntryGet") == nil {
		t.Error("rota para recuperar dados de publicação não está registrada")
	}	

}
*/