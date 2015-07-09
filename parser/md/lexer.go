package md

import (
	"errors"
	"github.com/kamichidu/go-jclass/parser/fd"
)

func Parse(descriptor string) ([]string, string, int, error) {
	length := len(descriptor)
	if length < 3 {
		return []string{}, "", 0, errors.New("Given string is not a method descriptor")
	}

	p := 0
	if descriptor[p] != '(' {
		return []string{}, "", 0, errors.New("Method descriptor must start with `('.")
	}
	p++

	paramTypes := make([]string, 0)
	for descriptor[p] != ')' {
		paramType, n, err := fd.Parse(descriptor[p:])
		if err != nil {
			return paramTypes, "", p, err
		}
		paramTypes = append(paramTypes, paramType)
		p += n
	}

	if descriptor[p] != ')' {
		return paramTypes, "", p, errors.New("Method descriptor must have balanced parenthesis.")
	}
	p++

	if descriptor[p] == 'V' {
		p++
		return paramTypes, "void", p, nil
	} else {
		retType, n, err := fd.Parse(descriptor[p:])
		return paramTypes, retType, p + n, err
	}
}
