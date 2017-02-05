import { Component } from '@angular/core';
import { URLSearchParams } from '@angular/http';

import { HttpBase } from '../../app/httpbase';

import { WorkoutTemplate } from './workout-template';
import { ViewController,ModalController } from 'ionic-angular';

@Component({
	selector: 'template-popover',
	templateUrl: 'template-popover.html',
	providers:[HttpBase]
})
export class TemplatePopover {

	userTemplate:any = {};
	constructor(public http:HttpBase, public modalCtrl:ModalController, public viewCtrl:ViewController){
		http.get('templates/' + window.localStorage.getItem('username')).subscribe((userTemplate) => {
			this.userTemplate = userTemplate
		})
	}

	showModal(){
		this.viewCtrl.dismiss();
		let template = this.modalCtrl.create(WorkoutTemplate,{}, { showBackdrop: false, enableBackdropDismiss: false });
		template.onDidDismiss((selectedInfo) => {
			this.userTemplate.templates = selectedInfo.names;
			this.userTemplate.templatesId = selectedInfo.ids;
		})
		template.present();
	}

	unassignTemplate(index:number){
		var param:URLSearchParams = new URLSearchParams();
		param.set('userid', window.localStorage.getItem('username'));
		param.set('templatesid',this.userTemplate.templates[index]);
		//console.info("deleting template:" + this.userTemplate.templates[index])
		this.http.delete('templates/user', param).subscribe(_ => {})
	}
}
