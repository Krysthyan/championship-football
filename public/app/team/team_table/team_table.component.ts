import {Component, OnInit, Input, Output, EventEmitter} from '@angular/core';

import {Http} from "@angular/http";

import  {HeroService} from "../../providers/team.service"

@Component({
    moduleId: module.id,
    selector: 'team-table',
    templateUrl: 'team_table.component.html',
    providers: [HeroService]
})

export class TeamTableComponent implements OnInit {

    items: any;
    selectClickedRow: Function;

    _id_select: number;
    @Output() id_nameChange = new EventEmitter();

    @Input()
    get id_name() {
        return this._id_select;
    }

    set id_name(id: number) {
        this._id_select = id;
        this.id_nameChange.emit(this._id_select);
    }

    constructor(private hero: HeroService) {


    }

    ngOnInit() {
        this.hero.getTeam().subscribe(
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
            this.id_name = item.id;
        }

    }


}