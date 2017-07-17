<template>
<div>
    <md-input-container v-on:click.native="showDialog()">
        <label>训练动作</label>
        <md-input v-model="name"></md-input>
    </md-input-container>
    <md-dialog ref="movementSelection">
    	<movement-list v-if="step == 'list'" @to-add="step = 'add'" :movements="movements" 
    	@close="closeDialog()" @selected="selectedMovement($event)">
    	</movement-list>
    	<movement-add v-if="step == 'add'" @to-list="step = 'list'" @add-movement="addMovements($event)"></movement-add>
    </md-dialog>
</div>


</template>
<script type="text/javascript">
import MovementList from './movement-list.vue'
import MovementAdd from './movement-add.vue'
export default {

    name:'movement-selection',
    data(){
        return {
            name:'',
            step:'list',
            movements:[]
        }
    },
    methods:{
        showDialog(){
            this.$refs['movementSelection'].open();
            this.step = 'list'
        },
        closeDialog(){
            this.$refs['movementSelection'].close();
        },
        loadMovements(){
        	this.$http.get('movements').then((res)=>{
        		this.movements = res.data
        	})
        },
        addMovements(movement){
        	this.movements.push(movement);
        	this.step = 'list'
        },
        selectedMovement(movement){
        	this.name = movement.name;
        	this.closeDialog();
        	this.$emit('selected-movement', movement)
        }
    },
    components:{
    	'movement-list': MovementList,
    	'movement-add':MovementAdd
    }

} 
</script>
<style>
.md-dialog{width:90%;height: 70%}
#criteria {
    margin-left: 10px
}
</style>