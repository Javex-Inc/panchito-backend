package model

type Status string

const (
	PENDING   Status = "Pendente"
	PREPARING Status = "Preparando"
	READY     Status = "Pronto"
	UNDERWAY  Status = "A caminho"
	DELIVERED Status = "Entregue"
	CANCELED  Status = "Cancelado"
)
