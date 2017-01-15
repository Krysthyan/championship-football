import {Component} from '@angular/core';


@Component({
    moduleId: module.id,
    selector: 'my-team',
    templateUrl: 'team_desktop.component.html'
})

export class TeamDesktopComponent {

    public id_team_select = 0;
    public name_team_select = '';

    constructor() {
    }

}