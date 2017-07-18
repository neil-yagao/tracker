<template>
<div>
    <md-list class="md-double-line">
        <md-list-item v-for='plan in exsitingPlan' v-on:click.native='checkPlan(plan)'>
            <md-icon>visibility</md-icon>
            <div class='md-list-text-container'>
                <span class="list-head">{{plan.name}}</span>
                <span class="list-content">{{plan.createby.username}}</span>
                <md-ink-ripple></md-ink-ripple>
            </div>
            <md-divider></md-divider>
        </md-list-item>
        <md-button class='md-raised md-warn' v-on:click.native='openCreationPromot()'>未找到满意的？自己创建一个吧</md-button>
    </md-list>
    <md-dialog-prompt md-title="创建新计划" md-ok-text="创建" md-cancel-text="取消" @close="onClose" v-model="$store.state.plan.planName" md-input-maxlength='20' md-input-placeholder='新计划名称' ref="createPlanPromot">
    </md-dialog-prompt>
    <loading-modal ref="loadingModal"></loading-modal>
</div>

</template>
<script>
import LoadingModal from '../general/loading.vue'
export default {
	name:'plan-list',
	data(){
		return {
			exsitingPlan: [	
			]
		}
	},
	methods:{
		checkPlan(plan){
			console.info(plan)
			this.$refs['loadingModal'].open();
			this.$http.get('/plan/' + plan.name).then((res) =>{
				var planDetail = res.body.data;
				console.info(planDetail);
				this.$store.commit('resetPlan', planDetail);
				this.$router.push('/working/plan/sessions');
				this.$refs['loadingModal'].close();
			})
	
		},
		onClose(type){
			if(type == 'ok'){
				this.$router.push('/working/plan/sessions');
				this.$store.commit('editingPlan')
			}
		},
		openCreationPromot(){
			this.$refs['createPlanPromot'].open();
		},
		loadPlans(){
			this.$http.get('/plans').then((res)=>{
				console.info(res)
				this.exsitingPlan = res.body.data;
			})
		}

	},
	components:{
		'loading-modal': LoadingModal

	},
	mounted:function(){
		this.loadPlans();
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