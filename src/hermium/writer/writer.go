package main

import "hermium/settings"
import "path/filepath"
import "os"

const CLIENT_SETTINGS_FILE = ".hermium_client_settings"
const COORD_SETTINGS_FILE = ".hermium_coordinator_settings"

	
func check(e error) {
    if e != nil {
        panic(e)
    }
}

func main() {
	clientSettingsFilePath := filepath.Join(os.Getenv("HOME"), CLIENT_SETTINGS_FILE)
	coordinatorSettingsFilePath := filepath.Join(os.Getenv("HOME"), COORD_SETTINGS_FILE)
	
	clientFile, err := os.Create(clientSettingsFilePath)
	check(err)
		
	clientSettings := &settings.ClientSettings {
		ListenPort: 1337,
		CoordinatorAddr: "localhost:8008",
	}

	clientSettings.Write(clientFile)
	defer clientFile.Close()

	coordinatorFile, err := os.Create(coordinatorSettingsFilePath)
	check(err)

	coordinatorSettings := &settings.CoordinatorSettings {
		ListenPort: 8008,
		Shards: 1000,
	}

	coordinatorSettings.Write(coordinatorFile)
	defer coordinatorFile.Close()

}
