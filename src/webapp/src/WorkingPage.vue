<template>
<div>
    <!-- route outlet -->
    <!-- component matched by the route will render here -->
    <ul class="nav nav-pills">
        <li role="presentation" :class="{'active': $route.path.indexOf('workout') >=0 }"><a href="#/lift/workouts">Available Sessions</a></li>
        <li role="presentation" :class="{'active': $route.path.indexOf('movement') >=0 }"><a href="#/lift/movements">Movements</a></li>
    </ul>
    <router-view class="container-fluid" style="margin-top: 2em"></router-view>
</div>
</template>
<script>
export default {
    name :'working-page',
    data() {
        return {}
    }
}
</script>
<style></style>
