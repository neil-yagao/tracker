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
                        <td @click="editing = false" v-if="editing && !eachSide">{{set.totalWeight}}</td>
                        <td v-if="!editing && !eachSide">
                            <div class="input-group">
                                <input type="text" class="form-control" v-model="set.totalWeight">
                                <span class="input-group-btn">
                                    <button class="btn btn-default" type="button" @click="saveTotalWeight(set)">OK</button>
                                </span>
                            </div>
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
        <rest-modal rest="60" :startFlag="restingFlag"></rest-modal>
</template>
<script>
import RestModal from './RestModal.vue'
var debounce
export default {
    name: 'workset',
    data() {
        return {
            restingFlag: false,
            editing: true,
            eachSide:false
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
            let self = this
            if (debounce) {
                debounce.cancel()
            }
            debounce = _.debounce(() => {
                self.$data.restingFlag = !self.$data.restingFlag;
                $('#rest-modal').modal({
                    'show': true
                })
            }, 1000)
            debounce()
        }
        ,saveTotalWeight: function(set){
            this.$data.editing = true
        }
    },
    components: {
        'rest-modal': RestModal
    }
}
</script>
<style scoped>
.btn {
    min-width: 40px
}
.input-group {
    max-width: 200px
}
</style>
