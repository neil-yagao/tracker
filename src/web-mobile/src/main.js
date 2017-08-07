// The following line loads the standalone build of Vue instead of the runtime-only build,
// so you don't have to do: import Vue from 'vue/dist/vue'
// This is done with the browser options. For the config, see package.json
import Vue from 'vue'
import VueRouter from 'vue-router'
import VueResource from 'vue-resource'
import router from './router.js'
import Vuex from 'vuex'
import VueMaterial from 'vue-material'
import planState from './components/plan/plan-state.js'
import sessionState from './components/session/session-state.js'
import movementState from './components/movement/movement-state.js'


Vue.use(VueMaterial)
Vue.use(VueRouter)
Vue.use(VueResource)
Vue.use(Vuex)
window._ = require('lodash');

const store = new Vuex.Store({

    state: {
        user: {},
        goback:'',
        guide:{
            new:false,
            instruction:'aaa'
        }
    },
    mutations: {
        setUser(state, user) {
            state.user = user
        },
        setBackUrl(state, url){
            state.goback = url;
        },
        setGuideInfo(state, guide){
            state.guide.new = guide.new;
            state.guide.instruction = guide.instruction;
        }
    },
    modules: {
        plan: planState,
        session: sessionState,
        movement:movementState
    }
})


window.Vue = new Vue({ // eslint-disable-line no-new
    store,
    router
}).$mount('#app')