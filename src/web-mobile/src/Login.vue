<template>
<div class="vcenter">
	<md-card class="md-primary">

	    <md-card-header>
	        <div class="md-title">给自己个响亮的外号，开撸！</div>
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

            var str = window.localStorage.getItem('user');
            var localStore = null;
            try{
                localStore = JSON.parse(str);
            }catch (e){
                console.info(str + " translate to json fail")
                console.info(e)
            }
            if(localStore && localStore.id){
                this.$store.commit('setUser', localStore);
                window.localStorage.setItem('user', localStore)
                this.$router.push('/working/workouts');
            }else {
                if(!this.username){
                    return
                }
                var md5Encode = md5(this.username);
                this.$http.post('login',{
                'username': this.username,
                'usr':md5Encode
                }).then((res)=>{
                    var user = res.body.data
                    console.info(JSON.stringify(user))
                    this.$store.commit('setUser',user);
                    window.localStorage.setItem('user', JSON.stringify(user))
                    this.$router.push('/working/workouts');
                })
            }
        }
    },
    mounted:function(){
        this.login()
    }
}
</script>
<style scoped>
.vcenter {
    margin-top: 50%;
}
</style>
