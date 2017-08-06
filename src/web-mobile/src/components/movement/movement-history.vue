<template>
<div>
    <md-dialog ref="movement-history" style="max-height:35em" @close="closing()">
        <md-dialog-title>动作历史</md-dialog-title>
        <md-dialog-content>
            
        </md-dialog-content>
    </md-dialog>
</div>
</template>
<script>
export default {
	name:'movement-history',
	data(){
		return {
			exercises:[]
		}
	},
	props:['movement'],
	methods:{
		open(){
			var user = this.$store.state.user.id
			this.$http.get('/session-movement/history/' + user + "/" + movement).then((res)=>{
				this.exercises = res.body.data.exercises;
			})
		}
	}
}
</script>