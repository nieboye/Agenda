package config

import (
	"convention/codec"
	"log"
	"os"
	"time"
)

var debugMode = true

func DebugMode() bool { return debugMode }

// type Config = map[string](interface{})

// Config holds all configure of Agenda system.
var Config = make(map[string](interface{}))

func Load(decoder codec.Decoder) {
	cfg := &(Config)
	// CHECK: Need check if have already exactly loaded ALL config (i.e. eof) ?
	if err := decoder.Decode(cfg); err != nil {
		log.Fatal(err)
	}
}

func Save(encoder codec.Encoder) error {
	return encoder.Encode(Config)
}

// ... paths

// WorkingDir for agenda.
func WorkingDir() string {
	location, existed := os.LookupEnv("HOME")
	if !existed || DebugMode() {
		location = "."
	}
	ret := location + "/.agenda.d/"
	return ret
}

func init() {
	files := make(map[string](interface{}))
	files["all-user-registered-data"] = "user-registered.json"
	files["all-meeting-data"] = "meeting-data.json"
	files["user-logined-data"] = "curUser.txt"
	// "config.json"

	Config["files"] = files

}

var neededFilepaths = []string{
	UserDataRegisteredPath(),
	MeetingDataPath(),
	AgendaConfigPath(),
	UserLoginStatusPath(),
}

func NeededFilepaths() []string {
	return neededFilepaths
}

func UserDataRegisteredPath() string { return WorkingDir() + "user-registered.json" }
func MeetingDataPath() string        { return WorkingDir() + "meeting-data.json" }

func AgendaConfigPath() string { return WorkingDir() + "config.json" }

// func LogPath() string             { return WorkingDir() + "agenda_" + time.Now().Format("20060102_0304") + ".log" }
func LogPath() string             { return WorkingDir() + "agenda_" + time.Now().Format("20060102_15") + ".log" }
func UserLoginStatusPath() string { return WorkingDir() + "curUser.txt" }

func BackupDir() string {
	return WorkingDir() + "backup/"
}

var (
	// files     = Config["flies"].(map[string](interface{}))
	filepaths = NeededFilepaths()
)

func ensurePathsNeededExist() {
	if err := os.MkdirAll(WorkingDir(), 0777); err != nil {
		log.Fatal(err)
	}

	for _, path := range filepaths {
		if _, err := os.Stat(path); os.IsNotExist(err) {
			f, err := os.Create(path)
			defer f.Close()
			if err != nil {
				log.Fatal(err)
			}
		}
	}
}

func init() {
	ensurePathsNeededExist()
}
