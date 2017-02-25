import { Component } from '@angular/core';

import { NavController /*, PopoverController */ } from 'ionic-angular';

import { HttpBase } from '../../../app/httpbase';
import { PlanPage } from '../create/plan';
import { TemplateDetail } from './template-detail';



@Component({
    selector: 'template-list',
    templateUrl: 'plan-list.html'
})
export class PlanList {
    templates: Array < any > ;
    constructor(public navCtrl: NavController,
        public http: HttpBase /*, public popCtrl: PopoverController*/ ) {
        this.freshTemplateList()
    }

    freshTemplateList() {
        this.http.get('templates').subscribe((templates) => {
            this.templates = templates
        })
    }

    goToDetail(template) {
        this.navCtrl.push(TemplateDetail, { 'template': template })
    }

    openPopover(template) {
        this.navCtrl.push(PlanPage, { 'template': template });

    }

    ionViewWillEnter() {
        this.freshTemplateList()
    }
}
