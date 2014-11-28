package middlewares

import (
	"github.com/go-martini/martini"
	"github.com/venjiang/mons/services"
	"log"
	"os"
)

func CoreService() martini.Handler {

	return func(c martini.Context) {
		svc, err := services.InitCoreServie()

		if err != nil {
			panic(err)
		}
		defer svc.DbMap.Db.Close()
		c.Map(svc)
		svc.DbMap.TraceOn("[gorp]", log.New(os.Stdout, "mons:", log.Lmicroseconds))
		c.Next()
	}
}
