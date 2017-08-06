import Movement from './movement-main.vue'
import MovementDetail from './movement-detail.vue'

const router = [{
	path:'movement',
	component:Movement
},{
	path:'movement/:id',
	component:MovementDetail
}
]
export default router