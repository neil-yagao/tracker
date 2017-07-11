<template>
<div>
	<md-dialog ref="loading">
		<md-dialog-content>
			<md-layout md-flex="35" md-align="center">
				<md-spinner md-indeterminate></md-spinner>
			</md-layout>
		</md-dialog-content>
	</md-dialog>
</div>
</template>
<script>
export default {
	name:'loading-modal',
	methods:{
		open(){
			this.$refs['loading'].open();
		},
		close(){
			this.$refs['loading'].close();
		}
	}
}	
</script>