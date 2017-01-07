// The following line loads the standalone build of Vue instead of the runtime-only build,
// so you don't have to do: import Vue from 'vue/dist/vue'
// This is done with the browser options. For the config, see package.json
import Vue from 'vue'
import VueRouter from 'vue-router'
import VueResource from 'vue-resource'

import Session from './components/workouts/Session.vue'
import WorkoutList from './components/workouts/WorkoutList.vue'
import WorkoutCreate from './components/workouts/WorkoutCreate.vue'


import Movements from './components/movements/MovementGeneral.vue'
import MovementCreate from './components/movements/MovementCreate.vue'
import MovementList from './components/movements/MovementList.vue'

import Login from './Login.vue'
import WorkingSpace from './WorkingPage.vue'


Vue.use(VueRouter)
Vue.use(VueResource)



const routes = [{
    path: "/lift",
    component: WorkingSpace,
    children: [{
            path: 'workouts',
            component: WorkoutList
        }, {
            path: 'movements',
            component: Movements,
            children: [{
                path: 'new-movement',
                component: MovementCreate
            }, {
                path: '',
                component: MovementList
            }]
        }, {
            path: 'workout/:id',
            component: Session
        },

        {
            path: 'workout/template/new',
            component: WorkoutCreate
        },
    ]

}, {
    path: '/',
    component: Login
}]


const router = new VueRouter({
    routes // short for routes: routes
})

new Vue({ // eslint-disable-line no-new
    router
}).$mount('#app')
