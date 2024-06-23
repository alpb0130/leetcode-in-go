package serialize_dictionary

import (
	"strings"
)

type Item struct {
	IsString bool
	StrVal   string
	MapVal   map[string]Item
}

func Serialize(input map[string]Item) string {
	var res strings.Builder
	res.WriteString("{")
	strs := []string{}
	for key, item := range input {
		var subStr strings.Builder
		subStr.WriteString(key + ":")
		if item.IsString {
			subStr.WriteString(item.StrVal)
		} else {
			subStr.WriteString(Serialize(item.MapVal))
		}
		strs = append(strs, subStr.String())
	}
	res.WriteString(strings.Join(strs, ","))
	res.WriteString("}")
	return res.String()
}

func Deserialize(s string) map[string]Item {
	if len(s) <= 1 {
		return nil
	}
	res := map[string]Item{}
	itemsStr := s[1 : len(s)-1]
	prev := -1
	bracketCnt := 0
	for i := 0; i < len(itemsStr); i++ {
		if itemsStr[i] == '{' {
			bracketCnt++
		} else if itemsStr[i] == '}' {
			bracketCnt--
		} else if itemsStr[i] == ',' {
			if bracketCnt == 0 {
				key, value := processItemString(itemsStr[prev+1 : i])
				res[key] = value
				prev = i
			}
		}
	}
	key, value := processItemString(itemsStr[prev+1:])
	res[key] = value
	return res
}

func processItemString(str string) (string, Item) {
	var key, value string
	for j := 0; j < len(str); j++ {
		if str[j] == ':' {
			key = str[:j]
			value = str[j+1:]
			break
		}
	}
	if strings.Contains(value, "{") {
		return key, Item{false, "", Deserialize(value)}
	} else {
		return key, Item{true, value, nil}
	}
}
