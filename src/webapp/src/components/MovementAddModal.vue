<template>
<div id="movement-modal" class="modal fade bs-example-modal-lg" tabindex="-1" role="dialog">
    <div class="modal-dialog modal-lg" role="document">
        <div class="modal-content">
            <div class="modal-header">
                <button type="button" class="close" data-dismiss="modal" aria-label="Close"><span aria-hidden="true">&times;</span></button>
                <h4>Add New Movement</h4>
            </div>

                <movement :movement="mv" editing="true"></movement>
        </div>
    </div>
</div>
</template>
<script>
import Movement from "./Movement.vue"
export default {
    name: 'movement-modal',
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
