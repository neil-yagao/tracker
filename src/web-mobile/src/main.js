// The following line loads the standalone build of Vue instead of the runtime-only build,
// so you don't have to do: import Vue from 'vue/dist/vue'
// This is done with the browser options. For the config, see package.json
import Vue from 'vue'
import VueRouter from 'vue-router'
import VueResource from 'vue-resource'
import router from './router.js'
import Vuex from 'vuex'
window.jQuery = require("jquery");
require("bootstrap")
import Vuetify from 'vuetify'

Vue.use(Vuetify)
Vue.use(VueRouter)
Vue.use(VueResource)
Vue.use(Vuex)
window._ = require('lodash');

const store = new Vuex.Store({
    state: {
        username: '',
        newSession: {}
    },
    mutations: {

    }
})

window.Vue = new Vue({ // eslint-disable-line no-new
    store,
    router
}).$mount('#app')