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
require('../node_modules/bootstrap/dist/css/bootstrap.css')
import VueMaterial from 'vue-material'


Vue.use(VueMaterial)
Vue.use(VueRouter)
Vue.use(VueResource)
Vue.use(Vuex)
window._ = require('lodash');

const store = new Vuex.Store({

	state: {
		username: '',
		newSession: {
			sessions: []
		}
	},
	mutations: {
		pushSession(state, session) {
			state.newSession.sessions.push(sessions);
		},
		removeSessions(state, session) {
			state.newSession.sessions = _.without(state.newSession.sessions, session);
		},
		pushExercise(state, index, exercise) {
			if (state.newSession.sessions[index].workouts) {
				state.newSession.sessions[index].workouts.push(exercise)
			} else {
				state.newSession.sessions[index].workouts = [exercise]
			}
		}，
		removeExercise(state, index, exercise) {
			state.newSession.sessions[index].workouts = _.without(state.newSession.sessions[index].workouts, exercise)
		}
	}
})

window.Vue = new Vue({ // eslint-disable-line no-new
	store,
	router
}).$mount('#app')