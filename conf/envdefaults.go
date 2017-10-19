package conf

import (
	"fmt"
	"os"
	"reflect"
	"strconv"
	"time"

	"github.com/sirupsen/logrus"
)

const envtag = "env"
const defaulttag = "default"

func set(field reflect.Value, refType reflect.StructField, value string) error {
	switch field.Kind() {
	case reflect.String:
		field.SetString(value)
	case reflect.Bool:
		bvalue, err := strconv.ParseBool(value)
		if err != nil {
			return err
		}
		field.SetBool(bvalue)
	case reflect.Int:
		intValue, err := strconv.ParseInt(value, 10, 32)
		if err != nil {
			return err
		}
		field.SetInt(intValue)
	case reflect.Uint:
		uintValue, err := strconv.ParseUint(value, 10, 32)
		if err != nil {
			return err
		}
		field.SetUint(uintValue)
	case reflect.Float32:
		v, err := strconv.ParseFloat(value, 32)
		if err != nil {
			return err
		}
		field.SetFloat(v)
	case reflect.Float64:
		v, err := strconv.ParseFloat(value, 64)
		if err != nil {
			return err
		}
		field.Set(reflect.ValueOf(v))
	case reflect.Int64:
		if refType.Type.String() == "time.Duration" {
			dValue, err := time.ParseDuration(value)
			if err != nil {
				return err
			}
			field.Set(reflect.ValueOf(dValue))
		} else {
			intValue, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				return err
			}
			field.SetInt(intValue)
		}
	}
	return nil
}

// Parse parses the function
func Parse(in interface{}) error {
	ptrRef := reflect.ValueOf(in)
	ref := ptrRef.Elem()
	refType := ref.Type()
	for i := 0; i < refType.NumField(); i++ {
		f := ref.Field(i)
		ff := refType.Field(i)
		summary := fmt.Sprintf("%s.%s", refType.Name(), ff.Name)
		if p, ok := os.LookupEnv(ff.Tag.Get(envtag)); ok {
			if err := set(f, ff, p); err != nil {
				return err
			}
		}
		if d := ff.Tag.Get(defaulttag); d != "" {
			switch f.Kind() {
			case reflect.Int:
				if f.Int() == 0 {
					logrus.WithFields(logrus.Fields{"field": summary, "default": d}).Warn("No configured value, using default")
					if err := set(f, ff, d); err != nil {
						return err
					}
				}
			case reflect.Int64:
				if f.Int() == 0 {
					logrus.WithFields(logrus.Fields{"field": summary, "default": d}).Warn("No configured value, using default")
					if err := set(f, ff, d); err != nil {
						return err
					}
				}
			case reflect.Float64:
				if f.Float() == 0 {
					logrus.WithFields(logrus.Fields{"field": summary, "default": d}).Warn("No configured value, using default")
					if err := set(f, ff, d); err != nil {
						return err
					}
				}
			case reflect.String:
				if f.String() == "" {
					logrus.WithFields(logrus.Fields{"field": summary, "default": d}).Warn("No configured value, using default")
					if err := set(f, ff, d); err != nil {
						return err
					}
				}
			case reflect.Bool:
				if !f.Bool() {
					logrus.WithFields(logrus.Fields{"field": summary, "default": d}).Warn("No configured value, using default")
					if err := set(f, ff, d); err != nil {
						return err
					}
				}
			}
		}
	}

	return nil
}
