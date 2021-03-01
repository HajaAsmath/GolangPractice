package customfilereader

import (
	"encoding/json"
	"fmt"
	"os"
)

func JsonConfigurator(file *os.File, container interface{}) error {
	fmt.Println("In jsonConfigurator")
	return json.NewDecoder(file).Decode(container)
}
