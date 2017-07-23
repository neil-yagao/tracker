<template>
<div >
	<md-toolbar class="md-transparent">
        <md-button class="md-icon-button md-dense md-accent" v-on:click.native="openReturnConfirm()">
            <md-icon>reply</md-icon>
        </md-button>
        <span class="md-title" style="flex:1">{{title}}训练内容</span>
    </md-toolbar>
    <div style="overflow-y: auto; max-height: 25em;">
        <md-list md-triple-line>
            <md-list-item v-for="(row, rowIndex) in exercises">
                <div class="md-list-text-container">
                    <span>{{ row.name || row.movement.name}}</span>
                    <span>重复组数： {{ row.sets }}</span>
                    <p>每组数量 {{ row.reps }}个</p>
                </div>
                <md-button class="md-icon-button md-accent md-list-action" v-if="$store.state.plan.planEditing" v-on:click.native="removeExercise(row)">
                    <md-icon>delete</md-icon>
                </md-button>
                <md-divider></md-divider>
            </md-list-item>
        </md-list>
    </div>
    <md-button class="md-raised" v-on:click.native="showNewExerciseModal()" v-if="$store.state.plan.planEditing">添加新动作</md-button>
    <md-dialog ref="newExerciseModal" style="max-height:70%">
	    <md-list>
	        <movement id="movement-element" @selected-movement="selectMovement($event)"></movement>
	        <md-list-item>
	            <md-input-container :md-clearable="true" md-numeric >
	                <label>重复组数</label>
	                <md-input type="number" v-model="activeExercise.sets" required></md-input>
	            </md-input-container>
	        </md-list-item>
	        <md-list-item>
	            <md-input-container :md-clearable="true" >
	                <label>每组次数</label>
	                <md-input type="number" v-model="activeExercise.reps" required></md-input>
	            </md-input-container>
	        </md-list-item>
	         <md-dialog-actions>
	            <md-button class="md-primary" v-on:click.native="addExercise()">添加动作</md-button>
	            <md-button class="md-primary" v-on:click.native="closeModal()">取消</md-button>
	        </md-dialog-actions>
	    </md-list>
    </md-dialog>
</div>


	
</template>
<script>
import Movement from '../movement/movement.vue';
import _ from 'lodash';
	export default {
	    name: 'session-exercise',
	    data() {
	        return {
	            activeExercise: {}

	        }
	    },
	    methods: {
	        addExercise: function() {
                this.activeExercise.sequence = this.exercises.length + 1
                this.activeExercise.sets = Number(this.activeExercise.sets);
                this.activeExercise.reps = Number(this.activeExercise.reps);
                this.$store.commit('pushExercise',  _.clone(this.activeExercise))
                this.$http.post("/session/"+ this.session.id + "/movements",this.activeExercise).then(()=>{
                	this.activeExercise.sets = '';
	                this.activeExercise.reps = '';
	                this.activeExercise = {};
	                this.closeModal()
                })
                
	        },
	        removeExercise(row){
	        	this.$store.commit('removeExercise', row )
	        },
            doneEdit(){
                this.$router.replace("/working/plan/" +this.$route.params.session+ "/sessions")
            },
            selectMovement(movement){
            	this.activeExercise.movement = movement;
            },
	        openReturnConfirm(){
        		this.doneEdit();
	        },
	        showNewExerciseModal(){
	        	this.$refs['newExerciseModal'].open()
	        },
	        closeModal(){
	        	this.$refs['newExerciseModal'].close()
	        }
	    },
        computed:{
            sessionIndex: function(){
                return this.$route.params.id;
            },
            exercises: function(){
                return this.session.workouts;
            },
            title:function(){
                return this.session.name
            },
            session:function(){
            	return this.$store.state.plan.sessions[this.sessionIndex]
            }
        },
        components:{
        	Movement
        }
	}
</script>
<style>
#movement-element {
	padding-left: 16px;
	padding-right: 15px;
}
</style>