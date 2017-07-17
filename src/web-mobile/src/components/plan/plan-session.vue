<template>
<div style="margin-top:1em">
<md-button class="md-fab md-raised md-accent" v-on:click.native="openReturnConfirm()">
	<md-icon>reply</md-icon>
</md-button>
    <md-table-card>
        <md-toolbar>
            <h1 class="md-title">{{title}}已添加的动作</h1>
        </md-toolbar>
        <md-table>
            <md-table-header>
                <md-table-row>
                    <md-table-head>#</md-table-head>
                    <md-table-head>动作名称</md-table-head>
                    <md-table-head> 组数 </md-table-head>
                    <md-table-head>每组数量</md-table-head>
                    <md-table-head  v-if="$store.state.plan.planEditing"></md-table-head>
                </md-table-row>
            </md-table-header>
             <md-table-body>
		      <md-table-row v-for="(row, rowIndex) in exercises" :key="rowIndex" :md-item="row">
		        <md-table-cell>
		          {{ rowIndex + 1}}
		        </md-table-cell>
		        <md-table-cell>
		          {{ row.name || row.movement.name}}
		        </md-table-cell>
		        <md-table-cell md-numeric>
		          {{ row.sets }}
		        </md-table-cell>
		        <md-table-cell md-numeric>
		          {{ row.reps }}
		        </md-table-cell>
		        <md-table-cell  v-if="$store.state.plan.planEditing">
                    <md-icon class="md-accent"  v-on:click.native="removeExercise(row)">delete</md-icon>
		        </md-table-cell>
		      </md-table-row>
		    </md-table-body>
        </md-table>
    </md-table-card>
    <hr>
    <md-list v-if="$store.state.plan.planEditing">
        <movement id="movement-element" @selected-movement="selectMovement($event)"></movement>
        <md-list-item>
            <md-input-container :md-clearable="true">
                <label>重复组数</label>
                <md-input type="number" v-model="activeExercise.sets"></md-input>
            </md-input-container>
        </md-list-item>
        <md-list-item>
            <md-input-container :md-clearable="true">
                <label>每组次数</label>
                <md-input type="number" v-model="activeExercise.reps"></md-input>
            </md-input-container>
        </md-list-item>
        <md-list-item>
            <md-button class="md-raised" v-on:click.native="addExercise()">添加动作</md-button>
            <md-button class="md-raised md-primary" v-on:click.native="doneEdit()">课程编辑完成</md-button>
        </md-list-item>
    </md-list>
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
                this.$store.commit('pushExercise',  _.clone(this.activeExercise))
                this.activeExercise.name = '';
                this.activeExercise.weight = '';
                this.activeExercise.sets = '';
                this.activeExercise.reps = '';
                this.activeExercise = {};
	        },
	        removeExercise(row){
	        	this.$store.commit('removeExercise', row )
	        },
            doneEdit(){
                this.$router.replace("/working/plan/sessions")
            },
            selectMovement(movement){
            	this.activeExercise.movement = movement;
            },
	        openReturnConfirm(){
	        	if(this.$store.state.plan.planEditing){
	            	this.$refs['confirmReturn'].open();
	        	}else{
	        		this.doneEdit();
	        	}
	        }
	    },
        computed:{
            sessionIndex: function(){
                return this.$route.params.id;
            },
            exercises: function(){
                return this.$store.state.plan.sessions[this.sessionIndex].workouts;
            },
            title:function(){
                return this.$store.state.plan.sessions[this.sessionIndex].name
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