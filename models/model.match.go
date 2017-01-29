package models


type Match struct {
	Id string `json:"id" orm:"pk"`
	Team_winner string `json:"team_winner" orm:"null;on_update(set_null)"`
	Team_losser string `json:"team_losser" orm:"null;on_update(set_null)"`
	Stage_id string `json:"stage_id"`
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
