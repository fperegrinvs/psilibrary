package controllers

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
	"github.com/lstern/psilibrary/server/medline"
)

func ImportMedline(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		panic(err)
	}
	if err := r.Body.Close(); err != nil {
		panic(err)
	}

	var	ml medline.Medline;
	xml := string(body[:])
	entries, err := ml.InsertFromXml(xml);

	var importResult medline.ImportResult;
	importResult.Count = len(entries);

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	addCors(w, r)

	if err != nil {
		importResult.Error = err.Error()
	}

	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(importResult); err != nil {
		panic(err)
	}

	return
}
