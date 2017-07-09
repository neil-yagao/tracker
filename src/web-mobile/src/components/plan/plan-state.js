const planState = {
	state:{
		planName: '',
		sessions: [],
		index: 0,
		planEditing: true
	},
	mutations: {
		pushSession(state, session) {
			state.sessions.push(session);
		},
		removeSession(state, session) {
			state.sessions = _.without(state.sessions, session);
		},
		pushExercise(state, exercise) {
			var preAdd = state.sessions[state.index];
			preAdd.workouts.push(exercise)
			state.sessions.splice(state.index, 1, preAdd);
		},
		removeExercise(state, exercise) {
			state.sessions[state.index].workouts = _.without(state.sessions[state.index].workouts, exercise)
		},
		editSession(state, index) {
			state.index = index;
		},
		reset(state){
			state.sessions = [];
			state.planName = '';
		},
		editingPlan(state){
			state.planEditing = true;
		},
		readOnlyPlan(state){
			state.planEditing = false;
		}
	}
}
export default planState;