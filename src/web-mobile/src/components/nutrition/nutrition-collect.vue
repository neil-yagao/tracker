<template>
<div>
    <md-card md-with-hover>
        <md-card-header>
            <div class="md-title">我们需要一些数据来计算</div>
        </md-card-header>
        <md-card-content>
        	<md-list>
        		<md-list-item>
		        	<div class="md-list-text-container" :class="showError && gender == 0?'error':''"><label :reuqired="showError">您的性别是</label></div>
		            <md-button-toggle md-single class="md-list-action" >
		                <md-button class="md-primary" v-on:click.native="gender = 1" :class="gender == 1?'md-toggle':''">男性</md-button>
		                <md-button class="md-accent" v-on:click.native="gender = 2" :class="gender == 2?'md-toggle':''">女性</md-button>
		            </md-button-toggle>
	            </md-list-item>
	            <md-list-item>
	            	<md-layout>
		        	<md-layout md-flex="60"><label>您的年龄是</label></md-layout>
		        	<md-layout md-flex="20">
			            <md-input-container style="margin-top:-1.4em" :class="showError && !age?'md-input-invalid':''">
			            	<md-input v-model="age" md-numeric></md-input>
			            </md-input-container>
		            </md-layout>
		            </md-layout>
	            </md-list-item>
	            <md-list-item>
	            	<md-layout>
		        	<md-layout md-flex="60"><label>您的体重是多少(kg)</label></md-layout>
		        	<md-layout md-flex="20">
			            <md-input-container style="margin-top:-1.4em" :class="showError && !weight?'md-input-invalid':''">
			            	<md-input v-model="weight" md-numeric></md-input>
			            </md-input-container>
		            </md-layout>
		            </md-layout>
	            </md-list-item>
	            <md-list-item>
	            	<md-layout>
		        	<md-layout md-flex="60"><label>您的身高是多少(cm)</label></md-layout>
		        	<md-layout md-flex="20">
			            <md-input-container style="margin-top:-1.4em" :class="showError && !height?'md-input-invalid':''">
			            	<md-input v-model="height" md-numeric></md-input>
			            </md-input-container>
		            </md-layout>
		            </md-layout>
	            </md-list-item>
            </md-list>
        </md-card-content>
        <md-card-actions>
            <md-button v-on:click.native="doSave()">计算</md-button>
        </md-card-actions>
    </md-card>
</div>
</template>
<script>
export default {
	name:'nutrition-collect',
	data(){
		return {
			gender:0,
			weight:'',
			height:'',
			age:'',
			showError:false
		}
	},
	methods:{
		doSave(){
			if(this.gender == 0 || !this.weight || !this.height || !this.age){
				this.showError = true;
				return
			}
			var physique = {
				gender:Number(this.gender),
				weight:Number(this.weight),
				age:Number(this.age),
				height:Number(this.height)
			}
			this.$http.put("/physique/" + this.$store.state.user.id, physique).then(() =>{
				setTimeout(()=>{
					this.$router.replace('/working/nutrition')
				},500)
			})
		}
	}

}
</script>
<style scoped>
.error{
	color: #ff5722;
}
</style>