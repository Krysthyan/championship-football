import {NgModule}             from '@angular/core';
import {RouterModule, Routes} from '@angular/router';

import {TeamDesktopComponent}   from './team/team_desktop/team_desktop.component';
import {ChampionshipComponent} from './championship/init/championship.component';
import {ChampionshipInitComponent} from './championship/Desktop/chamship.component';
import {RoundComponent} from './championship/Round/round.component'



const routes: Routes = [
    {path: '', redirectTo: '/desktop', pathMatch: 'full'},
    {path: 'desktop', component: ChampionshipComponent},
    {path: 'championship/team', component:TeamDesktopComponent},
    {path: 'championship/desktop',  component:ChampionshipInitComponent},
    {path: 'championship/playRound', component: RoundComponent}
];

@NgModule({
    imports: [RouterModule.forRoot(routes, {useHash: true})],
    exports: [RouterModule]
})

export class AppRoutingModule {
}

