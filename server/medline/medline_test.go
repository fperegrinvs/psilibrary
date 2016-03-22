package medline

import  (
	"testing"
	//"github.com/lstern/psilibrary/server/medline/generated"
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
