package workflowPattern

// Settings is simple struct
type Settings struct {
	Port int
	Host string
}

var instance *Settings

// GetInstance returns a single instance of the settings
func GetInstance() *Settings {
	if instance == nil {
		instance = &Settings{} // <--- NOT THREAD SAFE
	}
	return instance
}

/**

//Client
import (
    "fmt"
    Settings "../Singleton"
)

settings := Settings.GetInstance()

settings.Host = "192.168.100.1"
settings.Port = 33

settings1 := Settings.GetInstance()
//settings1.Port is 33
*/
