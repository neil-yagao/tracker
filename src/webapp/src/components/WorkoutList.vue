<template>
    <div id="workout-list">
        <div class="list-group">
            <template v-for="workout in workouts">
                <a href="#" class="list-group-item">
                    <h4 class="list-group-item-heading">{{workout.name}}}</h4>
                    <p class="list-group-item-text">{{workout.description}}</p>
                </a>
            </template>
        </div>
    </div>
</template>
<script>
export default {
    name: 'workout-list',
    data() {
        return {
            workouts: []
        }
    },
    created: function() {
        this.$http.get("/workouts", (response) => {
            this.$set('workouts', response.body.data)
        })
    }
}
</script>
<style scoped>
</style>
