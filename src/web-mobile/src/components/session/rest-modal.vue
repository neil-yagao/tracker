<template>
<div>
    <md-dialog ref="resting-modal" style="max-height:35em" @close="closing()">
        <md-dialog-title>组间休息</md-dialog-title>
        <md-dialog-content>
            <md-spinner :md-size="250" :md-progress="progress" :class="color"></md-spinner>
            <span class="count" :class="left" >{{this.count}}</span>
        </md-dialog-content>
    </md-dialog>
</div>
</template>
<script>
export default {
	name:'resting',
	props:['rest'],
	data(){
		return {
			progress:0,
			count:0,
			interval:''
		}
	},
	methods:{
		open(){
			this.$refs['resting-modal'].open();
			this.interval = setInterval(()=>{
				if(this.progress > 100){
					this.progress = 0
				}
				this.progress += (10/6)
				this.count += 1
			}, 1000)
		},
		closing(){
			clearInterval(this.interval)
			this.progress = 0;
			this.count = 0
		}
	},
	computed:{
		color:function(){
			if(this.count < 60){
				return ''
			}else if(this.count < 120) {
				return 'md-warn'
			}else {
				return 'md-accent'
			}
		},
		left:function(){
			if(this.count < 10){
				return 'count-1'
			}else if(this.count < 100){
				return 'count-2'
			}else {
				return 'count-3'
			}
		}
	}
}
</script>
<style>
.count {
	position: absolute;
    top: 25%;
    font-size: 6em;
}
.count-1 {
	left: 39%;
}
.count-2 {
	left: 35%;
}
.count-3 {
	left: 26%;
}
</style>