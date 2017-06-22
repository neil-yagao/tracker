import Login from './Login.vue'
import WorkingSpace from './WorkingPage.vue'
import Plan from './components/plan/plan.vue'

import VueRouter from 'vue-router'


const routes = [{
    path: "/working",
    component: WorkingSpace,
    children: [{
        'path': 'plan',
        component: Plan
    }]
}, {
    path: '/',
    component: Login
}]

const router = new VueRouter({
    routes // short for routes: routes
})

export default router;