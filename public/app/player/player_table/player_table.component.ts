import {Component, Input, OnChanges, SimpleChanges} from '@angular/core';

import  {HeroService} from "../../providers/team.service"


@Component({
    moduleId:    module.id,
    selector:    'player',
    templateUrl: 'player_table.component.html',
    providers:   [HeroService]
})

export class JugadorComponet implements OnChanges {


    items: any;
    @Input() id_team: number;

    selectClickedRow: Function;

    public name: string;
    public lastname: string;

    constructor(private hero: HeroService) {

    }

    ngOnChanges(changes: SimpleChanges): void {
        console.log("hoola mundo");
        this.hero.getPersonTeam(this.id_team).subscribe(
            data => {
                this.items = data.json();
            },
            err => {
                console.error(err)
            },
            () => {
                console.log("Datos obtenidos")
            }
        );
        this.selectClickedRow = function (item: any) {
            this.name = item.name;
            this.lastname = item.lastname
        }
    }


}