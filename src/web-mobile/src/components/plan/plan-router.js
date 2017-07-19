import PlanSession from './plan-session.vue'
import PlanSessions from './plan-sessions.vue'
import PlanList from './plan-list.vue'
import Plan from './plan.vue'

var router = {
    path: "plan",
    component: Plan,
    children: [{
        path: "",
        component: PlanList
    }, {
        path: ":id/sessions",
        component: PlanSessions
    }, {
        path: ":session/session/:id",
        component: PlanSession
    }]
}


export default router;