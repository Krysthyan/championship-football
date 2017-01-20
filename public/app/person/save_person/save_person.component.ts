import {Component, Input} from '@angular/core';

@Component({
    moduleId: module.id,
    selector: 'person-save',
    templateUrl: 'save_person.component.html',
    styleUrls: ['save_person.component.css']
})

export class SavePersonComponent  {


    @Input() name_team :string;
    constructor(){

    }


}