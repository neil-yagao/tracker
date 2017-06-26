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
        <!-- <md-list-item>
            <md-input-container>
                <label>目标肌群</label>
                <md-select name="muscle" id="muscle" v-model="$store.state.newSession.muscle">
                    <md-option value="肩部">肩部</md-option>
                    <md-option value="胸部">胸部</md-option>
                    <md-option value="背部">背部</md-option>
                    <md-option value="手臂">手臂</md-option>
                    <md-option value="核心">核心</md-option>
                    <md-option value="腿部">腿部</md-option>
                    <md-option value="体能">体能(HIIT)</md-option>
                </md-select>
            </md-input-container>
        </md-list-item>
        <md-list-item>
            
        </md-list-item> -->
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