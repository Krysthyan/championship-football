package models

func IsChampionshipInTable(id string) (is_exist bool) {

	count, _ := ORM().QueryTable("stage_team").Filter("championship_id", id).Count()

	if int(count) > 0 {
		is_exist = true
	}
	return
}
