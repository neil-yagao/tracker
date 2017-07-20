<template>    
	<div>
        <md-dialog-title>
        <md-toolbar class="md-transparent">
        	<md-input-container style="flex: 1" >
                <md-icon>search</md-icon>
                <md-input class="md-input" v-model="criteria" id="criteria"></md-input>
            </md-input-container>
	        <md-button  class="md-raised md-primary" v-on:click.native="moveToAdd()">
	  			新增
	      	</md-button>
        </md-toolbar>
        </md-dialog-title>
        <md-dialog-content>
        	<div  style="max-height: 20em;overflow-y: auto;">
	            <md-list ref="m-list">
	                <md-list-item v-for="(movements, muscle) in movementList">
	                    <span>{{muscle}}</span>
	                    <md-list-expand>
	                        <md-list>
	                            <md-list-item class="md-inset" v-for="m in movements" v-on:dblclick.native="selectMovement(m)" v-on:click.native="selectingMovement(m)">{{m.name}}</md-list-item>
	                        </md-list>
	                    </md-list-expand>
	                </md-list-item>
	            </md-list>
            </div>
        </md-dialog-content>
        <md-dialog-actions>
            <md-button class="md-primary" v-on:click.native="closeDialog()">取消</md-button>
            <md-button class="md-primary" v-on:click.native="selectMovement(selectedMovement)">确定</md-button>
        </md-dialog-actions>
        </div>
</template>
<script>
import Vue from 'vue'
export default {
	name:'movment-add',
	data(){
		return {
			criteria:'',
			selectedMovement:{}
		}
	},
	methods:{
		moveToAdd(){
			this.$emit('to-add')
		},
		closeDialog(){
			this.$emit('close')
		},
		selectMovement(m){
			this.$emit('selected', m);
			this.selectedMovement = {};
		},
		selectingMovement(m){
			this.selectedMovement = m;
		}

	},
	props:['movements'],
	computed:{
		movementList:function(){
			var filtered = this.movements;
			if(this.criteria){
				filtered = _.filter(this.movements, (m) =>{
				return m.name.indexOf(this.criteria) >= 0
				})
				Vue.nextTick(() =>{
					for(var i in this.$refs['m-list'].$children){
						this.$refs['m-list'].$children[i].toggleExpandList()
					}
				})
			 
			}
			
			return _.groupBy(filtered,'targetMuscle')
		}
	}
}
</script>