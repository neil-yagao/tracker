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
                <button class="btn btn-success">Add</button>
            </div>
        </div>
        <template v-for="(movement, index) in filtedMovement">
            <movement :movement="movement" :selected="selected" :editing="false" @selectingMovement="selecting"></movement>
        </template>
    </div>
</template>
<script>
import Movement from './Movement.vue'
export default {
    name: 'movement-list',
    data() {
        return {
            movements: [{
                id: 1,
                name: 'Bench Press',
                description: "Lie back on a flat bench.  Using a medium width grip (a grip that creates a 90-degree angle in the middle of the movement between the forearms and the upper arms), lift the bar from the rack and hold it straight over you with your arms locked.  This will be your starting position.",
                targetMuscle: 'Chest; UpperBody'
            }, {
                id: 2,
                name: 'Deadlift',
                description: "Begin with a bar loaded on the ground. Approach the bar so that the bar intersects the middle of the feet. The feet should be set very wide, near the collars. Bend at the hips to grip the bar. The arms should be directly below the shoulders, inside the legs, and you can use a pronated grip, a mixed grip, or hook grip. Relax the shoulders, which in effect lengthens your arms",
                targetMuscle: 'Back; Legs'
            }],
            selected: -1,
            nameKey: ''
        }
    },
    created: function() {
        console.info("movement list created")
    },
    methods: {
        selecting: function(index) {
            if (this.$data.selected == index) {
                this.$data.selected = -1
            } else {
                this.$data.selected = index
            }
        }
    },
    computed: {
        filtedMovement: function() {
            return _.filter(this.$data.movements, (m) => {
                return m.name.indexOf(this.$data.nameKey) > -1
            })
        }
    },
    components: {
        Movement
    }
}
</script>
<style scoped>
</style>
