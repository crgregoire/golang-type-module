package scoping

import (
	"reflect"
	"regexp"
)

//
// FilterByScopes filters the fields on a struct
// by the provided scopes
//
func FilterByScopes(scopes []string, data interface{}) interface{} {
	t := reflect.TypeOf(data)
	out := reflect.New(reflect.Indirect(reflect.ValueOf(data)).Type()).Elem()
	dataValue := reflect.ValueOf(data)
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		value := dataValue.Field(i)
		if v, ok := field.Tag.Lookup("scope"); ok {
			if contains(scopes, v) {
				out.Field(i).Set(value)
			}
		}
	}
	return out.Interface()
}

//
// GetAllScopeTagsOnStruct returns a string slice of scope
// tags on a struct
//
func GetAllScopeTagsOnStruct(data interface{}) []string {
	t := reflect.TypeOf(data)
	tags := []string{}
	for i := 0; i < t.NumField(); i++ {
		tag := t.Field(i).Tag.Get("scope")
		if tag != "" && tag != "-" {
			tags = append(tags, tag)
		}
	}
	return tags
}

func contains(s []string, v string) bool {
	for _, str := range s {
		if checkPattern(str, v) {
			return true
		}
	}
	return false
}

func checkPattern(pattern, match string) bool {
	var validator *regexp.Regexp
	if pattern == "*" {
		pattern = "." + pattern
	}
	validator = regexp.MustCompile("^" + pattern)
	return validator.MatchString(match)
}
