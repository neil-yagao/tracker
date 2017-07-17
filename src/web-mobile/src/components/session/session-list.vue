<template>
<div>
	<md-list>
		<md-list-item v-for="session in userSessions">
	    	<div class='md-list-text-container'>
	  			<span class="list-head" >{{session.originSession.name}}</span>
	  			<span class="list-content">{{session.expectingDate}}</span>
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
			userSessions:[{
				id:1,
				assignTo:{
					id:1,
					username:'tester',
					useridentity:'awer4'
				},
				expectingDate:moment().format("MM-DD"),
				originSession:{
					id:1,
					name:"session1"
				},
				status:'assigned',
				workouts:[
					{
						id:1,
						belong:''

					}
				]
			}]
		}
	}
}

</script>