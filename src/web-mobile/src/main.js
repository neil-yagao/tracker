// The following line loads the standalone build of Vue instead of the runtime-only build,
// so you don't have to do: import Vue from 'vue/dist/vue'
// This is done with the browser options. For the config, see package.json
import Vue from 'vue'
import VueRouter from 'vue-router'
import VueResource from 'vue-resource'
import router from './router.js'
import Vuex from 'vuex'
require('../node_modules/bootstrap/dist/css/bootstrap.css')
import VueMaterial from 'vue-material'
import planState from './components/plan/plan-state.js'
import sessionState from './components/session/session-state.js'

Vue.use(VueMaterial)
Vue.use(VueRouter)
Vue.use(VueResource)
Vue.use(Vuex)
window._ = require('lodash');

const store = new Vuex.Store({

    state: {
        user: {
        }
    },
    mutations: {
        setUser(state, user) {
            state.user = user
        }
    },
    modules: {
        plan: planState,
        session: sessionState
    }
})


window.Vue = new Vue({ // eslint-disable-line no-new
    store,
    router
}).$mount('#app')