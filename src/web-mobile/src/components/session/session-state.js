const sessionState = {
    state: {
        detail: {}
    },
    mutations: {
        startExecutingSession(state, session) {
            state.detail = session
        }
    }
}
export default sessionState;