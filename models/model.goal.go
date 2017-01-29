package models

import "log"

type Goal struct {
	Player_id string `json:"player_id" orm:"pk"`
	Match_id string `json:"match_id"`
	Number_goals int `json:"number_goals"`
}

func InsertGoal(goal Goal) []byte{
	return ORM_INSERT(goal)
}

func IsExistPlayerGoal(goal Goal) (isExist bool) {

	count,_ :=ORM().QueryTable("goal").Filter("player_id",goal.Player_id).Filter("match_id", goal.Match_id).Count()

	if int(count) > 0 {
		isExist = true
	}

	return
}

func UpdateGoal(goal Goal)  {
	if ORM().Read(&goal) == nil {
		goal.Number_goals += 1

		if _, err := ORM().Update(&goal); err == nil {
			log.Println(err)
		}
	}
}