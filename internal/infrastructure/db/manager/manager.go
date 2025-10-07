package manager

import (
	trmpgx "github.com/avito-tech/go-transaction-manager/drivers/pgxv5/v2"
	"github.com/avito-tech/go-transaction-manager/trm/v2"
	"github.com/avito-tech/go-transaction-manager/trm/v2/manager"
	"github.com/jackc/pgx/v5/pgxpool"
)

func NewPgxTrManagerDefaultCtxGetter() *trmpgx.CtxGetter {
	return trmpgx.DefaultCtxGetter
}

func NewPgxTrManager(pool *pgxpool.Pool) trm.Manager {
	return manager.Must(trmpgx.NewDefaultFactory(pool))
}
