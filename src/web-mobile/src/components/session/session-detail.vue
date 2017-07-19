<template>
<div>

	<div v-for="workout in $store.state.session.detail.workouts">
		<session-movement :movement="workout" @movement-done="setMovementDone($event)">
		</session-movement>
	</div>
	<md-dialog-alert
	  md-title="恭喜"
	  md-content="本次训练课所有动作均已完成！"
	  md-ok-text="确定"
	  ref="workout-finalize">
	</md-dialog-alert>
	</div>
</template>
<script>
import SessionMovement from './session-movement.vue'
export default {
		name:'session-detail',
		data(){
			return {}
		},
		methods:{
			setMovementDone(movement){
				var workouts = this.$store.state.session.detail.workouts;
				var acheived = 0;

				for(var i in workouts){
					if(workouts[i].status == 2){
						acheived += 1;
					}
				}
				if(acheived == workouts.length){
					//all done
					console.info('workout acheived');
					this.$router.replace("/working/workouts")
					this.$http.post('/session/' + this.$store.state.session.detail.id).then((res) =>{
						this.$refs['workout-finalize'].open();
					})
				}
				
			},
			openReturnConfirm(){
				this.$router.push("/working/workouts")
			}
		},
		components:{
			'session-movement':SessionMovement
		}
}
</script>