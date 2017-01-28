package models

type Team_stage struct {
	Team_id string `json:"team_id" orm:"pk"`
	Stage_id string `json:"stage_id"`
	Championship_id string `json:"championship_id"`
}

func InsertTeamStage(teamStage Team_stage) []byte {
	return ORM_INSERT(teamStage)
}


func IsChampionshipInTable(id string) (is_exist bool) {

	count, _ := ORM().QueryTable("stage_team").Filter("championship_id", id).Count()

	if int(count) > 0 {
		is_exist = true
	}
	return
}
