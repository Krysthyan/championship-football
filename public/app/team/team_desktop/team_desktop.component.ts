import {Component} from '@angular/core';
import {HeroService} from '../../providers/team.service'
import {ActivatedRoute} from "@angular/router";

@Component({
    moduleId: module.id,
    selector: 'my-team',
    templateUrl: 'team_desktop.component.html',
    providers: [HeroService]

})


export class TeamDesktopComponent {

    public id_team_select: string;
    public name_team_select :string;

    public name_save_team = '';

    public teams_save_return : string;

    constructor(private hero: HeroService, private route: ActivatedRoute) {
        this.route.queryParams.subscribe(params => {
            this.name_save_team = params["name_championship"];
            console.log(this.name_save_team);
        });

    }

    saveTeam(){

        this.hero.postSaveTeam(this.name_save_team).subscribe(
            data => {

                this.teams_save_return = data.json()[0].name;

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

