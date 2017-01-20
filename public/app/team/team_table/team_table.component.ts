import {Component, OnInit, Input, Output, EventEmitter, OnChanges, SimpleChanges} from '@angular/core';

import {TeamDesktopComponent} from '../team_desktop/team_desktop.component'

import  {HeroService} from "../../providers/team.service"

@Component({
    moduleId: module.id,
    selector: 'team-table',
    templateUrl: 'team_table.component.html',
    providers: [HeroService]

})

export class TeamTableComponent implements OnInit, OnChanges {


    public items: Array<{id: string, name: string}>;
    selectClickedRow: Function;
    @Input() team_save_return: Object;


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
    //todo arreglar despues ya que esta haciendo de nuevo consulta en vez de solo de agregar el item
    ngOnChanges(changes: SimpleChanges): void {
       this.ngOnInit()

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