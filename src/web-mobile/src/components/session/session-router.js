import SessionList from './session-list.vue'
import SessionDetail from './session-detail.vue'
import Session from './Session.vue'
var router = {
    path: 'workouts',
    component: Session,
    children: [{
        path: '',
        component: SessionList
    },{
    	path:'detail/:id',
    	component: SessionDetail
    }]
}

export default router