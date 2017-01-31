import {Injectable} from '@angular/core';
import {Http, Headers, RequestOptions} from "@angular/http";

@Injectable()
export class HeroService {
    constructor(private http: Http) {

    }

    getTeam() {
        let json = this.http.get("http://localhost:8080/team/list");
        return json;
    }

    getPlayers(){
        let json = this.http.get("http://localhost:8080/player/getList");
        return json;
    }

    postSaveTeam(num_teams: any){
        let json = this.http.get("http://localhost:8080/team/addTeamsRamdon?num_team="+ num_teams);
        return json

    }
    postSavePlayer(num_players: number){
        let json = this.http.get("http://localhost:8080/player/addPlayersRandom?num_players="+ num_players);
        return json
    }



    postSaveChampionship(championship: any){
        let headers = new Headers({ 'Content-Type': 'application/json' });
        let options = new RequestOptions({ headers: headers });
        let body = JSON.stringify(championship);
        let json = this.http.post('http://localhost:8080/championship/save', body, headers);
        return json
    }

    getListChampionship(){
        let json = this.http.get("http://localhost:8080/championship/getList" );
        return json;
    }


}
