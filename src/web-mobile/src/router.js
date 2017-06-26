import Login from './Login.vue'
import WorkingSpace from './WorkingPage.vue'
import PlanRouter from './components/plan/plan-router.js'

import VueRouter from 'vue-router'


const routes = [{
    path: "/working",
    component: WorkingSpace,
    children: [PlanRouter]
}, {
    path: '/',
    component: Login
}]

const router = new VueRouter({
    routes // short for routes: routes
})

export default router;