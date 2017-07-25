<template>
<div>
    <md-dialog ref="set" id="set-confirm-modal">
        <md-dialog-content>
            <md-layout md-gutter>
                <md-layout md-flex="44">
                    <md-input-container :md-clearable="true" md-numeric>
                        <label>使用的重量(KG)</label>
                        <md-input type="number" v-model="weight"></md-input>
                    </md-input-container>
                </md-layout>
                <md-layout md-flex="33" md-flex-offset="10">
                    <md-input-container md-numeric>
                        <label>完成个数</label>
                        <md-input type="number" v-model="reps"></md-input>
                    </md-input-container>
                </md-layout>
            </md-layout>
        </md-dialog-content>
        <md-dialog-actions>
            <md-button class="md-primary" v-on:click.native="closeDialog()">确定</md-button>
        </md-dialog-actions>
    </md-dialog>
</div>
</template>
<script>

export default {
	name:'complete-set',
	data(){
		return {
			weight:'',
			reps:this.defaultReps
		}
	},
	props:['defaultReps'],
	methods:{
		open(){
			this.$refs['set'].open()
		},
		closeDialog(){
			this.$refs['set'].close();
			this.$emit('closed',{
				weight:Number(this.weight),
				reps:Number(this.reps)
			})
		}
	}
}
</script>
<style>
#set-confirm-modal{
	height:50%;
}
</style>