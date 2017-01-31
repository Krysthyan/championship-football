import {Component, OnInit, Input, Output, EventEmitter, OnChanges, SimpleChanges} from '@angular/core';


import  {HeroService} from "../../providers/team.service"

@Component({
    moduleId: module.id,
    selector: 'team-table',
    templateUrl: 'team_table.component.html',
    providers: [HeroService]

})

export class TeamTableComponent implements OnInit, OnChanges {


    items: any;
    selectClickedRow: Function;
    @Input() team_save_return: Object;


    _id_select: string;
    @Output() id_nameChange = new EventEmitter();

    @Input()
    get id_name() {
        return this._id_select;
    }

    set id_name(id: string) {
        this._id_select = id;
        this.id_nameChange.emit(this._id_select);
    }


    _name_select: string;
    @Output() name_teamChange = new EventEmitter();

    @Input()
    get name_team() {
        return this._name_select;
    }

    set name_team(name: string) {
        this._name_select = name;
        this.name_teamChange.emit(this._name_select);
    }


    constructor(private hero: HeroService) {
        this.selectClickedRow = function (item: any) {
            this.id_name = item.id;
            this.name_team = item.name;
            console.log(this.id_name);
        }

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



    }


}