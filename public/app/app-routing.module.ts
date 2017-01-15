import {NgModule}             from '@angular/core';
import {RouterModule, Routes} from '@angular/router';

import {TeamDesktopComponent}   from './team/team_desktop/team_desktop.component';
import {TeamComponent} from './team/team_description/team_description.component'

const routes: Routes = [
    {path: '', redirectTo: '/desktop', pathMatch: 'full'},
    {path: 'desktop', component: TeamDesktopComponent},
    {path: 'team', component: TeamComponent}
];

@NgModule({
    imports: [RouterModule.forRoot(routes, {useHash: true})],
    exports: [RouterModule]
})

export class AppRoutingModule {
}

