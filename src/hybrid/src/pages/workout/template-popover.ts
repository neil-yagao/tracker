import { Component } from '@angular/core';
import { URLSearchParams } from '@angular/http';

import { HttpBase } from '../../app/httpbase';


import { ViewController } from 'ionic-angular';

@Component({
    selector: 'template-popover',
    templateUrl: 'template-popover.html'
})
export class TemplatePopover {

    userTemplate: any = {};
    constructor(public http: HttpBase, public viewCtrl: ViewController) {
        http.get('templates/' + window.localStorage.getItem('username')).subscribe((userTemplate) => {
            this.userTemplate = userTemplate
        })
    }

    showModal() {
        this.viewCtrl.dismiss({
            action: "add"
        })

    }

    unassignTemplate(index: number) {
        var param: URLSearchParams = new URLSearchParams();
        param.set('user', window.localStorage.getItem('username'));
        param.set('templatesid', this.userTemplate.templatesId[index]);
        //console.info("deleting template:" + this.userTemplate.templates[index])
        this.http.delete('templates/user', param).subscribe(_ => {
            this.viewCtrl.dismiss({
                action: "delete"
            })
        })
    }
}
