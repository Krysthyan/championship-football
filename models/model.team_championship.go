package models

type Team_championship struct {
	Championship_id string`orm:"pk;ref(fk)"`
	Team_id string `orm:"ref(fk)"`

}

func InsertTeamChampionship(team_championship Team_championship)  {
	ORM().Insert(&team_championship)
}

