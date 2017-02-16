import { Component, ViewChild, AfterViewInit } from '@angular/core';
import { URLSearchParams } from '@angular/http';
import { NavParams, Navbar, ModalController } from 'ionic-angular';

import * as _ from "lodash";
import { Md5 } from 'ts-md5/dist/md5';

import { HttpBase } from '../../app/httpbase';
import { RestingModal } from './resting-modal';

@Component({
    selector: 'workout-detail',
    templateUrl: 'workout-detail.html',
    providers: [HttpBase]
})
export class WorkoutDetail implements AfterViewInit {

    @ViewChild(Navbar) navBar: Navbar;
    private exercises: Array < any > ;
    private debounce: any;
    constructor(public params: NavParams, private http: HttpBase, public modalCtrl: ModalController) {
        var workoutId = params.data.workout.id;
        let param = new URLSearchParams();
        param.set('user', String(Md5.hashStr('powerlift')));
        http.get('session/' + workoutId, param).subscribe(session => {
            var movements = _.groupBy(session, 'movementName')
            var beforeAssign = _.map(movements, function(value, key) {
                var sequence = 0;
                var dividable = value[0]["dividable"]
                var newSets = _.map(value, function(set: any) {
                    var eachSide = (set.targetWeight - 20) / 2;
                    if (!set.dividable) {
                        eachSide = 1 / 0
                    }
                    if (set.sequence != 0) {
                        sequence += 1
                    }
                    return {
                        'id': set.id,
                        'sequence': set.sequence == 0 ? "*" : sequence,
                        'repeat': set.targetNumber,
                        'eachSide': eachSide,
                        'totalWeight': set.targetWeight,
                        "eachSideShow": false,
                        'achieved': -1
                    }
                })
                return {
                    "name": key,
                    "dividable": dividable,
                    "sets": newSets
                }
            })
            console.info(beforeAssign)
            this.exercises = beforeAssign;

        })

    }

    finishOneSet(set) {
        if (set.achieved == undefined || set.achieved <= 0) {
            set.achieved = set.repeat;
        } else {
            set.achieved -= 1
        }

        let self = this
        if (self.debounce) {
            self.debounce.cancel()
        }
        self.debounce = _.debounce(() => {
            let resting = this.modalCtrl.create(RestingModal, { timeout: 60 }, { showBackdrop: false, enableBackdropDismiss: false });
            resting.present();
        }, 1000)
        self.debounce()
    }

    finishSession() {
        var sessionData = _.reduce(this.exercises, function(arry, value) {
            return _.concat(arry, _.flatMap(value.sets, (set) => {
                return {
                    'id': set.id,
                    'acheiveNumber': +set.achieved,
                    'acheiveWeight': +set.totalWeight
                }
            }))

        }, [])
        console.info(sessionData)
        let param = new URLSearchParams();
        param.set('user', window.localStorage.getItem('username'));
        this.http.post('session/' + this.params.data.workout.id, sessionData, param)
            .subscribe(() => {})
    }

    resetTotalWeight(set, $event) {
        set.totalWeight = $event
        set.eachSide = ($event - 20) / 2
    }


    ngAfterViewInit() {
        this.navBar.setBackButtonText("返回列表");
    }
}
