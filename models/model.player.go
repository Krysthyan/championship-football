package models

import (
	"championship-football/tools"
	"encoding/json"
	"github.com/astaxie/beego/orm"
	"github.com/icrowley/fake"
	"strconv"
)

type Player struct {
	Id        string `json:"id" orm:"pk"`
	Name      string `json:"name"`
	Lastname  string `json:"lastname"`
	Direction string `json:"direction"`
	Position  string `json:"position"`
	Number    int    `json:"number"`
}

func InsertPlayer(player Player) []byte {
	return ORM_INSERT(player)
}

func GetPlayer(id string) (mapB []byte) {
	var player Player

	ORM().QueryTable("player").
		Filter("id", id).
		OrderBy("Id").
		One(&player)

	mapB, _ = json.Marshal(player)

	return
}

func GetPlayerList() (mapB []byte) {
	var players []Player

	ORM().QueryTable("player").
		OrderBy("Id").All(&players)

	mapB, _ = json.Marshal(players)
	return
}

func GetPlayerRamdon() Player {

	number, _ := strconv.Atoi(fake.DigitsN(2))
	return Player{
		Name:      fake.MaleFirstName(),
		Lastname:  fake.MaleLastName(),
		Direction: fake.StreetAddress(),
		Id:        fake.DigitsN(10),
		Position:  tools.GetPositionPlayer(fake.DigitsN(1)),
		Number:    number,
	}
}

func GetPlayerListRamdon(size int) (mapB []byte) {
	var listValues []orm.ParamsList

	var players []Player

	qb, _ := orm.NewQueryBuilder("mysql")
	qb.Select("*").From("player").OrderBy("RAND()").Limit(size)

	ORM().Raw(qb.String()).ValuesList(&listValues)

	for _, element := range listValues {
		number, _ := strconv.Atoi(element[4].(string))
		players = append(players, Player{
			Id:        element[0].(string),
			Name:      element[1].(string),
			Lastname:  element[2].(string),
			Direction: element[3].(string),
			Number:    number,
			Position:  element[5].(string),
		})
	}

	mapB, _ = json.Marshal(players)

	return
}
