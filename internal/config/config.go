package config

type Config struct {
	DataFile        string
	DefaultPriority string
	StoragePath     string
}

var AppConfig = Config{
	DataFile:        "tasks.json",
	DefaultPriority: "low",
	StoragePath:     ".",
}
