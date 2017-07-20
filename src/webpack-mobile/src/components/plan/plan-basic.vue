<template>
<div class="vertical-center">
    <md-list>
        <md-list-item>
            <md-input-container>
                <label>计划名称</label>
                <md-input v-model="$store.state.newSession.plan"></md-input>
            </md-input-container>
        </md-list-item>


        <md-list-item>
            <md-input-container>
                <md-textarea v-model="$store.state.newSession.planDescription" placeholder="简要的介绍一下你的计划吧！"></md-textarea>
            </md-input-container>
        </md-list-item>
 
        <md-list-item>
        	<md-button class="md-raised md-accent" v-on:click.native="startCreatePlan()">开始创建全新的计划吧</md-button>
        	</md-button>
        </md-list-item>
    </md-list>
</div>

</template>
<script>
export default {
	name:'session',
	data(){
		return {
		}
	},
	methods:{
		startCreatePlan(){
			window.location.href = '#/working/plan/sessions'
		}
	}
}
	
</script>
<style type="text/css">
.vertical-center {
	position: absolute;
	top:30%;
	width: 100%
}
</style>