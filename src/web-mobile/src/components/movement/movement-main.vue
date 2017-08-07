<template>
<div>
	 <md-layout>
        <md-layout md-flex="80" style="margin-left:2vw">
            <md-input-container>
                <md-icon>search</md-icon>
                <label>Search</label>
                <md-input type="text" v-model="name"></md-input>
            </md-input-container>
        </md-layout>
        <md-layout md-flex="10" md-flex-offset="3" style="margin-top:2vh">
            <md-menu id="muscles" md-size="4">
                <md-button class="md-icon-button" md-menu-trigger>
                    <md-icon>filter_list</md-icon>
                </md-button>
                <md-menu-content>
                	<md-menu-item v-on:click.native="muscle = ''">所有部位</md-menu-item>
				    <md-menu-item v-on:click.native="muscle = '胸部'">胸部动作</md-menu-item>
				    <md-menu-item v-on:click.native="muscle = '肩部'">肩部动作</md-menu-item>
                    <md-menu-item v-on:click.native="muscle = '背部'">背部动作</md-menu-item>
                    <md-menu-item v-on:click.native="muscle = '二头肌'">二头肌动作</md-menu-item>
                    <md-menu-item v-on:click.native="muscle = '三头肌'">三头肌动作</md-menu-item>
                    <md-menu-item v-on:click.native="muscle = '核心'">核心动作</md-menu-item>
                    <md-menu-item v-on:click.native="muscle = '腿部'">腿部动作</md-menu-item>
				  </md-menu-content>
            </md-menu>
        </md-layout>
    </md-layout>
    <div style="overflow-y: auto; max-height: 73vh; margin-top:-3vh">
    <md-list md-dense>
        <md-list-item v-for="movement in showList" v-on:click.native="checkMovement(movement)">
        <div class="md-list-text-container has-ripple">
           {{movement.name}}
           <md-ink-ripple/>
           </div>
           <md-button class="md-icon-button md-list-action">
	        <md-icon>keyboard_arrow_right</md-icon>
	      </md-button>
           <md-divider></md-divider>
        </md-list-item>
    </md-list>
    </div>
</div>
</template>
<script>
export default {
	name:'movement-panel',
	data(){
		return {
			muscle:'',
			movements:[],
			name:''
		}
	},
	methods:{
		loadMovements(){
			return this.$http.get('movements').then((res)=>{
		        		this.movements = res.body.data;
		        	})
		},
		checkMovement(movement){
			console.info(movement)
			this.$store.commit('setMovement', movement)
			this.$router.push('/working/movement/' + movement.id)
		}
	},
	mounted:function(){
		this.loadMovements()
	},
	computed:{
		showList:function(){
			var filterList = _.filter(this.movements, (m)=>{
				if(this.muscle && this.muscle != ''){
					return m.targetMuscle == this.muscle
				}else {
					return true
				}
			});
			filterList = _.filter(filterList, (m)=>{
				return m.name.indexOf(this.name) >= 0
			})

			return filterList
		}
	}
}
</script>