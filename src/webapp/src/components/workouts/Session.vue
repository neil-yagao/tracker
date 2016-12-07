<template>
<div id="movement-list">
    <div class="row"><a class="btn btn-lick pull-right " href="#/lift/workouts">Back to List</a></div>
    <div class="row">
        <workset v-for="movement in movements" :title="movement.name" :sets="movement.sets" @finishOneSet="finishOneSet"
        @changeMovement="changeMovement($event, movement)" :key="movement.name"></workset>
    </div>
    <div class="row">
        <button type="button" class="btn btn-info" style="max-width:450px; width:100%" @click="finishSession">Finish</button>
    </div>
    <rest-modal rest="60" :startFlag="restingFlag"></rest-modal>
</div>
</template>
<script>
import RestModal from './RestModal.vue'
import WorkSet from './WorkingSet.vue'
export default {
    name: 'movement-list',
    data() {
        return {
            restingFlag: false,
            movements: []
        }
    },
    created: function() {
        this.getAndTransformMovement()
    },
    methods: {
        finishSession: function() {
            console.info(this.$data.movements)
            var sessionData = _.reduce(this.$data.movements, function(arry, value) {
                return _.concat(arry, _.flatMap(value.sets, (set) => {
                    return {
                        'id': set.id,
                        'acheiveNumber': set.achieved,
                        'acheiveWeight': set.totalWeight
                    }
                }))

            }, [])
            console.info(sessionData)
            this.$http.post('/session/' + this.$route.params["id"], JSON.stringify(sessionData))
        },
        getAndTransformMovement:function(){
            let self = this
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
                        if(!set.dividable){
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
                this.$data.movements = beforeAssign
            }).bind(this)
        },
        finishOneSet: function() {
            let self = this
            if (self.debounce) {
                self.debounce.cancel()
            }
            self.debounce = _.debounce(() => {

                $('#rest-modal').modal({
                    'show': true
                })
                self.$data.restingFlag = !self.$data.restingFlag;
            }, 1000)
            self.debounce()
        },
        changeMovement: function(newMovement, old){
            var workingSetId = _.flatMap(old.sets, "id")
            console.info(workingSetId)
            console.info(newMovement)
            this.$http.post('/session/'+this.$route.params["id"]+'/movement', {
                "movement": newMovement.id,
                "workset": workingSetId
            }).then((response) =>{
                this.getAndTransformMovement()
            })
            // this.getAndTransformMovement()
            // this.$data.movements = []
        }
    },
    components: {
        'workset': WorkSet,
        'rest-modal': RestModal
    }
}
</script>
<style scoped>

</style>
