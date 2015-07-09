package fd

import (
	"errors"
	"strings"
)

func Parse(descriptor string) (string, int, error) {
	length := len(descriptor)
	if length < 1 {
		return "", 0, errors.New("Empty string is not a field descriptor")
	}

	switch descriptor[0] {
	case 'B':
		return "byte", 1, nil
	case 'C':
		return "char", 1, nil
	case 'D':
		return "double", 1, nil
	case 'F':
		return "float", 1, nil
	case 'I':
		return "int", 1, nil
	case 'J':
		return "long", 1, nil
	case 'S':
		return "short", 1, nil
	case 'Z':
		return "boolean", 1, nil
	case 'L':
		i := 1
		for ; i < length; i++ {
			if descriptor[i] == ';' {
				break
			}
		}
        className := strings.Replace(descriptor[1:i], "/", ".", -1)
		return className, i + 1, nil
	case '[':
		ret, n, err := Parse(descriptor[1:])
		return ret + "[]", n + 1, err
	default:
		return "", 0, errors.New("Syntax error: " + descriptor)
	}
}
