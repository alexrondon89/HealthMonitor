package app

import (
	"HealthMonitor/server"
	"HealthMonitor/server/service/client/elastic"
	"HealthMonitor/server/service/client/postgresPromise"
	"HealthMonitor/server/service/client/redis"
	"HealthMonitor/server/service/client/serviceUrl"
	"HealthMonitor/server/service/doctormonitor"
	"HealthMonitor/server/service/repository/local"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func Start() {

	elasticCli := elastic.New()
	postgresPromiseCli := postgresPromise.New()
	redisCli := redis.New()
	serviceUrlCli := serviceUrl.New()
	localRepository := local.New()
	srv := doctormonitor.New(elasticCli, postgresPromiseCli, redisCli, serviceUrlCli, localRepository)
	handler := server.New(srv)

	r := mux.NewRouter()
	r.HandleFunc("/registration", handler.ResourceRegister).Methods(http.MethodPost)
	r.HandleFunc("/health", handler.HealthCheck).Methods(http.MethodGet)
	log.Fatal(http.ListenAndServe("localhost:8080", r))

}
