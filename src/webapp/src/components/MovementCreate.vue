<template>
    <div id="movement-create">
        <div class="row">
            <div class="col-md-3 col-md-offset-2">
                <div class="input-group">
                    <span class="input-group-addon">
                        Movement
                    </span>
                    <input type="text" class="form-control" v-model="name">
                </div>
            </div>
            <div class="col-md-2">
                <div class="input-group">
                    <span class="input-group-addon">
                        Weight
                    </span>
                    <input type="text" class="form-control" v-model="weight">
                    <span class="input-group-addon">
                        KG
                    </span>
                </div>
            </div>
            <div class="col-md-1">
                <div class="input-group">
                    <span class="input-group-addon">
                        Sets
                    </span>
                    <input type="text" class="form-control" v-model="sets">
                </div>
            </div>
            <div class="col-md-1">
                <div class="input-group">
                    <span class="input-group-addon">
                        Reps
                    </span>
                    <input type="text" class="form-control" v-model="repeats">
                </div>
            </div>
            <div class="col-md-1">
                <button type="button" class="btn btn-default" @click="addMovement()">Add</button>
            </div>
        </div>
    </div>
</template>
<script>
export default {
    name: 'movement-create',
    data() {
        return {
            'name': this.movement.name,
            'repeats': this.movement.repeats,
            'weight': this.movement.weight,
            'sets': this.movement.sets
        }
    },
    props: {
        'movement': {}
    },
    methods: {
        addMovement: function() {
            this.$emit('addMovement', {
                'name': this.$data.name,
                'weight': this.$data.weight,
                'repeats': this.$data.repeats,
                'sets': this.$data.sets
            })
            this.name = '';
            this.repeats = '';
            this.weight = '';
            this.sets = ''
        }
    }

}
</script>
<style scoped>
.input-group-addon {
    padding-right: 3px;
    padding-left: 2px
}
</style>
