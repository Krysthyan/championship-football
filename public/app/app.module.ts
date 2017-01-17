import {NgModule, Compiler}  from '@angular/core';
import {BrowserModule}       from '@angular/platform-browser';
import {FormsModule}         from '@angular/forms';
import {HttpModule}          from '@angular/http';
import {Ng2PaginationModule} from 'ng2-pagination'

import {AppComponent}         from './desktop/desktop.component';
import {TeamDesktopComponent}   from './team/team_desktop/team_desktop.component';
import {TeamTableComponent} from './team/team_table/team_table.component'
import {HeroService}          from './providers/team.service';
import {TeamComponent}        from './team/team_description/team_description.component'
import {JugadorComponet}      from './player/player_table/player_table.component'
import {AppRoutingModule}     from './app-routing.module';


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
        JugadorComponet
    ],

    providers: [HeroService],
    bootstrap: [AppComponent]
})

export class AppModule {
    constructor(private _compiler: Compiler) {
        _compiler.clearCache()
    }
}
