<template>
<div>
    <md-dialog ref="movement-history" style="max-height:35em" @close="closing()">
        <md-dialog-title>动作历史</md-dialog-title>
        <md-dialog-content>
            <md-list md-double-line v-if="!loading">
                <md-list-item class="md-inset" v-for="sm in sessionMovements">
                    <div class="md-list-text-container">
                        <span>sm.BelongSession.executionDate</span>
                        <span>{{mergeExercise(sm)}}</span>
                    </div>
                </md-list-item>
            </md-list>
            <md-layout md-flex="35" md-align="center" v-if="loading">
				<md-spinner md-indeterminate></md-spinner>
			</md-layout>
        </md-dialog-content>
    </md-dialog>
</div>
</template>
<script>
export default {
	name:'movement-history',
	data(){
		return {
			loading:true,
			sessionMovements:[]
		}
	},
	props:['movement'],
	methods:{
		open(){
			this.$http.get('/session-movement/history/' + this.$store.state.user.id + "/" + this.movement)
			.then((res)=>{
				this.$refs['movement-history'].open();
				this.sessionMovements = res.body.data;
				this.loading = false;
			})
		},
		close(){
			this.$refs['movement-history'].close();
		},
		mergeExercise(sm){
			var history = ''
			for (var i in sm.exercises){
				var exercise = sm.exercises[i]
				history += exercise.weight + "KG X " +exercise.reps + ";"
			}
			return history;
		}
	}
}
</script>