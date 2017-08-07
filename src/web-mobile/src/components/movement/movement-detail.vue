<template>
    <div style="overflow-y: auto; max-height: 83vh">
    <md-toolbar class="md-transparent">

        	<h2 class="md-title"  style="flex:1">{{movement.name}}</h2>
        	<md-button  class="md-raised md-accent md-icon-button" v-on:click.native="addMovement()">
	  			<md-icon v-if="editing">save</md-icon>
	  			<md-icon v-if="!editing">edit_mode</md-icon>
	      	</md-button>
	      	<md-button  class="md-raised md-primary md-icon-button" v-on:click.native="addMovement()">
	  			<md-icon>videocam</md-icon>
	      	</md-button>
        </md-toolbar>
        <md-list>
            <md-list-item>
                <md-input-container required :class="showError && !movement.targetMuscle?'md-input-invalid':''">
                    <label>主要目标肌群</label>
                    <md-select name="muscle" id="muscle" v-model="movement.targetMuscle" :disabled="!editing">
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
                    <md-select name="muscle" id="muscle" v-model="movement.secondaryMuscleArray" multiple :disabled="!editing">
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
                <md-input-container md-clearable :class="showError && !movement.description?'md-input-invalid':''" >
                    <label>动作描述</label>
                    <md-textarea v-model="movement.description" required :disabled="!editing"></md-textarea>
                </md-input-container>
            </md-list-item>
        </md-list>
        <movement-video></movement-video>
    </div>
</template>
<script>
import MovementVideo from './movement-videos.vue'
export default {
	name:'movement-detail',
	data(){
		return {
			showError:false,
			editing:false
		}
	},
	computed:{
		id: function(){
			return this.$router.params.id
		},
		movement:function(){
			return this.$store.state.movement.movement
		}
	},
	components:{
		'movement-video':MovementVideo
	}
}
</script>