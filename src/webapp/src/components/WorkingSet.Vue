<template>
    <div id="workset">
        <div class="panel panel-default">
            <div class="panel-heading">
                <h4>{{title}}</h4>
            </div>
            <div class="panel-body" style="padding-bottom:0px">
                <table class="table table-hover table-condensed" style="margin-bottom:0px">
                    <thead>
                        <tr>
                            <th>#</th>
                            <th>Repeat</th>
                            <th>Each Side(KG)</th>
                            <th>Total Weight(KG)</th>
                            <th>Finished</th>
                        </tr>
                    </thead>
                    <tbody>
                        <tr v-for="(set, index) in sets">
                            <td>{{set.sequence}}</td>
                            <td>{{set.repeat}}</td>
                            <td>{{set.eachSide}}</td>
                            <td>{{set.totalWeight}}</td>
                            <td>
                                <button @click="handleAchieve(set)" :class="['btn', {'btn-success': set.achieved== set.repeat, 'btn-warning': set.achieved < set.repeat && set.achieved > 0 , 'btn-danger' : set.achieved == 0, 'btn-default' : set.achieved == -1}]">{{set.achieved == -1 ? '--' : set.achieved}}</button>
                            </td>
                        </tr>
                    </tbody>
                </table>
            </div>
        </div>
</template>
<script>
export default {
    name: 'workset',
    props: ['title', 'sets'],
    methods: {
        handleAchieve: function(set) {
            if (set.achieved == undefined || set.achieved <= 0) {
                set.achieved = set.repeat;
            } else {
                set.achieved -= 1
            }
        }
    }
}
</script>
<style scoped>
.btn {
    min-width: 40px
}
</style>
