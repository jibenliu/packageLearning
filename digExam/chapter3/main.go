package main

import (
	"digExam/chapter3/controllers"
	"digExam/chapter3/mgo"
	"digExam/chapter3/services/impl"
	"go.mongodb.org/mongo-driver/mongo"
	"go.uber.org/dig"
	"net/http"
)

type ServerConfig struct {
	Host string
	Port string
}

type Router struct {
	routes *map[string]http.HandlerFunc
}

type Server struct {
	serverConfig *ServerConfig
	router       *Router
}

func (svc *Server) Run() {
	for path := range *svc.router.routes {
		http.HandleFunc(path, (*svc.router.routes)[path])
	}
	http.ListenAndServe(svc.serverConfig.Host+":"+svc.serverConfig.Port, nil)
}

func NewMogConfig() *mgo.Config {
	return &mgo.Config{
		Host:     "192.168.139.140",
		Port:     "27017",
		Username: "admin",
		Password: "123456",
		Database: "test",
	}
}

func NewDB(mgoCfg *mgo.Config) *mongo.Database {
	Db, _ := mgo.ConnectDatabase(mgoCfg)
	return Db
}
func NewPersonService(Db *mongo.Database) *impl.Person {
	return &impl.Person{
		Db: Db,
	}
}

func NewPersonController(personService *impl.Person) *controllers.Person {
	return &controllers.Person{
		Service: *personService,
	}
}

func NewRouter(personController *controllers.Person) *Router {
	routes := make(map[string]http.HandlerFunc)
	routes["/person"] = personController.FindAll
	return &Router{
		&routes,
	}
}

func NewServerConfig() *ServerConfig {
	return &ServerConfig{
		Host: "127.0.0.1",
		Port: "8080",
	}
}

func NewServer(svcCfg *ServerConfig, router *Router) *Server {
	return &Server{
		svcCfg,
		router,
	}
}

func BuildContainer() *dig.Container {
	container := dig.New()
	container.Provide(NewMogConfig)
	container.Provide(NewDB)
	container.Provide(NewPersonService)
	container.Provide(NewPersonController)
	container.Provide(NewRouter)
	container.Provide(NewServerConfig)
	container.Provide(NewServer)
	return container
}

func main() {
	container := BuildContainer()
	err := container.Invoke(func(server *Server) {
		server.Run()
	})
	if err != nil {
		panic(err)
	}
}
