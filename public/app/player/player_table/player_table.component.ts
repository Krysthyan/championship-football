import {Component, Input, OnChanges, SimpleChanges, OnInit} from '@angular/core';

import  {HeroService} from "../../providers/team.service"


@Component({
    moduleId:    module.id,
    selector:    'player',
    templateUrl: 'player_table.component.html',
    providers:   [HeroService]
})

export class JugadorComponet implements OnChanges, OnInit {


    items: any;
    @Input() id_team: string;
    @Input() name_team: string;
    public num_jugadores:number;
    selectClickedRow: Function;

    public name: string;
    public lastname: string;

    constructor(private hero: HeroService) {
        this.selectClickedRow = function (item: any) {
            this.name = item.name;
            this.lastname = item.lastname
        }
    }
    agregarJugadores(){
        this.hero.postSavePlayer(this.num_jugadores).subscribe(
            data => {
                this.ngOnInit();
            }
        );
    }
    ngOnChanges(changes: SimpleChanges): void {
        if (this.id_team.length > 0){
            console.log("Player");
            console.log(this.id_team);
            this.ngOnInit();
        }

    }
    ngOnInit() {
        this.hero.getPlayers().subscribe(
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



    }


}