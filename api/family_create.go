package api

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

/// SE O FILHO JÁ TEM DOIS PAIS (BLOQUEIA)
/// VOCE NÃO PODE SER PAI DOS SEUS ANTEPASSADOS

func (h *Handler) InsertMember(w http.ResponseWriter, r *http.Request) {

	body, _ := ioutil.ReadAll(r.Body)
	var memberRequest MemberRequest
	json.Unmarshal(body, &memberRequest)

	resp, err := json.Marshal(memberRequest)
	if err != nil {
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(resp)
}
