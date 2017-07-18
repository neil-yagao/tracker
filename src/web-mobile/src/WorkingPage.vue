<template>
<div>
    <md-toolbar>
        <md-button>
            <md-icon class="md-size-2x">menu</md-icon>
        </md-button>
        <span class="md-display-1" style="color:white">举铁小助手</span>
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
import Plan from './components/plan/plan.vue'
import {MdToolbar, MdButton, MdIcon} from 'vue-material';
export default {
    name: 'working-page',
    data() {
        return {}
    },
    computed: {
        'username': function() {
            return this.$store.state.username;
        }
    },
    components:{
    	Plan
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
