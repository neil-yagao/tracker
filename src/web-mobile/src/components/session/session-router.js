import WorkoutList from './session-list.vue'
import Session from './Session.vue'
var router = {
    path: 'workouts',
    component: Session,
    children: [{
        path: '',
        component: WorkoutList
    }]
}

export default router