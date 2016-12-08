<template>
    <div id="movement-select-modal" class="modal fade bs-example-modal-lg" tabindex="-1" role="dialog">
        <div class="modal-dialog modal-lg" role="document">
            <div class="modal-content">
                <div class="modal-header">
                    <button type="button" class="close" data-dismiss="modal" aria-label="Close"><span aria-hidden="true">&times;</span></button>
                    <h4>选择动作</h4>
                </div>
                <div class="modal-body">
                    <movement-list @confirmSelection="confirmSelection"></movement-list>
                </div>
            </div>
        </div>
    </div>
</template>
<script>
import MovementList from './MovementList.vue'
export default {
    name:'movement-select-modal',
    data(){
        return {}
    },
    methods: {
        confirmSelection : function(movement){
            this.$emit('movement-modal-selected', movement )
        }
    },
    components:{
        'movement-list': MovementList
    }
}
</script>
