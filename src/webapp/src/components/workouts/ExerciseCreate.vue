<template>
    <div id="movement-create">
        <div class="row">
            <div class="col-md-2 col-md-offset-2">
                <div class="input-group">
                    <input type="text" class="form-control" v-model="name" placeholder="Movement">
                     <span class="input-group-btn">
                    <button class="btn btn-default" data-toggle="modal" data-target="#movement-select-modal"><i class="fa fa-search"  aria-hidden="true"></i></button
                    </span>
                </div>
            </div>
            <div class="col-md-2">
                <div class="input-group">
                    <input type="number" class="form-control" v-model="weight" placeholder="Weight">
                    <span class="input-group-addon">
                        KG
                    </span>
                </div>
            </div>
            <div class="col-md-1">
                <div class="input-group">
                    <input type="number" class="form-control" v-model="repeats" placeholder="Reps">
                </div>
            </div>
            <div class="col-md-1">
                <div class="input-group">
                    <input type="number" class="form-control" v-model="sets" placeholder="Sets">
                </div>
            </div>

            <div class="col-md-1">
                <button type="button" class="btn btn-default" @click="addMovement()">Add</button>
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
            'sets': this.movement.sets
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
                'sets': this.$data.sets
            })
            this.name = '';
            this.repeats = '';
            this.weight = '';
            this.sets = ''
        },
        selected: function(movement){
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
