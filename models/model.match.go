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
	Pj string `json:"pj"`
	Pg string `json:"pg"`
	Pp string `json:"pp"`
	Pe string `json:"pe"`
	Gf string `json:"gf"`
	Gc string `json:"gc"`
}

type Combinations struct {
	team1 Team `json:"team_1"`
	team2 Team `json:"team_2"`
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
			Pj: element[3].(string),
			Points: element[4].(string),
			Pg:element[5].(string),
			Pp:element[6].(string),
			Pe:element[7].(string),
			Gf:element[8].(string),
			Gc:element[9].(string),

		})
	}
	mapB, _ = json.Marshal(position)

	return
}

func GetMatchStage(id_stage string) (mapB []byte) {
	var match []Match
	ORM().QueryTable("match").Filter("Stage_id", id_stage).All(&match)

	mapB,_ = json.Marshal(match)
	return
}
