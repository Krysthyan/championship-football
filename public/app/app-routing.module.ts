import {NgModule}             from '@angular/core';
import {RouterModule, Routes} from '@angular/router';

import {TeamDesktopComponent}   from './team/team_desktop/team_desktop.component';


const routes: Routes = [
    {path: '', redirectTo: '/desktop', pathMatch: 'full'},
    {path: 'desktop', component: TeamDesktopComponent}

];

@NgModule({
    imports: [RouterModule.forRoot(routes, {useHash: true})],
    exports: [RouterModule]
})

export class AppRoutingModule {
}

