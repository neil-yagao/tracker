<template>
    <div id="movement-list">
        <div class="row">
            <div class="col-md-4">
                <div class="input-group" style="margin-bottom:1em">
                    <span class="input-group-addon">
                        Seaching Name
                    </span>
                    <input type="text" class="form-control" v-model="nameKey">
                </div>
            </div>
            <div class="col-md-1">
                <button class="btn btn-success" data-toggle="modal" data-target="#movement-modal">Add</button>
            </div>
        </div>
        <movement-modal></movement-modal>
        <template v-for="(movement, index) in filtedMovement">
            <movement :movement="movement" :selected="selected" :editing="false" @selectingMovement="selecting"></movement>
        </template>
    </div>
</template>
<script>
import Movement from './Movement.vue'
import MovementModal from './MovementAddModal.vue'
export default {
    name: 'movement-list',
    data() {
        return {
            movements: [],
            selected: -2,
            nameKey: ''
        }
    },
    created: function() {
        console.info("movement list created");
        this.$http.get('/movements').then((response) => {
            console.info("get movements response:")
            console.info(response)
            this.$data.movements = response.body.data
        })

    },
    methods: {
        selecting: function(index) {
            if (this.$data.selected == index) {
                this.$data.selected = -2
            } else {
                this.$data.selected = index
            }
        }
    },
    computed: {
        filtedMovement: function() {
            var filtered = _.filter(this.$data.movements, (m) => {
                var movementName = _.toLower(m.name)
                var searchKey = _.toLower(this.$data.nameKey)
                return movementName.indexOf(searchKey) > -1
            })
            return _.slice(_.unionBy(filtered, 'name'),0 , 20)
        }
    },
    components: {
        'movement':Movement,
        'movement-modal': MovementModal

    }
}
</script>
<style scoped>
</style>
