<template>
<div id="workset">
    <div class="panel panel-default" style="max-width:800px">
        <div class="panel-heading">
            <h4>{{title}}</h4>
        </div>
        <div class="panel-body" style="padding-bottom:0px;">
            <table class="table table-hover table-condensed" style="margin-bottom:0px">
                <thead>
                    <tr>
                        <th>#</th>
                        <th>Reps</th>
                        <th @click="eachSide = false" v-if="eachSide">Each Side(KG)</th>
                        <th @click="eachSide = true" v-if="!eachSide">Total Weight(KG)</th>
                        <th>Finished</th>
                    </tr>
                </thead>
                <tbody>
                    <tr v-for="(set, index) in sets">
                        <td>{{set.sequence}}</td>
                        <td>{{set.repeat}}</td>
                        <td v-if="eachSide">{{set.eachSide}}</td>
                        <td>
                            <editable-input :value="set.totalWeight" @save="changeTotalWeight($event, set)"></editable-input>
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
</template>
<script>

import EditableInput from '../../general/EditableInput.vue'
var debounce
export default {
    name: 'workset',
    data() {
        return {
            editing: true,
            eachSide: false
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
            set.totalWeight = totalWeight + 0;
            var eachSide = (totalWeight - 20) / 2;
            var movementName = _.toLower(set.name)
            if (movementName.indexOf('dumbbell') >= 0 || movementName.indexOf("body") >= 0) {
                eachSide = "--"
            }
            set.eachSide = eachSide
        }

    },
    components: {
        'editable-input': EditableInput
    }
}
</script>
<style scoped>
.btn {
    min-width: 40px
}
</style>
