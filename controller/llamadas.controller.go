package controller

import (
	"encoding/json"
	"fmt"
	"llamadas/modelos"
	"llamadas/setup"
	"net/http"
	"strconv"
)

func GetAllLlamadas(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()

	paginaStr := query.Get("pagina")
	pagina, err := strconv.Atoi(paginaStr)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	offset := (pagina - 1) * 10
	rows, err := setup.DBConn.Query(fmt.Sprintf("SELECT * FROM llamadas limit 10 OFFSET %d", offset))

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	defer rows.Close()
	var llamadas []modelos.Llamada

	for rows.Next() {
		var llamada modelos.Llamada
		err := rows.Scan(&llamada.ID, &llamada.FechaLlamada, &llamada.DuracionLlamada, &llamada.Remitente, &llamada.Destinatario)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}
		llamadas = append(llamadas, llamada)
	}

	jsonData, err := json.Marshal(llamadas)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonData)

}
