package customfilereader

import (
	"bufio"
	"errors"
	"log"
	"os"
	"reflect"
	"strconv"
	"strings"
)

type fileContents map[string]reflect.Value

func (fileContents *fileContents) Add(name, value, fieldType string) error {
	fileCon := *fileContents
	switch fieldType {
	case "STRING":
		fileCon[name] = reflect.ValueOf(value)
	case "BOOL":
		v, err := strconv.ParseBool(value)
		if err != nil {
			return err
		}
		fileCon[name] = reflect.ValueOf(v)
	case "INTEGER":
		v, err := strconv.ParseInt(value, 32, 64)
		if err != nil {
			return err
		}
		fileCon[name] = reflect.ValueOf(v)
	case "FLOAT":
		v, err := strconv.ParseFloat(value, 64)
		if err != nil {
			return err
		}
		fileCon[name] = reflect.ValueOf(v)
	}
	return nil
}

func CostomConfigurator(file *os.File, container interface{}) error {
	defer file.Close()
	scanner := bufio.NewScanner(file)
	fileContents := make(fileContents)

	for scanner.Scan() {
		line := scanner.Text()
		line1 := strings.Split(line, "|")
		line2 := strings.Split(line1[1], ";")
		name := line1[0]
		value := line2[0]
		fieldType := line2[1]
		fieldType = strings.ToUpper(fieldType)
		err := fileContents.Add(name, value, fieldType)
		if err != nil {
			log.Fatal("Error parsing file", err)
			return err
		}
	}
	containerValue := reflect.ValueOf(container).Elem()
	containerType := containerValue.Type()
	for i := 0; i < containerValue.NumField(); i++ {
		conType := containerType.Field(i)
		conValue := containerValue.Field(i)
		tagName := conType.Tag.Get("name")
		if tagName == "" {
			log.Fatal("Tagname is empty")
			return errors.New("Tegname is Empty")
		}
		if fieldValue, ok := fileContents[tagName]; ok {
			conValue.Set(fieldValue)
		}
	}
	return nil
}
