<template>
    <div id="workout-create">
        <div class="row"><a class="btn btn-lick pull-right " href="#/lift/workouts">Back to List</a></div>
        <div class="row" style="margin-bottom:1em;text-align: center;">
            <div class="col-md-2" style="margin-top:6px">
                <label>训练课名称</label>
            </div>
            <div class="col-md-2">
                <input class="form-control" placeholder="Name" v-model="名称" type="text">
            </div>
            <div class="col-md-2" style="margin-top:6px">
                <label>训练目标肌群</label>
            </div>
            <div class="col-md-2">
                <input class="form-control" placeholder="Target Muscle Group" v-model="目标肌群" type="text">
            </div>
            <div class="col-md-2" style="margin-top:6px">
                <label>训练起始日</label>
            </div>
            <div class="col-md-2">
                <input class="form-control" placeholder="Start At" v-model="起始于" type="text">
            </div>
        </div>
        <div class="row" style="margin-bottom:1em;text-align: center;">
            <div class="col-md-2" style="margin-top:6px">
                <label>训练课注意事项</label>
            </div>
            <div class="col-md-5">
                <input class="form-control" placeholder="Notes during the workout..." v-model="注意事项" type="text">
            </div>
        </div>
        <div class="panel panel-default">
            <div class="panel-heading">
                训练课内容
            </div>
            <div class="panel-body">
                <div class="col-md-6 col-md-offset-3">
                    <table class="table table-condensed table-hover">
                        <thead>
                            <tr>
                                <td><b>#</b></td>
                                <td>动作</td>
                                <td>重量</td>
                                <td>每组次数</td>
                                <td>组数</td>
                            </tr>
                        </thead>
                        <tbody>
                            <tr v-for="(movement, index) in movements">
                                <td><b>{{index + 1}}</b></td>
                                <td>{{movement.name}}</td>
                                <td>{{movement.weight}}</td>
                                <td>{{movement.repeats}}</td>
                                <td>{{movement.sets}}</td>
                                <td>
                                    <button type="button" class="btn btn-danger btn-xs" @click="remove(index)">&times;</button>
                                </td>
                            </tr>
                        </tbody>
                    </table>
                </div>
                <div class="row">
                    <exercise-create @addMovement="updateMovement" :movement="{}"></exercise-create>
                </div>
            </div>
            <div class="panel-footer">
                <div class="row">
                    <div class="col-md-3">
                        <div class="input-group">
                            <div class="input-group-btn">
                                <button type="button" class="btn btn-default dropdown-toggle" data-toggle="dropdown" aria-haspopup="true" aria-expanded="false">每周<span class="caret"></span></button>
                                <ul class="dropdown-menu">
                                    <li v-for="day in week" @click="weekly = day"><a>{{day.show}}</a>
                                    </li>
                                </ul>
                            </div>
                            <input type="text" class="form-control" v-model="weekly.show">
                        </div>
                    </div>
                    <div class="col-md-2">
                        <h5>每次训练课增加</h5>
                    </div>
                    <div class="col-md-3">
                        <div class="input-group">
                            <span class="input-group-addon">增加重量</span>
                            <input type="text" class="form-control" v-model="addition">
                            <span class="input-group-addon">KG</span>
                        </div>
                    </div>
                    <div class="col-md-1 col-md-offset-2">
                        <button class="btn btn-default" @click="createWorkoutTemplate">创建
                        </button>
                    </div>
                </div>
            </div>
        </div>
    </div>
</template>
<script>
import ExerciseCreate from "./ExerciseCreate.vue"
export default {
    name: 'workout',
    data() {
        return {
            name: '',
            movements: [],
            startAt: new Date().toISOString().slice(0, 10),
            week: [{
                'show': '周一',
                'value': 'Monday'
            },{
                'show': '周二',
                'value': 'Tuesday'
            } ,{
                'show': '周三',
                'value': 'Wednesday'
            } ,{
                'show': '周四',
                'value': 'Thursday'
            } ,{
                'show': '周五',
                'value': 'Friday'
            } ,{
                'show': '周六',
                'value': 'Saturday'
            } ,{
                'show': '周日',
                'value': 'Sunday'
            } ],
            weekly: '',
            addition: '',
            target: '',
            description: ''
        }
    },
    components: {
        'exercise-create': ExerciseCreate
    },
    methods: {
        updateMovement: function(newMovement) {
            if (newMovement.name) {
                this.$data.movements.push(newMovement)
            }
        },
        remove: function(index) {
            this.$data.movements.splice(index, 1)
        },
        createWorkoutTemplate: function() {
            var requestBody = JSON.stringify({
                'name': this.name,
                'movements': this.movements,
                'startAt': this.startAt,
                'addition': this.addition,
                'weekly': this.weekly.value,
                'targetMuscle': this.target,
                'description': this.description
            });
            console.info(requestBody)
            var userIdentity = window.localStorage.getItem('user')
            console.info("user:" + userIdentity)
            this.$http.put('/workouts', requestBody, {
                'params': {
                    'user': userIdentity
                }
            })
        }
    }
}
</script>
<style scoped>
</style>
