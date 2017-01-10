<template>
    <div class="panel panel-default">
        <div class="panel-heading" @click="selectedMovement(movement.id)" v-show="movement.name">{{movement.name}}</div>
        <div class="panel-body" v-show="selected == movement.id || !movement.id ">
            <div class="thumbnail" v-show="!edited">
                <p>
                    <img src="http://www.bodybuilding.com/exercises/exerciseImages/sequences/722/Male/l/722_1.jpg" width="100" height="100">
                </p>
                <div class="caption">
                    <h3>{{movement.name}}</h3>
                    <p>{{movement.description}}</p>
                    <p>
                        <b>目标肌群:</b>
                        <ul>
                            <li v-for="muscle in eachMuscleGroup" class="active"><a href="#">{{muscle}}</a></li>
                        </ul>
                    </p>
                    <p>
                        <button class="btn btn-warning" role="button" @click="edited = true">修改</button>
                    </p>
                </div>
            </div>
            <div class="thumbnail" v-show="edited">
                <div class="caption">
                    <p>
                        <div class="row">
                            <div class="col-md-4">
                                <b>动作名称</b>
                            </div>
                        </div>
                        <div class="row">
                            <div class="col-md-4">
                                <input v-model="mv.name" class="form-control"></input>
                            </div>
                            <div class="col-md-4 col-md-offset-2">
                                <div class="input-group" style="margin-bottom:1em">
                                    <input v-model="mv.src" class="form-control" disabled="true">
                                    <span class="input-group-btn">
                                        <button class="btn btn-default" disabled="true" type="button">上传图片</button>
                                    </span>
                                </div>
                            </div>
                        </div>
                    </p>
                    <div class="row">
                        <div class="col-md-4">
                            <b>动作描述或步骤</b>
                        </div>
                    </div>
                    <p>
                        <textarea class="form-control" v-model="mv.description"></textarea>
                    </p>
                    <p>
                        <div class="row">
                            <div class="col-md-4">
                                <b>目标肌群(请用";"进行分割):</b>
                            </div>
                        </div>
                        <div class="row">
                            <div class="col-md-7">
                                <input v-model="mv.targetMuscle" class="form-control"></input>
                            </div>
                            <div class="col-md-3">
                                <div class="checkbox">
                                    <label>
                                        <input type="checkbox" v-model="mv.dividable">能否计算每边重量
                                    </label>
                                </div>
                            </div>
                        </div>
                    </p>
                    <p>
                        <button class="btn btn-success" role="button" @click="updateOrInsert">保存</button>
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
            mv: this.movement,
            edited: this.editing?this.editing: false
        }
    },
    props: ['movement', 'selected', 'editing'],
    methods: {
        selectedMovement: function(id) {
            this.$emit('selectingMovement', id)
        },
        updateOrInsert: function() {
            this.$data.edited = false;
            var movementId = this.$data.mv.id;
            var requestBody = JSON.stringify(this.$data.mv);
            if (movementId && movementId != -1) {
                //update
                this.$http.post('/movement/' + movementId, requestBody)
                    .then((respons) => {
                        console.info(respons)
                    })
            } else {
                this.$http.put('/movements', requestBody)
            }
            this.$emit('saveMovement', this.$data.mv)
        }
    },
    beforeUpdate: function() {
        this.$data.mv = this.movement;
        //this.$data.edited = this.editing
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
