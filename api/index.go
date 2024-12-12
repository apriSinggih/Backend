package handler

import (
	"net/http"

	"github.com/TrinityKnights/Backend/config"
	_ "github.com/TrinityKnights/Backend/docs"
)

// Handler is main function to run the application in vercel function
func Handler(w http.ResponseWriter, r *http.Request) {
	viper := config.NewViper()
	log := config.NewLogrus(viper)
	db := config.NewDatabase(viper, log)
	redis := config.NewRedisClient(viper, log)
	jwt := config.NewJWT(viper)
	validate := config.NewValidator()
	midtrans := config.NewMidtransClient(viper)
	app, log := config.NewEcho()

	err := config.Bootstrap(&config.BootstrapConfig{
		DB:       db,
		Cache:    redis,
		App:      app,
		Log:      log,
		Validate: validate,
		JWT:      jwt,
		Midtrans: midtrans,
	})
	if err != nil {
		log.Fatalf("Failed to bootstrap application: %v", err)
	}

	app.ServeHTTP(w, r)
}