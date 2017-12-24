package settings

import "encoding/json"
import "fmt"
import "io/ioutil"
import "os"

type ClientSettings struct {
    ListenPort      uint32
    CoordinatorAddr string
}

type CoordinatorSettings struct {
    ListenPort      uint32
    Shards          uint32
}

func ReadClientSettings(path string) ClientSettings {
    raw, err := ioutil.ReadFile(path)
    if err != nil {
        fmt.Println(err.Error())
        os.Exit(1)
    }

    var s ClientSettings
    err = json.Unmarshal(raw, &s)
    if err != nil {
        fmt.Println(err.Error())
        os.Exit(1)
    }
    return s
}

func ReadCoordinatorSettings(path string) CoordinatorSettings {
    raw, err := ioutil.ReadFile(path)
    if err != nil {
        fmt.Println(err.Error())
        os.Exit(1)
    }

    var s CoordinatorSettings
    err = json.Unmarshal(raw, &s)
    if err != nil {
        fmt.Println(err.Error())
        os.Exit(1)
    }
    return s
}
