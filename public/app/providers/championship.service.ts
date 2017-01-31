import {Injectable} from '@angular/core';
import {Http, Headers, RequestOptions} from "@angular/http";

@Injectable()
export class ChampionshipService {
    constructor(private http: Http) {

    }

    asignarEquipoAleatorio(name_championship:string, num_team:number){
        let json = this.http.get("http://localhost:8080/team/" +
            "asignarEquipos?num_teams="+num_team+"&championship_id="+ name_championship);
        return json;
    }

    getListPlayerFormTeam(id_team:string, championship_id: string){
        let json = this.http.get("http://localhost:8080/player/" +
            "getFromTeam?championship_id="+championship_id+"&team_id=" + id_team);
        return json;
    }

    getChampionshipFormNamr(name_championship:string){
        let json = this.http.get("http://localhost:8080/championship" +
            "/getFromName?name="+ name_championship);
        return json;
    }

    postAsignarJugadores(items: any, num_players: number, championship_id:string){
        let headers = new Headers({ 'Content-Type': 'application/json' });
        let options = new RequestOptions({ headers: headers });
        let body = JSON.stringify(items);
        let json = this.http.post('http://localhost:8080/player/' +
            'assingPlayerToTeam?' +
            'championship_id='+championship_id+'C&num_players='+num_players, body, headers);
        return json
    }
    postDividirGrupos(items: any, championship_id:string){
        let headers = new Headers({ 'Content-Type': 'application/json' });
        let options = new RequestOptions({ headers: headers });
        let body = JSON.stringify(items);
        let json = this.http.post('http://localhost:8080/stage/' +
            'generateStage?championship_id='+championship_id, body, headers);
        return json
    }
    getStageChampionship(id_championship:string){
        let json = this.http.get("http://localhost:8080/" +
            "stage/getListChampionship?championship_id="+id_championship);
        return json;
    }
    getPlayOff(id_championship: string){
        let json = this.http.get("http://localhost:8080/" +
            "stage/getPlayOffs?championship_id="+ id_championship);
        return json;
    }
    getTablePosicion(id_stage:string){
        let json = this.http.get("http://localhost:8080/" +
            "stage/getTablaPosicion?stage_id="+ id_stage);
        return json;
    }

    getCountTeamChampionship(championship_id: string){
        let json = this.http.get("http://localhost:8080/championship/" +
            "getCountTeam?championship_id=" + championship_id);
        return json;
    }

    playGroup(championSip_id: string){
        let json = this.http.get("http://localhost:8080/" +
            "stage/playGroup?championship_id="+ championSip_id);
        return json;
    }

    getTeamOff(stage_id : string){
        let json = this.http.get("http://localhost:8080/" +
            "stage/getTeamPlayOffs?stage_id="+ stage_id);
        return json
    }
}