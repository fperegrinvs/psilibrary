package medline

import (
    "strconv"
    "time"
    "encoding/xml"
    "github.com/lstern/psilibrary/server/medline/generated"
    "github.com/lstern/psilibrary/server/models"
)

// reference https://gist.github.com/kwmt/6135123#file-parsetvdb-go
type Medline struct {

}

type PubmedArticleSet struct {
    XMLName xml.Name `xml:"PubmedArticleSet"`
    generated.TPubmedArticleSet
}

func (m Medline) ParseXML(file string) *PubmedArticleSet {
	bytes := []byte(file);

	a := new (PubmedArticleSet);
	xml.Unmarshal(bytes, a)
	return a;
}

func (m Medline) ConvertArticle(citation *generated.TMedlineCitation) *models.Entry{
    article := citation.Article;
    entry := new(models.Entry);
    entry.Abstract = article.Abstract.AbstractTexts[0].XCDATA;
    entry.Title = article.ArticleTitle.XCDATA;
    entry.Content = "";
    entry.EntryType = *new(models.EntryType);
    entry.EntryType.ID = 2;

    if (article.Journal != nil){
        if (article.Journal.JournalIssue != nil && article.Journal.JournalIssue.PubDate != nil){
            year, _ := strconv.Atoi(article.Journal.JournalIssue.PubDate.Year.String())
            month := time.January;
            day := 1;

            if (article.Journal.JournalIssue.PubDate.XHasElem_Month.Walk() == nil){

                if (len(article.Journal.JournalIssue.PubDate.Month.String()) > 2) {
                    
                    test, err := time.Parse("January", article.Journal.JournalIssue.PubDate.Month.String());
                    if (err == nil){
                        month = test.Month();
                    }               
                }

                if (article.Journal.JournalIssue.PubDate.XHasElem_Day.Walk() == nil){
                    nday, err := strconv.Atoi(article.Journal.JournalIssue.PubDate.Day.String())
                    if (err == nil){
                        day = nday;
                    }
                }
            }


            entry.PublishDate = time.Date(year, month, day, 0, 0, 0, 0, time.Local);
            entry.Journal = article.Journal.Title.String();
        }

        if (article.AuthorList != nil){
            entry.Author = "";
            for i:= 0; i < len(article.AuthorList.Authors); i++  {
                author := article.AuthorList.Authors[i];
                entry.Author += author.LastName.String() + ", " + author.ForeName.String();
                if i + 1 < len(article.AuthorList.Authors) {
                    entry.Author += "; ";
                }
            }
        }
    }

    entry.MedlineId = citation.Pmid.XCDATA;

    return entry;    
}

func (m Medline) ReadXML() string { 
return `<?xml version="1.0"?>
<!DOCTYPE PubmedArticleSet PUBLIC "-//NLM//DTD PubMedArticle, 1st January 2016//EN" "http://www.ncbi.nlm.nih.gov/corehtml/query/DTD/bookdoc_160101.dtd">

<PubmedArticleSet>

<PubmedArticle>
    <MedlineCitation Owner="NLM" Status="In-Data-Review">
        <PMID Version="1">26886152</PMID>
        <DateCreated>
            <Year>2016</Year>
            <Month>03</Month>
            <Day>10</Day>
        </DateCreated>
        <Article PubModel="Print-Electronic">
            <Journal>
                <ISSN IssnType="Electronic">1464-0678</ISSN>
                <JournalIssue CitedMedium="Internet">
                    <Volume>21</Volume>
                    <Issue>3</Issue>
                    <PubDate>
                        <Year>2016</Year>
                        <Month>May</Month>
                    </PubDate>
                </JournalIssue>
                <Title>Laterality</Title>
                <ISOAbbreviation>Laterality</ISOAbbreviation>
            </Journal>
            <ArticleTitle>Need for cognition moderates paranormal beliefs and magical ideation in inconsistent-handers.</ArticleTitle>
            <Pagination>
                <MedlinePgn>228-42</MedlinePgn>
            </Pagination>
            <ELocationID EIdType="doi" ValidYN="Y">10.1080/1357650X.2015.1125914</ELocationID>
            <Abstract>
                <AbstractText>A growing literature suggests that degree of handedness predicts gullibility and magical ideation. Inconsistent-handers (people who use their non-dominant hand for at least one common manual activity) report more magical ideation and are more gullible. The current study tested whether this effect is moderated by need for cognition. One hundred eighteen university students completed questionnaires assessing handedness, self-reported paranormal beliefs, and self-reported need for cognition. Handedness (Inconsistent vs. Consistent Right) and Need for Cognition (High vs. Low) were treated as categorical predictors. Both paranormal beliefs and magical ideation served as dependent variable's in separate analyses. Neither set of tests yielded main effects for handedness or need for cognition. However, there were a significant handedness by need for cognition interactions. Post-hoc comparisons revealed that low, but not high, need for cognition inconsistent-handers reported relatively elevated levels of paranormal belief and magical ideation. A secondary set of tests treating the predictor variables as continuous instead of categorical obtained the same overall pattern.</AbstractText>
            </Abstract>
            <AuthorList CompleteYN="Y">
                <Author ValidYN="Y">
                    <LastName>Prichard</LastName>
                    <ForeName>Eric C</ForeName>
                    <Initials>EC</Initials>
                    <AffiliationInfo>
                        <Affiliation>a Department of Psychology , University of Toledo , Toledo , OH , USA.</Affiliation>
                    </AffiliationInfo>
                </Author>
                <Author ValidYN="Y">
                    <LastName>Christman</LastName>
                    <ForeName>Stephen D</ForeName>
                    <Initials>SD</Initials>
                    <AffiliationInfo>
                        <Affiliation>a Department of Psychology , University of Toledo , Toledo , OH , USA.</Affiliation>
                    </AffiliationInfo>
                </Author>
            </AuthorList>
            <Language>eng</Language>
            <PublicationTypeList>
                <PublicationType UI="D016428">Journal Article</PublicationType>
            </PublicationTypeList>
            <ArticleDate DateType="Electronic">
                <Year>2016</Year>
                <Month>02</Month>
                <Day>17</Day>
            </ArticleDate>
        </Article>
        <MedlineJournalInfo>
            <Country>England</Country>
            <MedlineTA>Laterality</MedlineTA>
            <NlmUniqueID>9609064</NlmUniqueID>
            <ISSNLinking>1357-650X</ISSNLinking>
        </MedlineJournalInfo>
        <CitationSubset>IM</CitationSubset>
        <KeywordList Owner="NOTNLM">
            <Keyword MajorTopicYN="N">Handedness</Keyword>
            <Keyword MajorTopicYN="N">magical ideation</Keyword>
            <Keyword MajorTopicYN="N">need for cognition</Keyword>
            <Keyword MajorTopicYN="N">paranormal beliefs</Keyword>
        </KeywordList>
    </MedlineCitation>
    <PubmedData>
        <History>
            <PubMedPubDate PubStatus="aheadofprint">
                <Year>2016</Year>
                <Month>2</Month>
                <Day>17</Day>
            </PubMedPubDate>
            <PubMedPubDate PubStatus="entrez">
                <Year>2016</Year>
                <Month>2</Month>
                <Day>18</Day>
                <Hour>6</Hour>
                <Minute>0</Minute>
            </PubMedPubDate>
            <PubMedPubDate PubStatus="pubmed">
                <Year>2016</Year>
                <Month>2</Month>
                <Day>18</Day>
                <Hour>6</Hour>
                <Minute>0</Minute>
            </PubMedPubDate>
            <PubMedPubDate PubStatus="medline">
                <Year>2016</Year>
                <Month>2</Month>
                <Day>18</Day>
                <Hour>6</Hour>
                <Minute>0</Minute>
            </PubMedPubDate>
        </History>
        <PublicationStatus>ppublish</PublicationStatus>
        <ArticleIdList>
            <ArticleId IdType="pubmed">26886152</ArticleId>
            <ArticleId IdType="doi">10.1080/1357650X.2015.1125914</ArticleId>
        </ArticleIdList>
    </PubmedData>
</PubmedArticle>
</PubmedArticleSet>
`;
}