package main

import (
	"digExam/chapter2/controllers"
	"digExam/chapter2/mgo"
	"digExam/chapter2/services/impl"
	"net/http"
)

type ServerConfig struct {
	Host string
	Port string
}

type Server struct {
	ServerConfig *ServerConfig
	routes       *map[string]http.HandlerFunc
}

func (svc *Server) Run() {
	for path := range *svc.routes {
		http.HandleFunc(path, (*svc.routes)[path])
	}
	http.ListenAndServe(svc.ServerConfig.Host+":"+svc.ServerConfig.Port, nil)
}

func main() {
	mgoCfg := mgo.Config{
		Host:     "192.168.139.140",
		Port:     "27017",
		Username: "admin",
		Password: "123456",
		Database: "test",
	}

	Db, _ := mgo.ConnectDatabase(&mgoCfg)
	personService := impl.Person{
		Db: Db,
	}
	personController := controllers.Person{
		Service: personService,
	}
	routes := make(map[string]http.HandlerFunc)
	routes["/person"] = personController.FindAll
	svcConfig := ServerConfig{
		Host: "127.0.0.1",
		Port: "8080",
	}
	svc := Server{&svcConfig, &routes}
	svc.Run()
}
