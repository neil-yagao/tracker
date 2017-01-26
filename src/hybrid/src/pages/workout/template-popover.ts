import { Component } from '@angular/core';
import { HttpBase } from '../../app/httpbase';

@Component({
	selector: 'template-popover',
	templateUrl: 'template-popover.html',
	providers:[HttpBase]
})
export class TemplatePopover {

	currentTemplate:string;
	constructor(public http:HttpBase){
		http.get('templates').subscribe((templates) => {
			//this.templates = templates
		})
	}
}