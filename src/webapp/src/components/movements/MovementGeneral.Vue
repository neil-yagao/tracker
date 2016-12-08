<template>
    <div id="movement-general">
        <ul class="nav nav-tabs nav-justified" style="max-width:600px">
          <li role="presentation" :class="{'active': _.endsWith($route.path, 'movements') }"><a href="#/lift/movements">All Movements</a></li>
          <li role="presentation" :class="{'active': _.endsWith($route.path, 'new-movement') }"><a href="#/lift/movements/new-movement">Not Found? Add One Here!</a></li>
        </ul>
        <router-view class="container-fluid" style="margin-top:0.5em"></router-view>
    </div>
</template>
<script>
export default {
    name:'movement-general',
    data(){
        return {}
    }
}
</script>
<style scoped>
</style>
