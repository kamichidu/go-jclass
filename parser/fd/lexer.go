package fd

import (
	"errors"
	"strings"
)

func Parse(descriptor string) (string, error) {
	length := len(descriptor)
	if length < 1 {
		return "", errors.New("Empty string is not a field descriptor")
	}

	if length == 1 {
		switch descriptor {
		case "B":
			return "byte", nil
		case "C":
			return "char", nil
		case "D":
			return "double", nil
		case "F":
			return "float", nil
		case "I":
			return "int", nil
		case "J":
			return "long", nil
		case "S":
			return "short", nil
		case "Z":
			return "boolean", nil
		default:
			return "", errors.New("Syntax error: " + descriptor)
		}
	}
	if descriptor[0] == 'L' {
		for i := 1; i < length; i++ {
			if descriptor[i] == ';' {
				if i != length-1 {
					return "", errors.New("Syntax error: " + descriptor)
				}
				break
			}
		}
		return strings.Replace(descriptor[1:length-1], "/", ".", -1), nil
	}
	if descriptor[0] == '[' {
		ret, err := Parse(descriptor[1:])
		return ret + "[]", err
	}

	return "", errors.New("Syntax error: " + descriptor)
}
