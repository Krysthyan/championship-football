import {Component} from '@angular/core';

@Component({
    moduleId:    module.id,
    selector:    'desktop_app',
    styleUrls:   ['desktop.component.css'],
    templateUrl: 'desktop.component.html',
})
export class AppComponent {
    isActive = false;
    showMenu: string = '';

    eventCalled() {
        this.isActive = !this.isActive;
    }

    addExpandClass(element: any) {
        if (element === this.showMenu) {
            this.showMenu = '0';
        } else {
            this.showMenu = element;
        }
    }
}
