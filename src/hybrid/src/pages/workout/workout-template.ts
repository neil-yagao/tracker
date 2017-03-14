import { Component } from '@angular/core';
import { ViewController, AlertController, LoadingController } from 'ionic-angular';
import { HttpBase } from '../../app/httpbase';


import * as _ from "lodash";

@Component({
    selector: 'workout-template',
    templateUrl: './workout-template.html',
    providers: [HttpBase]
})
export class WorkoutTemplate {
    templates: Array < any > ;
    timeStarts: string = new Date().toISOString()
    selectedTemplate: Array < string > = [""];
    selectedTemplateId: Array < any > = [];
    constructor(public viewCtrl: ViewController, public http: HttpBase,
        public alertCtrl: AlertController, private loading: LoadingController) {
        http.get('templates').subscribe((templates) => {
            this.templates = templates;
        })
    }

    assignTemplate() {
        console.info(this.selectedTemplateId)
        var requestBody = {
            'userid': window.localStorage.getItem("username"),
            'templatesId': this.selectedTemplateId,
            'startDate': this.timeStarts.substring(0, 10)
        }
        let loader = this.loading.create({
            content: "正在为您应用计划，请稍后...",
            duration: 5000
        });
        loader.present()
        this.http.post('templates/user', requestBody).subscribe(_ => {
            this.viewCtrl.dismiss()
            loader.dismiss()
        })
    }

    dismiss() {
        this.viewCtrl.dismiss()
    }

    selectTemplate(value) {
        this.selectedTemplateId = value
    }
}
