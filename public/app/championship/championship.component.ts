import {Component, OnInit, OnChanges, SimpleChanges} from '@angular/core';
import {HeroService} from '../providers/team.service';
import {Router, NavigationExtras} from "@angular/router";


@Component({
    moduleId: module.id,
    selector: 'championship',
    templateUrl: 'championship.component.html',
    providers: [HeroService]
})
export class ChampionshipComponent  implements OnInit, OnChanges {
    items: any;
    public name_championship = '';

    constructor(
        private hero: HeroService,
        private router:Router) {

    }

    saveTeam(){

        this.hero.postSaveChampionship({name: this.name_championship}).subscribe(
            data => {
                let navigationExtras: NavigationExtras = {
                    queryParams: {
                        "name_championship": this.name_championship
                    }
                };
                this.router.navigate(['team'], navigationExtras);

            },
            err => {
                console.log(err)
            },
            () => {
                console.log("Championship save")
            }
        );
    }

    ngOnChanges(changes: SimpleChanges): void {

    }

    ngOnInit() {
        this.hero.getListChampionship().subscribe(
            data => {
                this.items = data.json();
                console.log(data.json());
            },
            err => {
                console.error(err)
            },
            () => {
                console.log("Get championship OK..")
            }
        );

    }
}
