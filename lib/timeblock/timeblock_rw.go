// functions for reading/writing timeblock json file

package timeblock_lib

import (
	"encoding/json"
	"fmt"
	"os"
)

// save timeblock dict to file. give filepath with extension
func SaveTimeblockFile(timeblocksDict TimeblocksDict,filepath string) error {
    var jsonOut []byte
    var e error
    jsonOut,e=json.Marshal(timeblocksDict)

    if e!=nil {
        fmt.Println("error in converting timeblocks to json")
        return e
    }

    var file *os.File
    file,e=os.Create(filepath)

    if e!=nil {
        fmt.Println("failed to create file")
        return e
    }

    file.Write(jsonOut)
    file.Close()

    return nil
}

// try to load timeblock dict from target file
func LoadTimeblockFile(filepath string) (TimeblocksDict,error) {
    var rfile []byte
    var e error
    rfile,e=os.ReadFile(filepath)

    if e!=nil {
        fmt.Println("failed to read timeblock file")
        return make(TimeblocksDict),e
    }

    var dict TimeblocksDict=make(TimeblocksDict)
    e=json.Unmarshal(rfile,&dict)

    if e!=nil {
        fmt.Println("failed to parse timeblock file")
        return make(TimeblocksDict),e
    }

    return dict,nil
}