const movementState = {
    state: {
        movement:{}
    },
    mutations:{
    	setMovement:function(state, movement){
    		console.info('set movement')
    		console.info(movement)
    		state.movement = _.clone(movement)
    	}
    }
}

export default movementState