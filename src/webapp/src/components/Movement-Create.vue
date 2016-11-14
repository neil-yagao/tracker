<template>
    <div id="movement-create" class="row head-space">
        <div class="row">
            <div class="col-lg-4 col-lg-offset-2">
                <div class="input-group">
                    <span class="input-group-addon">
        				Movement
      				</span>
                    <input type="text" class="form-control" v-model="name">
                </div>
            </div>
            <div class="col-lg-2">
                <div class="input-group">
                    <span class="input-group-addon">
        				Target Weight
      				</span>
                    <input type="text" class="form-control" v-model="weight">
                </div>
            </div>
            <div class="col-lg-2">
                <div class="input-group">
                    <span class="input-group-addon">
        				Repeats
      				</span>
                    <input type="text" class="form-control" v-model="repeats">
                </div>
            </div>
            <div class="col-lg-2">
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
            name: '',
            weight: 0,
            repeats: 8
        }
    },
    created: function() {

    },
    methods: {
        addMovement: function() {
            this.$emit('addMovement', {
                'name': this.$data.name,
                'weight': this.$data.weight,
                'repeats': this.$data.repeats
            })
            this.$data.name = '';
            this.$data.weight = 0;
            this.$data.repeats = 8
        }

    }

}
</script>
<style scoped>
.head-space {
    margin-top: 0.5em;
}
</style>
