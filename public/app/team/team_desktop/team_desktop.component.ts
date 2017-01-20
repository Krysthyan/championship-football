import {Component} from '@angular/core';
import {HeroService} from '../../providers/team.service'

@Component({
    moduleId: module.id,
    selector: 'my-team',
    templateUrl: 'team_desktop.component.html',
    providers: [HeroService]

})


export class TeamDesktopComponent {

    public id_team_select = 0;
    public name_team_select = '';

    public name_save_team = '';

    public teams_save_return : Team;

    constructor(private hero: HeroService) {
    }

    saveTeam(){

        this.hero.postSaveTeam({name: this.name_save_team}).subscribe(
            data => {

                this.teams_save_return = new Team(data.json().id,
                                                    data.json().name);
            },
            err => {
                console.log(err)
            },
            () => {
                console.log("team save")
            }
        );
    }

}

class Team{
    constructor(public id: string,
                public name: string){}
}