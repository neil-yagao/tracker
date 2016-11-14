<template>
    <div id="workout-create" class="row">
        <div class="input-group input-group-lg">
            <span class="input-group-addon">Workout</span>
            <input type="text" class="form-control" id="basic-url">
        </div>
        <movement-create :movement="currentMovement"></movement-create>
        <movement-create v-for="movement in movements" :movement="movement"></movement-create>
</template>
<script>
import MovementCreate from "./Movement-Create.vue"
export default {
    name: 'workout',
    data() {
        return {
            workout: {},
            movements: [],
            currentMovement: {}
        }
    },
    created: function() {
        this.$on('addMovement', function(newMovement) {
            this.$data.currentMovement = {}
            this.$data.movements.push(newMovement)
        })
        this.$data.movements = []
    },
    components: {
        'movement-create': MovementCreate
    },
    computed: {
        currentMovement: function() {
            if (this.$data.movements != null || this.$data.movements.length > 0) {
                return this.$data.movements[this.$data.movements.length - 1]
            } else {
                return {}
            }
        }
    }
}
</script>
<style scoped>
</style>
