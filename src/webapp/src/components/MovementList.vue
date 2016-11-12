<template>
    <div id="movement-list">
        <movement v-for="movement in movements" :title="movement.name" :sets="movement.sets"></movement>
    </div>
</template>
<script>
import Movement from './Movement.vue'
export default {
    name: 'movement-list',
    data() {
        return {
            movements: []
        }
    },
    created: function() {
        this.$http.get("/movements", (response) => {
            this.$set('movements', response.body.data)
        })
    },
    component: {
        Movement
    }
}
</script>
<style scoped>
</style>
