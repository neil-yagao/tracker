<template>
    <movement :movement="mv" editing="true">
</movement>
</template>
<script>
import Movement from "./Movement.vue"
export default {
    name: 'movement-create-tab',
    data() {
        return {
            mv: {
                targetMuscle: ""
            }
        }
    },

    components: {
        Movement
    }
}
</script>
<style></style>
