import { Component, Input, EventEmitter, Output } from '@angular/core';
import { URLSearchParams } from '@angular/http';
import { LoadingController } from 'ionic-angular';

import { WorkoutTemplate } from './workout-template';
import { Exercise } from './exercise';

import { HttpBase } from '../../app/httpbase';

import * as _ from "lodash";
import { Md5 } from 'ts-md5/dist/md5';

@Component({
    selector: 'plan-exercise',
    templateUrl: 'plan-exercise.html',
    providers: [HttpBase]

})

export class PlanExercisePage {
    @Input() template: WorkoutTemplate;
    @Output() toMain: EventEmitter < boolean > = new EventEmitter < boolean > ();
    activeExercise: Exercise = new Exercise();
    private movements: Array < any >
        constructor(public http: HttpBase, private loading: LoadingController) {
            http.get('movements').subscribe((movements) => {
                this.movements = movements
            })
        }

    addExercise() {
        if (this.activeExercise.name) {
            var tempExercise = new Exercise();
            //directly clone will cause type incorrect.
            //force type change
            tempExercise.name = this.activeExercise.name;
            tempExercise.repeats = Number(this.activeExercise.repeats);
            tempExercise.weight = Number(this.activeExercise.weight);

            tempExercise.sets = Number(this.activeExercise.sets);
            this.template.addExercise(tempExercise);
            this.activeExercise = new Exercise();
        }
    }

    deleteExercise(exercise: Exercise) {
        _.remove(this.template.movements, function(e) {
            return e.name == exercise.name
        });

    }

    createSession() {
        let param = new URLSearchParams();
        param.set('user', window.localStorage.getItem('username'));
        let loader = this.loading.create({
            content: "正在为您创建训练计划，请耐心等待..."
        });
        this.http.put('workouts', this.template, param).subscribe((res) => {
            if (res.success) {
                loader.dismiss()
                this.toMain.emit(true);
            }
        });
    }
}
