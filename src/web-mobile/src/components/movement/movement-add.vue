<template>
	<div style="height:45em">
		<md-dialog-title>
		 <md-toolbar class="md-transparent">
        	<h2 class="md-title"  style="flex:1">添加新动作</h2>
	        <md-button  class="md-raised md-accent" v-on:click.native="moveToList()">
	  			返回
	      	</md-button>
        </md-toolbar>
        </md-dialog-title>
		<md-dialog-content>
		<div style="overflow-y: auto; max-height: 23em">
		<md-list>
			<md-list-item>
			<md-input-container>
				<label>名称</label>
    			<md-input v-model="movement.name" required></md-input>
			</md-input-container>
			</md-list-item>

			<md-list-item>
			  <md-input-container required>
                <label>主要目标肌群</label>
                <md-select name="muscle" id="muscle" v-model="movement.targetMuscle">
                    <md-option value="肩部">肩部</md-option>
                    <md-option value="胸部">胸部</md-option>
                    <md-option value="背部">背部</md-option>
                    <md-option value="二头肌">二头肌</md-option>
                    <md-option value="三头肌">三头肌</md-option>
                    <md-option value="核心">核心</md-option>
                    <md-option value="腿部">腿部</md-option>
                </md-select>
            </md-input-container>
            </md-list-item>
            <md-list-item>
			  <md-input-container>
                <label>辅助肌群</label>
                <md-select name="muscle" id="muscle" v-model="movement.secondaryMuscleArray" multiple>
                    <md-option value="肩部">肩部</md-option>
                    <md-option value="胸部">胸部</md-option>
                    <md-option value="背部">背部</md-option>
                    <md-option value="肱二头肌">肱二头肌</md-option>
                    <md-option value="肱三头肌">肱三头肌</md-option>
                    <md-option value="核心">核心</md-option>
                    <md-option value="腿部">腿部</md-option>
                </md-select>
            </md-input-container>
            </md-list-item>
            <md-list-item>
			<md-input-container md-clearable>
				<label>动作描述</label>
    			<md-textarea v-model="movement.description" required ></md-textarea>
			</md-input-container>
			</md-list-item>
			</md-list>
			</div>
			 <md-button  class="md-raised md-primary" v-on:click.native="addMovement()">
	  			添加
	      	</md-button>
		</md-dialog-content>
	</div>
</template>
<script>
export default {
	name:'movement-add',
	data(){
		return{ 
			movement:{
				name:'',
				description:'',
				targetMuscle:'',
				secondaryMuscleArray:[]
			}
		}
	},
	methods:{
		moveToList(){
			this.$emit('to-list')
		},
		addMovement(){
			this.movement.secondaryMuscle = _.join(this.movement.secondaryMuscleArray,";")
			this.movement.addBy = this.$store.state.user
			console.info(this.movement)
			this.$emit('add-movement', this.movement)
			this.$http.put('movement',this.movement).then(() =>{
				this.moveToList()
			})
		}
	}
}
</script>