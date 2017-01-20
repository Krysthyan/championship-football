import {NgModule, Compiler}  from '@angular/core';
import {BrowserModule}       from '@angular/platform-browser';
import {FormsModule}         from '@angular/forms';
import {HttpModule}          from '@angular/http';
import {Ng2PaginationModule} from 'ng2-pagination'

import {AppComponent}         from './desktop/desktop.component';
import {TeamDesktopComponent}   from './team/team_desktop/team_desktop.component';
import {TeamTableComponent} from './team/team_table/team_table.component';
import {HeroService}          from './providers/team.service';
import {TeamComponent}        from './team/team_description/team_description.component';
import {JugadorComponet}      from './player/player_table/player_table.component';
import {AppRoutingModule}     from './app-routing.module';

import {SavePersonComponent} from "./person/save_person/save_person.component";
import {DeletePersonComponent} from './person/delete_person/delete_person.component';
import {EditPersonComponent} from './person/edit_person/edit_person.component';
import {ViewPersonComponent} from './person/view_person/view_person.component';


@NgModule({
    imports: [
        BrowserModule,
        HttpModule,
        FormsModule,
        AppRoutingModule,
        Ng2PaginationModule
    ],

    declarations: [
        AppComponent,
        TeamDesktopComponent,
        TeamTableComponent,
        TeamComponent,
        JugadorComponet,
        SavePersonComponent,
        DeletePersonComponent,
        EditPersonComponent,
        ViewPersonComponent
    ],

    providers: [HeroService],
    bootstrap: [AppComponent]
})

export class AppModule {
    constructor(private _compiler: Compiler) {
        _compiler.clearCache()
    }
}
