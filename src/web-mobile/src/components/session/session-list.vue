<template>
<div style="overflow-y: auto; max-height: 84vh;">
	<md-button-toggle md-single>
	  <md-button class="md-toggle" style="flex:1" v-on:click.native="showState = 'assigned'">待完成</md-button>
	  <md-button class="md-primary" style="flex:1" v-on:click.native="showState = 'achieved'">已完成</md-button>
	</md-button-toggle>
	<md-list>
		<md-list-item v-for="session in userSessions" v-on:click.native="toDetail(session)"> 
		<md-icon>directions_run</md-icon>
	    	<div class='md-list-text-container'>
	  			<h2 class="list-head" >{{session.originSession.name}}</h2>
	  			<md-ink-ripple></md-ink-ripple>
	  		</div>
  			<span class="md-list-action md-accent">{{translateDate(session.expectingDate)}}</span>
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
			userSessions:[],
			showState:'assigned'
		}
	},
	methods:{
		loadSessions(state){
			this.$http.get('/session/' + this.$store.state.user.username + "/" +  state).then((res)=>{
				console.info(res.body.data)
				this.userSessions = _.sortBy(res.body.data,'expectingDate');
			})
		},
		translateDate(date){
			return moment(date).format("MM-DD")
		},
		toDetail(session){
			this.$http.get('/session-detail/' + session.id).then((res) =>{
				this.$store.commit('startExecutingSession', res.body.data);
				console.info(res.body.data)
				this.$router.push("/working/workouts/detail/" + session.id);
			})

		}
	},
	mounted:function(){
		this.loadSessions(this.showState);
		this.$store.commit('setBackUrl','');
	},
	watch:{
		showState:function(value){
			console.info('change to '+ value)
			this.loadSessions(value);
		}
	}
}

</script>