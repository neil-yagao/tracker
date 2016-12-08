<template>
    <div id="workout-list">
        <div class="list-group">
            <template v-for="workout in workouts">
                <a :href="'#/lift/workout/' + workout.id" class="list-group-item">
                    <h4 class="list-group-item-heading" style="display: inline-block">{{workout.name}}</h4>
                    <h6 class="pull-right">{{workout.perform_date}}</h6>
                    <p class="list-group-item-text">{{workout.description}}</p>
                </a>
            </template>
        </div>
        <div class="row">
            <div class="col-lg-1">
                <a type="button" class="btn btn-default" href="#/lift/workout/template/new">创建新的训练课</a>
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

        this.$http.get("/workouts", {
            'params': {
                'user': window.localStorage.getItem('user')
            }
        }).then((response) => {
            console.info(response.body)
            this.$data.workouts = response.body.data
        })
    }
}
</script>
<style scoped>
</style>
