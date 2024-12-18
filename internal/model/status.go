package model

type Status string

const (
	PENDING   Status = "Pendente"
	PREPARING        = "Preparando"
	READY            = "Pronto"
	UNDERWAY         = "A caminho"
	DELIVERED        = "Entregue"
	CANCELED         = "Cancelado"
)
