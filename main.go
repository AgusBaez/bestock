package main

import (
	"log"

	"go.uber.org/fx"

	"github.com/AgusBaez/BeStock/settings"
)

func main() {
	app := fx.New(
		fx.Provide( //Tenemos que darle las funciones que llevan el struct,automaticamente devuelve un error en caso de falla
			settings.New,
		),
		fx.Invoke( //comando necesario para ejecutar antes que la aplicacion
			func(s *settings.Settings) {
				log.Println(s)
			},
		),
	)
	app.Run()
}
