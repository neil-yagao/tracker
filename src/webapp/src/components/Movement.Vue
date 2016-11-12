<template>
    <div id="movement">
        <div id="workset-header">{{movement}}</div>
        <table class="table table-hover">
            <thead>
                <tr>
                    <th>Repeat</th>
                    <th>Each Side</th>
                    <th>Total Weight</th>
                    <th>Finished</th>
                </tr>
            </thead>
            <tbody>
                <tr v-for="set in sets">
                    <td>{{set.repeat}}</td>
                    <td>{{set.eachSide}}</td>
                    <td>{{set.totalWeight}}</td>
                    <td>
                        <button @click="handleAchieve(set)">{{set.achieved}}</button>
                    </td>
                </tr>
            </tbody>
        </table>
    </div>
</template>
<script>
export default {
    name: 'movement',
    props: ['movement', 'sets'].
    methods: {
        handleAchieve: function(set) {
            if (set.achieved == undefined || set.achieved <= 0) {
                set.achieved = set.defaultRepeat;
            } else {
                set.achieved -= 1
            }
        }
    }
}
</script>
<style scoped>
</style>
