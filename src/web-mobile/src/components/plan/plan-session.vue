<template>
<div style="margin-top:1em">
    <md-table-card>
        <md-toolbar>
            <h1 class="md-title">{{$store.state.newSession.sessisons[$route.params.id].name}}已添加的动作</h1>
        </md-toolbar>
        <md-table>
            <md-table-header>
                <md-table-row>
                    <md-table-head>#</md-table-head>
                    <md-table-head>动作名称</md-table-head>
                    <md-table-head md-numeric>重量</md-table-head>
                    <md-table-head md-numeric> 组数 </md-table-head>
                    <md-table-head md-numeric>每组数量</md-table-head>
                    <md-table-head></md-table-head>
                </md-table-row>
            </md-table-header>
             <md-table-body>
		      <md-table-row v-for="(row, rowIndex) in exercises" :key="rowIndex" :md-item="row">
		        <md-table-cell>
		          {{ rowIndex + 1}}
		        </md-table-cell>
		        <md-table-cell>
		          {{ row.name }}
		        </md-table-cell>
		        <md-table-cell md-numeric>
		          {{ row.weight }}
		        </md-table-cell>
		        <md-table-cell md-numeric>
		          {{ row.sets }}
		        </md-table-cell>
		        <md-table-cell md-numeric>
		          {{ row.repeats }}
		        </md-table-cell>
		        <md-table-cell>
			        <md-button style="width:100%" class="md-raised md-accent" v-on:click.native="removeExercise(row)">从课程中删除</md-button>
		        </md-table-cell>
		      </md-table-row>
		    </md-table-body>
        </md-table>
    </md-table-card>
    <hr>
    <md-list>
        <md-list-item>
            <md-input-container md-clearable="true">
                <label>动作名称</label>
                <md-input v-model="activeExercise.name"></md-input>
            </md-input-container>
        </md-list-item>
        <md-list-item>
            <md-input-container md-clearable="true">
                <label>重量</label>
                <md-input type="number" v-model="activeExercise.weight"></md-input>
            </md-input-container>
        </md-list-item>
        <md-list-item>
            <md-input-container md-clearable="true">
                <label>重复组数</label>
                <md-input type="number" v-model="activeExercise.sets"></md-input>
            </md-input-container>
        </md-list-item>
        <md-list-item>
            <md-input-container md-clearable="true">
                <label>每组次数</label>
                <md-input type="number" v-model="activeExercise.repeats"></md-input>
            </md-input-container>
        </md-list-item>
        <md-button class="md-raised md-warn" v-on:click.native="addExercise()">添加动作</md-button>
    </md-list>
</div>


	
</template>
<script>
import _ from 'lodash';
	export default {
	    name: 'session-exercise',
	    data() {
	        return {
	            activeExercise: {},
	            exercises:[],
	            selectedExercise:{}
	        }
	    },
	    methods: {
	        addExercise: function() {
	        	this.exercises.push(_.clone(this.activeExercise));
	        	this.activeExercise = {};
	        },
	        removeExercise(row){
	        	console.info(row)
        		this.exercises = _.without(this.exercises, row);
	        }
	    }
	}
</script>