package app

type Config struct {
	Host     string
	Port     int
	User     string
	Pass     string
	ServerId int

	LogFile  string
	Position int
}
