import { Component} from '@angular/core';
import { ViewController, AlertController } from 'ionic-angular';
import { HttpBase } from '../../app/httpbase';


import * as _ from "lodash";

@Component({
    selector:'workout-template',
    templateUrl:'./workout-template.html',
    providers:[HttpBase]
})
export class WorkoutTemplate {
    templates:Array<any>;
    timeStarts:string;
    selectedTemplate:Array<string>;
    selectedTemplateId:Array<any> = [];
    constructor(public viewCtrl:ViewController, public http:HttpBase,
        public alertCtrl:AlertController){
        http.get('templates').subscribe((templates)=>{
            this.templates = templates;
        })
    }

    selectFromExistingTemplate(){
		var options:any = {};
        this.selectedTemplate = [];
		options.title = "已有训练计划";
		options.inputs = [];
        let self = this;
		_.forEach(this.templates, function(temp){
			options.inputs.push({
				'type':'checkbox',
				'label':temp.name,
				'value':temp.id,
				'checked': _.findIndex(self.selectedTemplateId, function(id){
                    return id == temp.id
                }) >= 0
			})
		})

		options.buttons = [
		{
			'text':'取消',
			'role':'cancel'
		},
		{
			'text':'确定',
			handler: (data:Array<any>) => {
				self.selectedTemplateId = data
                _.find(self.templates, function(tpl){
                    var index = _.findIndex(self.selectedTemplateId,function(id){
                        return id == tpl.id
                    })
                    if(index != -1){
                        self.selectedTemplate = _.concat(self.selectedTemplate, tpl.name)
                    }
                })

			}
		}
		]
		let templateSelector = this.alertCtrl.create(options);
		templateSelector.present();
	}

    assignTemplate(){
        console.info(this.selectedTemplateId)
        var requestBody = {
            'userid':window.localStorage.getItem("username"),
            'templatesId':this.selectedTemplateId
        }
        this.http.post('templates/user',requestBody).subscribe(_ =>{
            this.viewCtrl.dismiss({
                ids: this.selectedTemplateId,
                names:this.selectedTemplate
            })
        })
    }
}
