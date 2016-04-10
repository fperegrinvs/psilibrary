package medline

import  (
	"time"
	"testing"
    "github.com/lstern/psilibrary/server/models"
)

var ( 
	ml Medline
)

func TestReadXML(t *testing.T){
	xml := ml.ReadXML();

	if len(xml) == 0 {
		t.Error("Erro ao ler xml")
	}
}

func TestParseXmlHaveOneElement(t *testing.T) {
	xml := ml.ReadXML();
	result := ml.ParseXML(xml);

	if len(result.PubmedArticles) != 1 {
		t.Error("O xml tem 1 regisro");
	}

	entry := result.PubmedArticles[0];
	if entry.MedlineCitation.Pmid.XCDATA != "26886152" {
		t.Error("Erro ao processar id do registro");
	}
}

func TestInsertEntry(t *testing.T) {
	xml := ml.ReadXML();
	result := ml.ParseXML(xml).PubmedArticles[0];

	article := result.MedlineCitation;
	entry := ml.ConvertArticle(article)

	reference := new(models.Entry);
	reference.Title = "Need for cognition moderates paranormal beliefs and magical ideation in inconsistent-handers.";
    reference.PublishDate = time.Date(2016, time.May, 1, 0, 0, 0, 0, time.Local);
    reference.MedlineId = "26886152";
 
	if (entry.Title != reference.Title || entry.PublishDate != reference.PublishDate || entry.MedlineId != reference.MedlineId){
		t.Error("Erro ao converter artigo");
	}
}
