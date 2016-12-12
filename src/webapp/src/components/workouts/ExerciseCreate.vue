<template>
    <div id="movement-create">
        <div class="row">
            <div class="col-md-2 col-md-offset-2">
                <div class="input-group">
                    <input type="text" class="form-control" v-model="name" placeholder="训练动作">
                    <span class="input-group-btn">
                    <button class="btn btn-default" data-toggle="modal" data-target="#movement-select-modal"><i class="fa fa-search"  aria-hidden="true"></i></button
                    </span>
                </div>
            </div>
            <div class="col-md-2">
                <div class="input-group">
                    <input type="number" class="form-control" v-model="weight" placeholder="训练重量">
                    <span class="input-group-addon">
                        KG
                    </span>
                </div>
            </div>
            <div class="col-md-1">
                <div class="input-group">
                    <input type="number" class="form-control" v-model="repeats" placeholder="每组">
                </div>
            </div>
            <div class="col-md-1">
                <div class="input-group">
                    <input type="number" class="form-control" v-model="sets" placeholder="组数">
                </div>
            </div>
            <div class="col-md-2">
                <div class="checkbox">
                    <label>
                        <input type="checkbox" v-model="needWarmSet">为动作添加热身组
                    </label>
                </div>
            </div>
            <div class="col-md-1">
                <button type="button" class="btn btn-default" @click="addMovement()">添加</button>
            </div>
        </div>
        <movement-select-modal @movement-modal-selected="selected"></movement-select-modal>
    </div>
</template>
<script>
import MovementSelectModal from '../movements/MovementSelectModal.vue'
export default {
    name: 'movement-create',
    data() {
        return {
            'name': this.movement.name,
            'repeats': this.movement.repeats,
            'weight': this.movement.weight,
            'sets': this.movement.sets,
            'needWarmSet': 0
        }
    },


    props: {
        'movement': {}
    },
    methods: {
        addMovement: function() {
            this.$emit('addMovement', {
                'name': this.$data.name,
                'weight': this.$data.weight,
                'repeats': this.$data.repeats,
                'sets': this.$data.sets,
                'needWarmSet': this.$data.needWarmSet ? 1 : 0
            })
            this.name = '';
            this.repeats = '';
            this.weight = '';
            this.sets = ''
        },
        selected: function(movement) {
            this.$data.name = movement.name
        }
    },
    components: {
        'movement-select-modal': MovementSelectModal
    }

}
</script>
<style scoped>
.input-group-addon {
    padding-right: 3px;
    padding-left: 2px
}
</style>
