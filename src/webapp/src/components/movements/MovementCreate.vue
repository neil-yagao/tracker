<template>
    <movement :movement="mv" :editing="editing" @saveMovement="reloadMovement">
</movement>
</template>
<script>
import Movement from "./Movement.vue"
export default {
    name: 'movement-create-tab',
    data() {
        return {
            mv: {
                targetMuscle: "",
                dividable:1
            },
            editing: true
        }
    },
    methods:{
        reloadMovement: function(){
            this.$data.mv = {
                targetMuscle: "",
                dividable:1
            },
            this.$data.editing = true
        }
    },
    components: {
        Movement
    }
}
</script>
<style></style>
