import {Component, Input} from '@angular/core';

@Component({
    moduleId: module.id,
    selector: 'person-delete',
    templateUrl: 'delete_person.component.html',
    styleUrls: ['delete_person.component.css']
})

export class DeletePersonComponent  {


    @Input() name_team :string;
    constructor(){

    }


}