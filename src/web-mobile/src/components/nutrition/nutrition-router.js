import NutritionCollect from './nutrition-collect.vue'

import Nutrition from './nutrition.vue'

var router = {
    path: "nutrition",
    component: Nutrition,
    children: [{
        path: "",
        component: NutritionCollect
    }]
}


export default router;