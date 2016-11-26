<template>
    <div class="panel panel-default">
        <div class="panel-heading" @click="selectedMovement(movement.id)">{{movement.name}}</div>
        <div class="panel-body" v-show="selected == movement.id">
            <div class="thumbnail" v-show="!editing">
                <img src="http://www.bodybuilding.com/exercises/exerciseImages/sequences/722/Male/l/722_1.jpg">
                <div class="caption">
                    <h3>{{movement.name}}</h3>
                    <p>{{movement.description}}</p>
                    <p>
                        <b>Target Muscle:</b>
                        <ul>
                            <li v-for="muscle in eachMuscleGroup" class="active"><a href="#">{{muscle}}</a></li>
                        </ul>
                    </p>
                    <p>
                        <button class="btn btn-warning" role="button" @click="editing=true">Edit</button>
                    </p>
                </div>
            </div>
            <div class="thumbnail" v-show='editing'>
                <div class="caption">
                    <p>
                        <div class="row">
                            <div class="col-md-4">
                                <input v-model="mv.name" class="form-control"></input>
                            </div>
                            <div class="col-md-3 col-md-offset-3">
                                <div class="input-group" style="margin-bottom:1em">
                                    <input v-model="mv.src" class="form-control">
                                    <span class="input-group-btn">
                                        <button class="btn btn-default" type="button">Upload</button>
                                    </span>
                                </div>
                            </div>
                        </div>
                    </p>
                    <p>
                        <textarea class="form-control" v-model="mv.description"></textarea>
                    </p>
                    <p>
                        <b>Target Muscle(split by";"):</b>
                        <input v-model="mv.targetMuscle" class="form-control"></input>
                    </p>
                    <p>
                        <button class="btn btn-success" role="button" @click="editing=false">Save</button>
                    </p>
                </div>
            </div>
        </div>
</template>
<script>
export default {
    name: 'movement',
    data() {
        return {
            mv: this.movement
        }
    },
    props: ['movement', 'selected', 'editing'],
    methods: {
        selectedMovement: function(id) {
            this.$emit('selectingMovement', id)
        }
    },
    computed: {
        eachMuscleGroup: function() {
            return this.$data.mv.targetMuscle.split(';')
        }
    }
}
</script>
<style scoped>
.panel {
    margin-bottom: 5px
}
</style>
