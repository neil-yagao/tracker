<template>
<div>
    <md-toolbar class="md-dense">
        <span class="md-title" style="color:white; flex:1">举铁小助手</span>
        <span style="color:white">撸铁吧：{{username}}</span>
    </md-toolbar>
    <router-view class="container-fluid"></router-view>
    <md-bottom-bar class="bottom-stick" >
	  <md-bottom-bar-item md-icon="history" href='#/working/workouts' :md-active="$route.path.indexOf('workouts') > 0">训练课</md-bottom-bar-item>
	  <md-bottom-bar-item md-icon="favorite" href="#/working/plan" :md-active="$route.path.indexOf('plan') > 0" > 训练计划</md-bottom-bar-item>
	  <md-bottom-bar-item md-icon="near_me">营养跟踪</md-bottom-bar-item>
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
