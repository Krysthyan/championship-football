import {Component, OnInit, Input} from '@angular/core';

@Component({
    moduleId: module.id,
    selector: 'team',
    templateUrl: 'team_description.component.html',
    styleUrls: ['team_description.component.css']
})

export class TeamComponent implements OnInit {

    @Input() name: string;

    constructor() {

    }

    ngOnInit(): void {

    }

}