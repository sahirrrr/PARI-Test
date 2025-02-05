package repo

type postgresql struct{ tc SQLTxConn }

//nolint:interfacebloat
type PostgreSQL interface {
}

func NewPostgreSQL(txc SQLTxConn) PostgreSQL { return &postgresql{txc} }
