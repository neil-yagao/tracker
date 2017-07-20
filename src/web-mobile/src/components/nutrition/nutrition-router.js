import NutritionCollect from './nutrition-collect.vue';
import NutritionResult from './nutrition-calculate.vue';
import Nutrition from './nutrition.vue'

var router = {
    path: "nutrition",
    component: Nutrition,
    children: [{
        path: "collect",
        component: NutritionCollect
    },{
    	path:"",
    	component: NutritionResult
    }
    ]
}

export default router;