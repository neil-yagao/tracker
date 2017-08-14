<template>
<div>
    <md-toolbar class="md-dense">
        <md-button class="md-icon-button" v-if="$store.state.goback" v-on:click.native="goback()">
            <md-icon>keyboard_arrow_left
            </md-icon>
        </md-button>
        <span class="md-title" style="color:white; flex:1">举铁小助手</span>
        <span style="color:white">撸铁吧：{{username}}</span>
        <md-button class="md-icon-button" v-on:click.native="logout()">
            <md-icon>exit_to_app</md-icon>
        </md-button>
    </md-toolbar>
    <router-view></router-view>
    <md-bottom-bar class="bottom-stick">
        <md-bottom-bar-item md-icon="history" href='#/working/workouts' :md-active="$route.path.indexOf('workouts') > 0">训练课</md-bottom-bar-item>
        <md-bottom-bar-item md-icon="favorite" href="#/working/plan" :md-active="$route.path.indexOf('plan') > 0" :class="highlight" v-on:click="clearHighlight()"> 训练计划</md-bottom-bar-item>
        <md-bottom-bar-item md-icon="fitness_center" href="#/working/movement" :md-active="$route.path.indexOf('movement') > 0">训练动作</md-bottom-bar-item>
        <md-bottom-bar-item md-icon="near_me" href="#/working/nutrition" :md-active="$route.path.indexOf('nutrition') > 0">营养跟踪</md-bottom-bar-item>
    
    </md-bottom-bar>
    <md-dialog-alert :md-content="instruction" :md-ok-text="确定" @close="startHighlight" ref="showDialog" id="guide-promot">
    </md-dialog-alert>
</div>
</template>
<script>
export default {
    name: 'working-page',
    data() {
        return {
            highlight:'',
            shine:''
        }
    },
    computed: {
        username: function() {
            return this.$store.state.user.username;
        },
        isNew:function(){
            return this.$store.state.guide.new;
        },
        instruction:function(){
            return this.$store.state.guide.instruction;
        }
    },
    methods:{
        logout:function(){
            window.localStorage.setItem('user','')
            this.$router.replace('/')
        },
        goback(){
            this.$router.replace(this.$store.state.goback)
            this.$store.commit('setBackUrl','')
        },
        startHighlight(){
            if(this.$store.state.guide.new){
              this.shine = setInterval(()=>{
                console.info('interval')
                if(this.highlight){
                    this.highlight = ''
                }else {
                    this.highlight = 'highlight'
                }
            },1000)
            setTimeout(()=>{
                this.clearHighlight()
            },20000)  
            }
            
        },
        clearHighlight(){
            this.highlight = '';
            clearInterval(this.shine);
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
    },
    watch:{
        'instruction':function(){
            this.$refs['showDialog'].open()
        }
    }
}
</script>
<style type="text/css" global>
.md-tab-header-container{
	font-size:180%;
}
.bottom-stick{
	position: fixed;
    bottom: 0;
    width: 100%;
}
.impact {
	color:#ff5722;
}

.highlight {
  box-shadow: 3px 3px 15px #666;
  border-color:#C76C0C;
  background: #C76C0C;
  color: #fff;
  cursor: pointer;
  z-index:100;
  /*Opacity*/
  zoom: 1;
  filter: alpha(opacity=100);
  opacity: 1;
}
#guide-promot{
    height:30vh;
}
</style>
