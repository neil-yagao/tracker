<template>
<div>

	<div v-for="workout in $store.state.session.detail.workouts">
		<session-movement :movement="workout" @movement-done="setMovementDone($event)">
		</session-movement>
	</div>

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
					this.$http.post('/session/' + this.$store.state.session.detail.id)
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