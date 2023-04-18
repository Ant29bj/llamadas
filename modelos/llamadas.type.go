package modelos

type Llamada struct {
	ID              int    `json:"id"`
	FechaLlamada    string `json:"fecha_llamada"`
	DuracionLlamada string `json:"duracion_llamada"`
	Remitente       string `json:"remitente"`
	Destinatario    string `json:"destinatario"`
}
