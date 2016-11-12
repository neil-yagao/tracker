<template>
    <div id="create-workout">
        <div id="inputarea">
        </div>
        <movement v-if="generated" :title="movement" :set="sets"></movement>
    </div>
</template>
<script>
import Movement from './Movement.vue'
export default {
    name: 'create-workout',
    data() {
        return {
            movement: "",
            sets: [],
            exsitedMovement: [],
            generated: false
        }
    },
    method: function generate() {
        this.$http.put("/movement", {
            'movement': this.movement,
            'sets': this.sets
        }, (response) => {
            this.$set('movement', response.body.data.movement)
            this.$set('sets', response.body.data.sets)
            this.$set('generated', true)
        })
    },
    create: function() {

    },
    component: {
        Movement
    }
}
</script>
</script>
<style scoped>
</style>
