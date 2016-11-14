<template>
    <div id="workout-list">
        <div class="list-group">
            <template v-for="workout in workouts">
                <a :href="'#/workout/' + workout.id" class="list-group-item">
                    <h4 class="list-group-item-heading">{{workout.name}}</h4>
                    <p class="list-group-item-text">{{workout.description}}</p>
                </a>
            </template>
        </div>
        <div class="row">
            <div class="col-lg-1">
                <a type="button" class="btn btn-default" href="#/new-workout">New Workout</a>
            </div>
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
        this.$data.workouts = [{
                id: '1',
                name: 'upper strength',
                description: 'upper strength session to improve upperbody'
            }, {
                id: "2",
                name: 'lower strength',
                description: 'lower strength session to improve squat and deadlift'
            }]
            //this.$http.get("/workouts", (response) => {
            //    this.$set('workouts', response.body.data)
            //})
    }
}
</script>
<style scoped>
</style>
