import { Component, ViewChild, AfterViewInit, EventEmitter, Output} from '@angular/core';
import { AlertController } from 'ionic-angular';
import { WorkoutTemplate} from './workout-template';
import { HttpBase } from '../../app/httpbase';


import * as _ from "lodash";

@Component({
	selector: 'plan-basic',
	templateUrl: 'plan-basic.html',
	providers:[HttpBase]

})

export class PlanBasicPage implements AfterViewInit {
    @ViewChild('dateTime') dateTime;
    private template: WorkoutTemplate = new WorkoutTemplate();
    private muscleGroup:any;
    private templates:Array<any>;
	@Output() nextStep: EventEmitter<WorkoutTemplate> = new EventEmitter<WorkoutTemplate>();
    constructor(public http:HttpBase, public alertCtrl: AlertController) {
        console.info(this.template.startAt);
        this.http.get('templates').subscribe((templates) => {
        	this.templates = templates
        })
    }

    ngAfterViewInit() {
        setTimeout(_ => {
			this.dateTime.setValue(new Date().toISOString());
		});
    }

	generateSession() {
		console.info(this.template);
		this.nextStep.emit(this.template)
	}

	combinedMuscle(targetArray: Array<string>) {
		this.template.targetMuscle = ""
		_.forEach(targetArray, (muscle) => {
			this.template.targetMuscle += muscle + ","
		});
		this.template.targetMuscle = _.trimEnd(this.template.targetMuscle, ",");
		this.muscleGroup = targetArray;
	}

	selectFromExistingTemplate(){
		var options:any = {};
		options.title = "已有训练计划";
		options.inputs = [];
		var currentTemplate = window.localStorage.getItem('template')
		_.forEach(this.templates, function(temp){
			options.inputs.push({
				'type':'radio',
				'label':temp.name,
				'value':temp.name,
				'checked': temp.name == currentTemplate
			})
		})
		let self = this;
		options.buttons = [
		{
			'text':'取消',
			'role':'cancel'
		},
		{
			'text':'确定',
			handler: (data) => {
				self.template.template = data
			}
		}
		]
		let templateSelector = this.alertCtrl.create(options);
		templateSelector.present();
	}
}
