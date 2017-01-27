package tools

import (
	"strings"
	"strconv"
	"github.com/fatih/structs"
	"fmt"
	"crypto/rand"
)

type ErrorChampionship struct {
	Error string `json:"error"`
}

func ChampionshipToken(size int) string {
	b := make([]byte, size)
	rand.Read(b)
	return strings.ToUpper(fmt.Sprintf("%x", b))
}

func ChampionshipError(err error, structCham interface{}) (num_err string) {

	if err != nil {
		typeErr,_:= strconv.Atoi(
			strings.Split(
				strings.Split(
					err.Error(),
					" ")[1],
				":")[0])
	num_err = GetTypeError(typeErr, structs.Name(structCham))
	}
	return
}

func ListError(list *[]ErrorChampionship, err error) {
	if err != nil {
		*list = append(*list, ErrorChampionship{Error:err.Error()})
	}
}

func GetTypeError(err int, typeStruct string) string {
	switch err {
	case 1062:
		return "The " + typeStruct + " is already registered"
	}
	return string(err)

}
