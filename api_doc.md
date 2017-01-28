
# Api Championship

## Team
- insert team
 
    > `http://localhost:8080/team/save`


- get all list teams
 
    > `http://localhost:8080/team/list`


- get teams from a championship
 
    > `http://localhost:8080/team/getFromChampionship`

        params
            - championship_id
            
- Agregar equipos automáticamente pasando un numero de equipos a guardar
 
    > `http://localhost:8080/team/addTeamsRamdon`

        params
            - num_team


## Player


- insert player

    >`http://localhost:8080/player/save`
    
            
- get player

    >`http://localhost:8080/player/get`

        params
            - player_id
            

- get player from a team and championship

    >`http://localhost:8080/player/getFromTeam`

        params
            - championship_id
            - player_id