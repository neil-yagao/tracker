import { Component, ViewChild, AfterViewInit, EventEmitter, Output} from '@angular/core';

import { WorkoutTemplate} from './workout-template'


@Component({
	selector: 'plan-basic',
	templateUrl: 'plan-basic.html'

})

export class PlanBasicPage implements AfterViewInit {
    @ViewChild('dateTime') dateTime;
    private template: WorkoutTemplate = new WorkoutTemplate();
	@Output() nextStep: EventEmitter<WorkoutTemplate> = new EventEmitter<WorkoutTemplate>();
    constructor() {
        console.info(this.template.startAt);
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
}
