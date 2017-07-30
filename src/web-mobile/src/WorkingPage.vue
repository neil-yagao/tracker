<template>
<div>
    <md-toolbar class="md-dense">
        <md-button class="md-icon-button" v-if="$store.state.goback" v-on:click.native="goback()">
            <md-icon>keyboard_arrow_left
            </md-icon>
        </md-button>
        <span class="md-title" style="color:white; flex:1">举铁小助手</span>
        <span style="color:white">撸铁吧：{{username}}</span>
        <md-button class="md-icon-button" v-on:click.native="logout()"><md-icon >exit_to_app</md-icon></md-button>
    </md-toolbar>
    <router-view class="container-fluid" ></router-view>
    <md-bottom-bar class="bottom-stick" >
	  <md-bottom-bar-item md-icon="history" href='#/working/workouts' :md-active="$route.path.indexOf('workouts') > 0">训练课</md-bottom-bar-item>
	  <md-bottom-bar-item md-icon="favorite" href="#/working/plan" :md-active="$route.path.indexOf('plan') > 0" > 训练计划</md-bottom-bar-item>
	  <md-bottom-bar-item md-icon="near_me" href="#/working/nutrition" :md-active="$route.path.indexOf('nutrition') > 0">营养跟踪</md-bottom-bar-item>
</md-bottom-bar>
</div>
</template>
<script>
export default {
    name: 'working-page',
    data() {
        return {}
    },
    computed: {
        username: function() {
            return this.$store.state.user.username;
        }
    },
    methods:{
        logout:function(){
            window.localStorage.setItem('user','')
            this.$router.replace('/')
        },
        goback(){
            this.$router.replace(this.$store.state.goback)
        }
    },

    mounted:function(){
        //load user info from local storage
        if(!this.username){
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
            }else {
                this.$router.replace('/')
            }
        }
        
    }
}
</script>
<style type="text/css">
.md-tab-header-container{
	font-size:180%;
}
.bottom-stick{
	position: fixed;
    bottom: 0;
    width: 100%;
}
</style>
