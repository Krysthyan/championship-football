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

    getPersonTeam(id: string){
        let json = this.http.get("http://localhost:8080/getPersonTeam?id="+ id );
        return json;
    }

    postSaveTeam(team: any){
        let headers = new Headers({ 'Content-Type': 'application/json' });
        let options = new RequestOptions({ headers: headers });
        let body = JSON.stringify(team);
        let json = this.http.post('http://localhost:8080/team/save', body, headers);
        return json

    }


}
