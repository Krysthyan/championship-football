import {Component, ViewChild, ElementRef, OnChanges, SimpleChanges} from "@angular/core";
import {ActivatedRoute, Router, NavigationExtras} from "@angular/router";
import {ChampionshipService} from '../../providers/championship.service';
import {SpinnerService} from '../../providers/spinner.service'


@Component({
    moduleId: module.id,
    selector: 'championship-init',
    templateUrl: 'chamship.component.html',
    providers: [ChampionshipService, SpinnerService],
    styleUrls:['chamship.component.css']
})

export class ChampionshipInitComponent{


    name_championship : string;
    id_championship: any;
    items: any;
    championship: any;
    players: any;
    selectClickedRow: Function;
    id_team: any;
    items_player: any;

    constructor(private route: ActivatedRoute,
                private http:ChampionshipService,
                public spinner:SpinnerService,
                private router:Router)  {
        this.route.queryParams.subscribe(params => {
            this.name_championship = params["name_championship"];
            this.id_championship = this.getIdChampionship();
        });

        this.selectClickedRow = function (item: any) {
            this.id_team = item.id;
            this.ngOnChanges();
        }


    }
    ngOnChanges(): void {

        this.http.getListPlayerFormTeam(this.id_team, this.id_championship).subscribe(
            data => {
                this.items_player  = data.json();
            },
            err => {

            },
            () => {
                console.log("Obtener jugadores")
            }

        )
    }

    asigarEquiposAleatorios(num:number){

        this.http.asignarEquipoAleatorio(this.id_championship, num).subscribe(
            data =>{
                console.log(data.json());
                this.items = data.json();

            },
            err => {
                console.error(err);
            },
            () => {
                console.log("asignado correctamente los equipos");
            }
        );

        // this.loadOptionPlayer(this.player)

    }

    jugarCampeonato(){
        this.http.postDividirGrupos(this.items, this.id_championship).subscribe(
            data => {
                let navigationExtras: NavigationExtras = {
                    queryParams: {
                        "name_championship": this.name_championship
                    }
                };
                this.router.navigate(['championship/playRound'], navigationExtras);
            },
            err => {

            },
            () => {
                console.log("asignado grupos");
            }
        );


    }

    asignarJugadores(){
        this.spinner.start();
        console.log("entradmos espera un momento");
        this.http.postAsignarJugadores(this.items,11, this.id_championship).subscribe(
            data => {
                console.log(data.json());
                this.players = data.json();
            },
            err => {
                console.error(err);
            },
            () => {
                console.log("Realizado");
            }


        );
        this.spinner.stop();
    }

    getIdChampionship(){
        this.http.getChampionshipFormNamr(this.name_championship).subscribe(
            data =>{
                this.championship = data.json();
                console.log("estamos aqui");
                console.log(this.championship.id);
                this.id_championship = this.championship.id;
                console.log("id de campeonato");
                console.log(this.id_championship);
            },
            err => {
                console.error(err);
            },
            () => {
                console.log("get id of championship");
            }
        )
    }

}