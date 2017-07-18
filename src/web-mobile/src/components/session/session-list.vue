<template>
<div>
	<md-list class="md-double-line">
		<md-list-item v-for="session in userSessions" v-on:click.native="toDetail(session)"> 
		<md-icon>directions_run</md-icon>
	    	<div class='md-list-text-container'>
	  			<h2 class="list-head" >{{session.originSession.name}}</h2>
	  			<span class="list-content">{{translateDate(session.expectingDate)}}</span>
	  			<md-ink-ripple></md-ink-ripple>
	  		</div>
	  		<md-button class="md-icon-button md-list-action">
		        <md-icon class="md-primary">keyboard_arrow_right</md-icon>
		    </md-button>
		    <md-divider></md-divider>
		</md-list-item>
	</md-list>
	<loading-modal ref="loadingModal"></loading-modal>
</div>
</template>
<script>
import moment from 'moment'
import LoadingModal from '../general/loading.vue'

export default {
	name:'session-list',
	data(){
		return {
			userSessions:[]
		}
	},
	methods:{
		loadSessions(){
			this.$http.get('/session/' + this.$store.state.user.username).then((res)=>{
				console.info(res.body.data)
				this.userSessions = res.body.data;
			})
		},
		translateDate(date){
			return moment(date).format("MM-DD")
		},
		toDetail(session){
			this.$refs['loadingModal'].open();
			this.$http.get('/session-detail/' + session.id).then((res) =>{
				console.info(res.body.data)
				this.$store.commit('startExecutingSession', res.body.data);
				this.$refs['loadingModal'].close()
				this.$router.push("/working/workouts/detail/" + session.id)
			})
		}
	},
	mounted:function(){
		this.loadSessions()
	},
	components:{
		'loading-modal': LoadingModal

	}
}

</script>