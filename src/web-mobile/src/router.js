import Login from './Login.vue'
import WorkingSpace from './WorkingPage.vue'
import PlanRouter from './components/plan/plan-router.js'
import SessionRouter from './components/session/session-router.js'
import NutritionRouter from './components/nutrition/nutrition-router.js'

import VueRouter from 'vue-router'


const routes = [{
    path: "/working",
    component: WorkingSpace,
    children: [PlanRouter, SessionRouter, NutritionRouter]
}, {
    path: '/',
    component: Login
}]

const router = new VueRouter({
    routes // short for routes: routes
})

export default router;