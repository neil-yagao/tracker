<template>
<div id="workset">
    <div class="panel panel-default" style="max-width:800px">
        <div class="panel-heading">
            <div class="row">
                <div class="col-md-4">
                    <h4>{{movementTitle}}<h4>
                </div>
                <div class="col-md-2 col-md-offset-6">
                    <button class="btn btn-link pull-right"  data-toggle="modal" data-target="#movement-select-modal">换训练动作</button>
                </div>
            </div>
        </div>
        <div class="panel-body" style="padding-bottom:0px;">
            <table class="table table-hover table-condensed" style="margin-bottom:0px">
                <thead>
                    <tr>
                        <th>#</th>
                        <th>每组目标</th>
                        <th @click="eachSide = false" v-if="eachSide">每边重量(KG)</th>
                        <th @click="eachSide = true" v-if="!eachSide">总重量(KG)</th>
                        <th>完成</th>
                    </tr>
                </thead>
                <tbody>
                    <tr v-for="(set, index) in workingSets">
                        <td>{{set.sequence}}</td>
                        <td><editable-input type="number" :value="set.repeat" @save="set.repeat = $event + 0"></editable-input></td>
                        <td v-if="eachSide">{{set.eachSide}}</td>
                        <td v-if="!eachSide">
                            <editable-input type="number" :value="set.totalWeight" @save="changeTotalWeight($event, set)"></editable-input>
                        </td>
                        <td>
                            <button @click="handleAchieve(set)" :class="['btn', {'btn-success': set.achieved== set.repeat,
                                'btn-warning': set.achieved < set.repeat && set.achieved > 0 ,
                                'btn-danger' : set.achieved == 0,
                                'btn-default' : set.achieved == -1}]">
                                    {{set.achieved == -1 ? '--' : set.achieved}}
                                </button>
                        </td>
                    </tr>
                </tbody>
            </table>
        </div>
    </div>
    <movement-select-modal @movement-modal-selected="selected"></movement-select-modal>
</template>
<script>

import EditableInput from '../../general/EditableInput.vue'
import MovementSelectModal from '../movements/MovementSelectModal.vue'
var debounce
export default {
    name: 'workset',
    data() {
        return {
            editing: true,
            eachSide: false,
            workingSets: this.sets,
            movementTitle: this.title
        }
    },
    props: ['title', 'sets'],
    methods: {
        handleAchieve: function(set) {
            if (set.achieved == undefined || set.achieved <= 0) {
                set.achieved = set.repeat;
            } else {
                set.achieved -= 1
            }
            this.$emit('finishOneSet')
        },
        changeTotalWeight: function(totalWeight, set) {
            set.totalWeight = Number(totalWeight);
            var eachSide = (totalWeight - 20) / 2;
            if(!set.dividable){
                eachSide = "--"
            }
            set.eachSide = eachSide
        },
        selected: function(movement){
            this.$emit('changeMovement', movement)
        }

    },
    beforeUpdate: function(){
        this.$data.workingSets = this.sets;
        this.$data.movementTitle = this.title;
    },
    components: {
        'editable-input': EditableInput,
        'movement-select-modal': MovementSelectModal
    }
}
</script>
<style scoped>
.btn {
    min-width: 40px
}
</style>
