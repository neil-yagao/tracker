<template>
	<p class="padding">
    <md-card md-with-hover >
        <md-card-header>
            <div class="md-title">{{movement.mappedMovement.movement.name}},需要完成{{movement.mappedMovement.sets}}组</div>
        </md-card-header>
        <md-card-expand ref="expand-toggle">
            <md-card-content>
                    <md-toolbar>
                        <h1 class="md-title">每组推荐完成数量为：{{movement.mappedMovement.reps}}</h1>
                    </md-toolbar>
                    <md-table md-sort="dessert" md-sort-type="desc" >
                        <md-table-header >
                            <md-table-row>
                                <md-table-head>完成重量(KG)</md-table-head>
                                <md-table-head>重复数量</md-table-head>
                                <md-table-head>删除该组</md-table-head>
                            </md-table-row>
                        </md-table-header>
                        <md-table-body>
                            <md-table-row v-for="exercise in exercises">
                                <md-table-cell>
                                    {{exercise.weight}}KG
                                </md-table-cell>
                                <md-table-cell>
                                    {{exercise.reps}}
                                </md-table-cell>
                                <md-table-cell>
			                        <md-icon  class="md-raised md-warn" :disabled="disabled" v-on:click.native="removeExercise(exercise)" >remove circle outline</md-icon>
                                </md-table-cell>
                            </md-table-row>
                        </md-table-body>
                    </md-table>
            </md-card-content>
            <md-card-actions>
                <md-button class="md-raised" v-on:click.native="executingMovement()">{{startTitle}}</md-button>
                <md-button class="md-raised md-primary" v-on:click.native="doneExecuting()">完成该动作训练</md-button>
                <span style="flex: 1"></span>
                <md-button class="md-icon-button" md-expand-trigger v-on:click.native="expandContent()">
                    <md-icon>keyboard_arrow_down</md-icon>
                </md-button>
            </md-card-actions>
        </md-card-expand>
    </md-card>
    <loading-modal ref="loadingModal"></loading-modal>
    <set-confirm ref="set-modal" @closed="pushExercise($event)" :defaultReps="movement.mappedMovement.reps"></set-confirm>
    <rest-modal ref="resting" rest="60"></rest-modal>
</p>
</template>
<script>
import SetConfirm from './set-confirm.vue'
import LoadingModal from '../general/loading.vue'
import RestModal from './rest-modal.vue'
export default {
	name:'session-movement',
	props:['movement'],
	data(){
		return {
			expanding:false,
			interval:'',
			executing:Number(this.movement.status)||0, //0 is not started, 1 is executing, 2 is done,
			exercises:this.movement.Exercises ||[]
		}
	},
	methods:{
		expandContent(){
			this.expanding = !this.expanding;
			console.info("expanding:" + this.expanding);
			this.flashContent();
			
		},
		executingMovement(){
			if(this.executing == 0){
				this.executing = 1;
				if(!this.expanding){
					this.$refs['expand-toggle'].toggle()
					this.expanding = true;
					this.showSetFinalizeModal()
				}
			}else if(this.executing == 2){
				this.expanding = true;
				this.$refs['expand-toggle'].toggle()
				this.flashContent()
				
			}else if(this.executing == 1){
				if(!this.expanding){
					this.expanding = true;
					this.$refs['expand-toggle'].toggle()
				}
				this.showSetFinalizeModal()
			}
			
		},
		doneExecuting(){
			this.executing = 2;
			if(this.expanding){
				this.$refs['expand-toggle'].toggle()
				this.expanding = false;
			}
			this.$http.post('/session-movement/' + this.movement.id, this.exercises).then((res) =>{
				console.info("finalize");
				this.movement.status = this.executing
				this.$emit('movement-done', this.movement)
			})
		},
		showSetFinalizeModal(){
			this.$refs['set-modal'].open();
		},
		pushExercise(exercise){
			this.exercises.push(exercise);
			this.$refs['resting'].open();
		},
		removeExercise(exercise){
			this.exercises = _.without(this.exercises, exercise)
		},
		flashContent(){
			if(this.interval){
				clearTimeout(this.interval)
			}
			if(this.executing != 1){
				this.interval = setTimeout(()=>{
					if(this.expanding){
						this.$refs['expand-toggle'].toggle()
						this.expanding = false;
					}
			}, 3000)
			}
		}
	},computed:{
		startTitle:function(){
			if(this.executing == 1){
				return '记录一组完成'
			}else if (this.executing == 0){
				return '开始该动作训练'
			}else if(this.executing == 2){
				return '查看已完成内容'
			}
		},
		disabled:function(){
			return this.executing == 2
		}
	},

	components:{
		'set-confirm':SetConfirm,
		'loading-modal': LoadingModal,
		'rest-modal': RestModal
	}
}
</script>
<style>
.padding{
	padding-top: 10px
}
th>div> .md-table .md-table-head-text {
	padding-right: 0px
}

</style>