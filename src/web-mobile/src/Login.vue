<template>
<div class="vcenter">
	<md-card>

	    <md-card-header>
	        <div class="md-title">欢迎加入撸铁大家庭，给自己起个名字吧！</div>
	    </md-card-header>

	    <md-card-content>
	        <md-input-container :md-clearable="true">
	            <md-input type="text" v-model="username"></md-input>
	        </md-input-container>
	    </md-card-content>
	    <md-card-actions>
	        <md-button v-on:click.native="login()">进入</md-button>
	    </md-card-actions>

	</md-card>
</div>

</template>
<script>
import md5 from 'md5'
export default {
    name: 'login',
    data() {
        return {
            username: ''
        }
    },
    methods: {
        login() {
            var md5Encode = md5(this.username);
            console.info(md5Encode);
            /*this.$http.post('login',{
            	'username': this.username,
            	'usr':md5Encode
            }).then((res)=>{
           		this.$store.state.username = this.username;
            	window.location.href = '#/working';
            })*/
            this.$store.commit('setUser',{
                username:this.username,
                userIdentity: md5Encode,
                id:1
            })
			this.$router.push('/working/workouts');

        }
    }
}
</script>
<style scoped>
.vcenter {
    margin-top: 50%;
}
</style>
