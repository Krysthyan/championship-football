package tools

import (
	"crypto/rand"
	"fmt"
	"github.com/fatih/structs"
	"strconv"
	"strings"
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
		typeErr, _ := strconv.Atoi(
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
		*list = append(*list, ErrorChampionship{Error: err.Error()})
	}
}

func GetTypeError(err int, typeStruct string) string {
	switch err {
	case 1062:
		return "The " + typeStruct + " is already registered"
	}
	return string(err)

}

func GetPositionPlayer(valid_string string) (position string) {

	valid, _ := strconv.Atoi(valid_string)
	if valid >= 0 && valid <= 1 {
		position = "Arquero"
	} else if valid >= 2 && valid <= 4 {
		position = "Defensa"
	} else if valid >= 5 && valid <= 7 {
		position = "Mediocampista"
	} else if valid >= 8 && valid <= 10 {
		position = "Delantero"
	}
	return
}
