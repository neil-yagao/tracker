<template>
    <div>
        <div class="row">
            <div class="col-md-4">
                <div class="input-group" style="margin-bottom:1em">
                    <span class="input-group-addon">
                        <i class="fa fa-search"  aria-hidden="true"></i>
                    </span>
                    <input type="text" class="form-control" v-model="nameKey">
                </div>
            </div>
            <div class="col-md-1">
                <button class="btn btn-info" @click="confirm" data-dismiss="modal">Select</button>
            </div>
        </div>
        <template v-for="(movement, index) in filtedMovement">
            <movement :movement="movement" :selected="selected" :editing="false" @selectingMovement="selecting(movement)"></movement>
        </template>
    </div>
</template>
<script>
import Movement from "./Movement.vue"
export default {
    name: 'movement-list',
    data() {
        return {
            movements: [],
            selected: -2,
            nameKey: '',
            selectedMovement:{}
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
    methods: {
        selecting: function(movement) {
            if (this.$data.selected == movement.id) {
                this.$data.selected = -2
            } else {
                this.$data.selected = movement.id
            }
            this.$data.nameKey = movement.name
            this.$data.selectedMovement = movement
        },
        confirm: function(){
            this.$emit('confirmSelection', this.$data.selectedMovement)
        }
    },
    created: function() {
        console.info("movement list created");
        this.$http.get('/movements').then((response) => {
            this.$data.movements = response.body.data
        })

    },

    components: {
        Movement
    }
}
</script>
