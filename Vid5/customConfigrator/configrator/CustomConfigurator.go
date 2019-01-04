package configrator

import (
	"bufio"
	"errors"
	"os"
	"reflect"
	"strconv"
	"strings"
	"fmt"
)

type ConfigFields map[string]reflect.Value

func (f ConfigFields) Add(name, value, t string) error {
	switch t {
	case "STRING":
		f[name] = reflect.ValueOf(value)
	case "INTEGER":
		i, err := strconv.Atoi(value)
		if err != nil {
			return err
		}
		f[name] = reflect.ValueOf(i)
	case "FLOAT":
		fl, err := strconv.ParseFloat(value, 64)
		if err != nil {
			return err
		}
		f[name] = reflect.ValueOf(fl)
	case "BOOL":
		b, err := strconv.ParseBool(value)
		if err != nil {
			return err
		}
		f[name] = reflect.ValueOf(b)
	}
	return nil
}

func MarshalCustomConfig(v reflect.Value, filename string) error {

	defer func() {
		if v:=recover(); v!=nil {
			fmt.Print("Panic occured !")
		}
	}()                                                  //To avoid exiting program in case of saome panic

	if !v.CanSet() {
		return errors.New("Value passed not settable")
	}

	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()
	fields := make(ConfigFields)
	scanner := bufio.NewScanner(file)               //to read the file line by line , newScanner takes a io.Reader interface
	for scanner.Scan() {
		//each itertion scans individual lines of the file
		line := scanner.Text()
		fmt.Println("Processing line", line)
		if strings.Count(line, "|") != 1 || strings.Count(line, ";") != 1 {
			continue
		}
		args := strings.Split(line, "|")
		name := args[0]
		valuetype := strings.Split(args[1], ";")
		name, value, vtype := strings.TrimSpace(name), strings.TrimSpace(valuetype[0]), strings.ToUpper(strings.TrimSpace(valuetype[1]))
		fields.Add(name, value, vtype)
	}

	if err := scanner.Err(); err != nil {               //check if any error occured during the scanning process
		return err
	}

	vt := v.Type()                                       // -> struct
	for i := 0; i < v.NumField(); i++ {                  // -> num of field in the struct
		fieldType := vt.Field(i)
		fieldValue := v.Field(i)

		name := fieldType.Tag.Get("name")
		if name == "" {
			name = fieldType.Name
		}

		if v, ok := fields[name]; ok {
			fieldValue.Set(v)
		}
	}

	return nil
}
