import {Injectable} from '@angular/core';
import {Http} from "@angular/http";

@Injectable()
export class HeroService {
    constructor(private http: Http) {

    }

    getTeam() {
        let json = this.http.get("http://localhost:8080/getTeam");
        return json;
    }

    getPersonTeam(id: number){
        let json = this.http.get("http://localhost:8080/getPersonTeam?id="+ id );
        return json;
    }


}
