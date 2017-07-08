import PlanBasic from './plan-basic.vue'
import PlanSession from './plan-session.vue'
import PlanSessions from './plan-sessions.vue'
import Plan from './plan.vue'

var router = {
    path: "plan",
    component: Plan,
    children: [{
        path: "",
        component: PlanBasic
    }, {
        path: "sessions",
        component: PlanSessions
    }, {
        path: "per-session/:id",
        component: PlanSession
    }]
}


export default router;