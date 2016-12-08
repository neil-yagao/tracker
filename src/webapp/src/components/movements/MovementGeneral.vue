<template>
    <div id="movement-general">
        <ul class="nav nav-tabs nav-justified" style="max-width:600px">
          <li role="presentation" :class="{'active': $route.path.indexOf('movements') > -1}"><a href="#/lift/movements">浏览动作</a></li>
          <li role="presentation" :class="{'active': $route.path.indexOf('new-movement')> -1 }"><a href="#/lift/movements/new-movement">未找到？自己加一个!</a></li>
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
