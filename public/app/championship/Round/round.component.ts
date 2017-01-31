import {Component, OnInit} from "@angular/core";
import {ActivatedRoute} from "@angular/router";
import {ChampionshipService} from '../../providers/championship.service';


@Component({
    moduleId: module.id,
    selector: 'round',
    templateUrl: 'round.component.html',
    providers: [ChampionshipService],
    styleUrls: ['round.component.css']
})

export class RoundComponent{


    name_championship:string;
    public id_championship: any;
    championship: any;
    public numTeams:number;
    public nummero:number;
    items: any;
    selectClickedRow: Function;
    selectPlayOff: Function;
    id_stage: any;
    items_equipos: any;



    item_stage_off : any;
    item_team_off: any;
    public id_team_off: string;

    constructor(private route: ActivatedRoute,
                private http:ChampionshipService){
        this.route.queryParams.subscribe(params => {
            this.name_championship = params["name_championship"];
            this.id_championship = this.getIdChampionship();
        });

        this.selectClickedRow = function (item: any) {
            this.id_stage = item.id;
            this.ngOnChanges();
        };

        this.selectPlayOff = function(item: any){
            this.id_team_off = item.id;
            this.getTeamOff();
        };

    }
    ngOnChanges(){
        this.http.getTablePosicion(this.id_stage).subscribe(
            data => {
                this.items_equipos = data.json();
            },
            err => {

            },
            () => {
                console.log('obtenida tabla de posicion');
            }
        );
    }


    trackByFn(index:number, item:any) {
    let playOff = ['ganador','final','Semifinal',"Cuartos de final","Octavos de final","Dieciseisavos de final"];
        let num:number = Math.log(this.numTeams)/Math.log(2);
        item.championship_id = playOff[3 - index ];
        return item;
    }

    getStage(id :string){
        this.http.getStageChampionship(id).subscribe(
            data => {
                this.items = data.json();
            },
            err => {

            },
            () => {
                console.log("obtenidos los grupos")
            }
        );
    }

    getPlayOff(){
        this.http.getPlayOff(this.id_championship).subscribe(
            data => {
                this.item_stage_off = data.json();
            }
        );
    }
    getTeamOff(){
        this.http.getTeamOff(this.id_team_off).subscribe(
            data => {
                this.item_team_off = data.json();
            }
        );
    }


    playGroup(){
        this.http.playGroup(this.id_championship).subscribe(
            data => {
                console.log(data.json());
                this.getPlayOff();
            }
        );
    }

    public getIdChampionship(){
        this.http.getChampionshipFormNamr(this.name_championship).subscribe(
            data =>{
                this.championship = data.json();
                this.id_championship = this.championship.id;
                this.getStage(this.id_championship);
                this.getPlayOff();


                this.http.getCountTeamChampionship(this.id_championship).subscribe(
                    data => {
                        this.numTeams = data.json();
                        this.nummero = this.numTeams;
                        console.log("Numero de equipos de este campeonato");
                        console.log(this.numTeams)
                    },
                    err => {
                        console.error(err)
                    },

                    () => {
                        console.log("Peticion de numero de equipos realizad");
                    }
                );
            },
            err => {
                console.error(err);
            },
            () => {
                console.log("get id of championship");
            }
        );

    }


}