package repo

import (
	"fmt"
	"reflect"
	"strings"

	"google.golang.org/api/iterator"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
)

func Fields(in interface{}) string {
	var fields strings.Builder
	t := reflect.TypeOf(in)
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i).Tag.Get("db")
		table := t.Field(i).Tag.Get("table")
		if table == "" {
			continue
		}
		if len(field) > 0 {
			field = fmt.Sprintf("%s.`%s`", table, field)
			fields.WriteString(field)
			if i+1 < t.NumField() {
				fields.WriteByte(',')
			}
		}
	}

	return fields.String()
}

func CheckIteratorNextError(err error) error {
	if err != nil {
		if err == iterator.Done || grpc.Code(err) == codes.NotFound {
			return nil
		} else {
			return err
		}
	}

	return nil
}
