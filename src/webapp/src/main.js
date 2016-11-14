// The following line loads the standalone build of Vue instead of the runtime-only build,
// so you don't have to do: import Vue from 'vue/dist/vue'
// This is done with the browser options. For the config, see package.json
import Vue from 'vue'
import VueRouter from 'vue-router'
import VueResource from 'vue-resource'
import WorkoutList from './components/WorkoutList.vue'
import MovementList from './components/MovementList.vue'
import WorkoutCreate from './components/Workout-Create.vue'

Vue.use(VueRouter)
Vue.use(VueResource)
const routes = [{
    path: '/',
    component: WorkoutList
}, {
    path: '/movements',
    component: MovementList
},/* {
     path: '/workout/:id',
     component: Workout
 }
*/, {
    path: '/new-workout',
    component: WorkoutCreate
}]


const router = new VueRouter({
    routes // short for routes: routes
})

new Vue({ // eslint-disable-line no-new
    router
}).$mount('#app')
