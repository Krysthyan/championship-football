import {Component, Input} from '@angular/core';

@Component({
    moduleId: module.id,
    selector: 'person-edit',
    templateUrl: 'edit_person.component.html',
    styleUrls: ['edit_person.component.css']
})

export class EditPersonComponent  {


    @Input() name_team :string;
    constructor(){

    }


}