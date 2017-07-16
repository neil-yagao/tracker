<template>
<div>
    <md-list class="md-double-line">
        <md-list-item v-for='plan in exsitingPlan' v-on:click.native='checkPlan(plan)'>
            <md-icon>visibility</md-icon>
            <div class='md-list-text-container'>
                <span class="list-head">{{plan.name}}</span>
                <span class="list-content">{{plan.createBy}}</span>
                <md-ink-ripple></md-ink-ripple>
            </div>
            <md-divider></md-divider>
        </md-list-item>
        <md-button class='md-raised md-warn' v-on:click.native='openCreationPromot()'>未找到满意的？自己创建一个吧</md-button>
    </md-list>
    <movement-selection></movement-selection>
    <md-dialog-prompt md-title="创建新计划" md-ok-text="创建" md-cancel-text="取消" @close="onClose" v-model="$store.state.plan.planName" md-input-maxlength='20' md-input-placeholder='新计划名称' ref="createPlanPromot">
    </md-dialog-prompt>
    <loading-modal ref="loadingModal"></loading-modal>
</div>

</template>
<script>
import LoadingModal from '../general/loading.vue'
import Movement from '../movement/movement.vue'
export default {
	name:'plan-list',
	data(){
		return {
			exsitingPlan: [
				{
					name:'plan1',
					createBy: 'builder'
				},
				{
					name:'plan2',
					createBy: 'builder'
				}
			]
		}
	},
	methods:{
		checkPlan(plan){
			console.info(plan)
			this.$refs['loadingModal'].open();
			setTimeout(()=>{
				this.$refs['loadingModal'].close();
			}, 1000)
		},
		onClose(type){
			if(type == 'ok'){
				this.$router.push('/working/plan/sessions')
			}
		},
		openCreationPromot(){
			this.$refs['createPlanPromot'].open();
		}

	},
	components:{
		'loading-modal': LoadingModal,
		'movement-selection': Movement

	}
}
</script>
<style>
.list-head {
	font-size: 48px
}
.list-content {
	font-size: 30px
}
</style>