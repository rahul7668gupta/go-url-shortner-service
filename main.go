package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/rahul7668gupta/go-url-shortner/pkg/constants"
	"github.com/rahul7668gupta/go-url-shortner/pkg/logger"
	"github.com/rahul7668gupta/go-url-shortner/pkg/redis"
	"github.com/rahul7668gupta/go-url-shortner/pkg/repository"
	"github.com/rahul7668gupta/go-url-shortner/pkg/service"
)

func main() {

	// Read port from env
	port := os.Getenv(constants.PORT)
	if len(port) == 0 {
		log.Panic("PORT env not found")
	}
	// Init logger
	loggerObj := logger.InitLogger()
	// Init redis client
	rdb, _ := redis.InitRedisClient() // error handling done inside this func already

	// Initialise repo
	shortnerRepo := repository.NewShortnerRepository(rdb, loggerObj)
	// Initialise service
	urlShortnerService := service.NewUrlShortnerService(shortnerRepo, loggerObj)

	// Initialise router
	routerSrv := mux.NewRouter().StrictSlash(true)

	routerSrv.HandleFunc("/shorten", urlShortnerService.ShortenURL).Methods("POST")
	routerSrv.HandleFunc("/r/{code}", urlShortnerService.Redirect).Methods("GET")
	routerSrv.HandleFunc("/metrics", urlShortnerService.Metrics).Methods("GET")

	loggerObj.Sugar().Infof("Server listening on port %s", port)
	log.Fatal(http.ListenAndServe(port, routerSrv))
}
