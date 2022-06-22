package models

type Realize struct {
	ID                       string  `json:"_id" bson:"_id, omitempty"`
	AnoMes                   int     `json:"iAnoMes"`
	NichoRealize             int     `json:"iNichoRealize"`
	ItemRealize              int     `json:"iItemRealize"`
	NivelRede                int     `json:"iNivelRede"`
	CgcUnidade               int     `json:"iCgcUnidade"`
	NaturalUnidade           int     `json:"iNaturalUnidade"`
	ObjetivoAcumulado        float64 `json:"nObjetivoAcumulado"`
	RealizadoAcumulado       float64 `json:"nRealizadoAcumulado"`
	PercRealizado            float64 `json:"fPercRealizado"`
	PercEsperado             float64 `json:"fPercEsperado"`
	PercAtingido             float64 `json:"fPercAtingido"`
	Ponto                    float64 `json:"fPonto"`
	GrupoRealize             int     `json:"iGrupoRealize"`
	PesoOficial              float64 `json:"fPesoOficial"`
	PesoReal                 float64 `json:"fPesoReal"`
	CgcUnidadeSuperior       int     `json:"iCgcUnidadeSuperior"`
	CgcVp                    int     `json:"iCgcVp"`
	Avaliacao                int     `json:"iAvaliacao"`
	Referencia               string  `json:"dReferencia"`
	AtualizacaoItemMetaSidem string  `json:"dAtualizacaoItemMetaSidem"`
	IdArquivo                int     `json:"idArquivo"`
}
