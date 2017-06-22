<template>
<div>
	<plan-basic v-if="step == 1"></plan-basic>
	<plan-session v-if="step == 2"></plan-session>
    <div>
        <md-radio v-model="step" id="my-test1" name="my-test-group1" md-value="1">计划的基本信息</md-radio>
        <md-radio v-model="step" id="my-test2" name="my-test-group1" md-value="2">计划内的课程</md-radio>
		<md-button class="md-primary md-raised" v-if="step == 1" v-on:click.native="step = 2">下一步</md-button>
    </div>
    <md-button class="md-accent md-raised" v-on:click.native="createSession()">创建训练课</md-button>
    <pre>
    	{{$store.state.newSession}}
    </pre>
</div>	
</template>
<script>
import PlanBasic from './plan-basic.vue'
import PlanSession from './plan-session.vue'
export default{
	name:'create-plan',
	data(){
		return {
			step:1,
			buttonLabel:'下一步'
		}
	},
	methods:{
		createSession:function(){
			console.info(this.$store.state.newSession)
			
		}
	},
	watch:{
		'step':function(val){
			console.info("step:" + this.step)
			if(val == 2){
				this.buttonLabel = '返回'
			}else if (val == 1){
				this.buttonLabel = '下一步'
			}
		}
	},
	components:{
		'plan-basic': PlanBasic,
		PlanSession
	}
}
</script>
