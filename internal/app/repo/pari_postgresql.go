package repo

type postgresql struct{ tc SQLTxConn }

type PostgreSQL interface {
	ItemsRespository
	ItemDetailsRespository
}

func NewPostgreSQL(txc SQLTxConn) PostgreSQL { return &postgresql{txc} }
