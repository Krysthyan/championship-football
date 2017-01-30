package models

import (
	"encoding/json"
	"github.com/astaxie/beego/orm"
)

type Match struct {
	Id string `json:"id" orm:"pk"`
	Team_winner string `json:"team_winner" orm:"null;on_update(set_null)"`
	Team_losser string `json:"team_losser" orm:"null;on_update(set_null)"`
	Is_draw int `json:"is_draw"`
	Stage_id string `json:"stage_id"`
}

type Posicion struct {
	Stage_id string `json:"stage_id" orm:"pk"`
	Name string `json:"name"`
	Points string `json:"points"`
}

type Combinations struct {
	team1 Team
	team2 Team
}

func InsertMatch(match Match) []byte {
	return ORM_INSERT(match)
}

func CombinationsTeam(iterable []Team) (partidos []Combinations)  {

	for index:=len(iterable)-1;index > 0; index-- {
		for index1:=index-1;index1 >= 0; index1-- {

			partidos = append(partidos, Combinations{
				team1:iterable[index],
				team2:iterable[index1],
			})

		}
	}
	return
}

func GetTablePosition(id_stage string) (mapB []byte)  {

	var position []Posicion
	var listValues []orm.ParamsList

	qb, _ := orm.NewQueryBuilder("mysql")
	qb.Select("*").From("posicion").Where("Stage_id = ? ")
	ORM().Raw(qb.String(), id_stage).ValuesList(&listValues)

	for _, element := range listValues {
		position = append(position, Posicion{
			Stage_id:element[0].(string),
			Name: element[1].(string),
			Points:   element[3].(string),
		})
	}
	mapB, _ = json.Marshal(position)

	return
}
