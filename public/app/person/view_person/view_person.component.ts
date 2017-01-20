import {Component, Input} from '@angular/core';

@Component({
    moduleId: module.id,
    selector: 'person-view',
    templateUrl: 'view_person.component.html',
    styleUrls: ['view_person.component.css']
})

export class ViewPersonComponent  {


    @Input() name_team :string;
    constructor(){

    }


}