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
</div>
</template>
<script>
import moment from 'moment'
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
			this.$router.push("/working/workouts/detail/" + session.id)
		}
	},
	mounted:function(){
		this.loadSessions()
	}
}

</script>