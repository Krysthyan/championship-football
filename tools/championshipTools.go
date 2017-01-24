package tools

import (
	"github.com/nkanish2002/token_generator"
	"strings"
	"strconv"
	"github.com/fatih/structs"
)

func ChampionshipToken(size int) string {
	g := token_generator.Generator{}
	g.New()
	return strings.ToUpper(g.GetToken(size))
}

func ChampionshipError(err error, structCham interface{}) string{

	if err != nil {
		typeErr,_:= strconv.Atoi(
			strings.Split(
				strings.Split(
					err.Error(),
					" ")[1],
				":")[0])
		return GetTypeError(typeErr, structs.Name(structCham))
	}else {
		return ""
	}


}

func GetTypeError(err int, typeStruct string) string {
	switch err {
	case 1062:
		return "The " + typeStruct + " is already registered"
	}
	return string(err)

}
