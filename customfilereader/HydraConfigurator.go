package customfilereader

import (
	"errors"
	"log"
	"os"
	"reflect"
)

const (
	CUSTOM string = "CUSTOM"
	JSON   string = "JSON"
)

func ReadFile(filename string, container interface{}, filetype string) error {
	defer func() {
		recover()
	}()
	containerValue := reflect.ValueOf(container)

	if containerValue.Kind() != reflect.Ptr {
		log.Fatal("The container has to be a pointer")
		return errors.New("The container is not a pointer")
	}

	containerType := containerValue.Type().Elem()

	if containerType.Kind() != reflect.Struct {
		log.Fatal("The container is not struct")
		return errors.New("The continer is not a struct")
	}

	file, err := os.Open(filename)

	if err != nil {
		log.Fatal("Error while opening file with filename " + filename)
		return err
	}
	switch filetype {
	case CUSTOM:
		err = CostomConfigurator(file, container)
	case JSON:
		err = JsonConfigurator(file, container)
	}
	return err
}
