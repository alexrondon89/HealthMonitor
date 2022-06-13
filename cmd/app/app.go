package app

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"

	localdb "HealthMonitor/platform/db/local"
	"HealthMonitor/server"
	"HealthMonitor/server/service/client"
	"HealthMonitor/server/service/client/elastic"
	"HealthMonitor/server/service/client/postgresClient"
	"HealthMonitor/server/service/client/postgresPool"
	"HealthMonitor/server/service/client/postgresPromise"
	"HealthMonitor/server/service/client/redis"
	"HealthMonitor/server/service/client/serviceUrl"
	"HealthMonitor/server/service/doctormonitor"
	"HealthMonitor/server/service/repository/local"
)

func Start() {
	clients := make(map[string]client.Client)
	clients["serviceUrl"] = serviceUrl.New()
	clients["redisClient"] = redis.New()
	clients["elasticsearchClient"] = elastic.New()
	clients["postgresPromiseClient"] = postgresPromise.New()
	clients["postgresClient"] = postgresClient.New()
	clients["postgresPool"] = postgresPool.New()

	database := localdb.New()
	localRepository := local.New(database)

	registrator := doctormonitor.NewRegistrator(localRepository)
	checker := doctormonitor.NewChecker(clients, localRepository)
	handler := server.New(registrator, checker)

	r := mux.NewRouter()
	r.HandleFunc("/registration", handler.ResourceRegister).Methods(http.MethodPost)
	r.HandleFunc("/health", handler.HealthCheck).Methods(http.MethodGet)
	log.Fatal(http.ListenAndServe("localhost:8080", r))

}
