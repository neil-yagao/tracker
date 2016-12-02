<template>
    <div id="movement-list">
        <div class="row"><a class="btn btn-lick pull-right " href="#/lift/workouts">Back to List</a></div>
        <div class="row">
            <workset v-for="movement in movements" :title="movement.name" :sets="movement.sets"></workset>
        </div>
        <div class="row">
            <button type="button" class="btn btn-info" style="max-width:450px; width:100%" @click="finishSession">Finish</button>
        </div>
    </div>
</template>
<script>
import WorkSet from './WorkingSet.vue'
export default {
    name: 'movement-list',
    data() {
        return {
            movements: []
        }
    },
    created: function() {
        this.$http.get("/session/" + this.$route.params["id"], {
            'params': {
                'user': window.localStorage.getItem('user')
            }
        }).then((response) => {
            var sets = response.body.data;
            var movements = _.groupBy(sets, 'movementName')
            console.info(movements)
            var beforeAssign = _.map(movements, function(value, key) {
                var sequence = 0;
                var newSets = _.map(value, function(set) {
                    var eachSide = (set.targetWeight - 20) / 2;
                    var movementName =_.toLower(key)
                    if (movementName.indexOf('dumbbell') >= 0 || movementName.indexOf("body") >= 0) {
                        eachSide = "--"
                    }
                    sequence += 1
                    return {
                        'id': set.id,
                        'sequence': sequence,
                        'repeat': set.targetNumber,
                        'eachSide': eachSide,
                        'totalWeight': set.targetWeight,
                        'achieved': -1
                    }
                })
                return {
                    "name": key,
                    "sets": newSets
                }
            })
            console.info("beforeAssign");
            console.info(beforeAssign);
            this.$data.movements = beforeAssign;

        })
    },
    methods: {
        finishSession: function() {
            console.info(this.$data.movements)
            var sessionData = _.reduce(this.$data.movements, function (arry, value){
                return _.concat(arry, _.flatMap(value.sets, (set)=>{
                    return {
                        'id': set.id,
                        'acheiveNumber': set.achieved,
                        'acheiveWeight' : set.totalWeight
                    }
                }))

            }, [])
            console.info(sessionData)
            this.$http.post('/session/'+  this.$route.params["id"], JSON.stringify(sessionData))
        }
    },
    components: {
        'workset': WorkSet
    }
}
</script>
<style scoped>
</style>
