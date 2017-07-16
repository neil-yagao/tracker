<template>
<div>
    <md-input-container v-on:click.native="showDialog()">
        <label>训练动作</label>
        <md-input v-model="name"></md-input>
    </md-input-container>
    <md-dialog ref="movementSelection">
        <md-dialog-title>选择需要的动作</md-dialog-title>
        <md-dialog-content>
            <md-toolbar md-theme="white">
                <md-input-container>
                <md-icon>search</md-icon>
                <md-input class="md-input" v-model="criteria" id="criteria"></md-input>
                </md-input-container>
            </md-toolbar>
            <md-list>
                <md-list-item>
                    <span>News</span>
                    <md-list-expand>
                        <md-list>
                            <md-list-item class="md-inset">World</md-list-item>
                            <md-list-item class="md-inset">Americas</md-list-item>
                            <md-list-item class="md-inset">Europe</md-list-item>
                        </md-list>
                    </md-list-expand>
                </md-list-item>
                <md-list-item>
                    <span>Games</span>
                    <md-list-expand>
                        <md-list>
                            <md-list-item class="md-inset">Console</md-list-item>
                            <md-list-item class="md-inset">PC</md-list-item>
                            <md-list-item class="md-inset">Phone</md-list-item>
                        </md-list>
                    </md-list-expand>
                </md-list-item>
            </md-list>
        </md-dialog-content>
        <md-dialog-actions>
            <md-button class="md-primary" @click="closeDialog()">Cancel</md-button>
            <md-button class="md-primary" @click="closeDialog()">Ok</md-button>
        </md-dialog-actions>
    </md-dialog>
</div>


</template>
<script type="text/javascript">
export default {
    name:'movement-selection',
    data(){
        return {
            name:''
        }
    },
    methods:{
        showDialog(){
            this.$refs['movementSelection'].open();
        },
        closeDialog(){
            this.$refs['movementSelection'].close();
        }
    },

} 
</script>
<style>
.md-dialog{width:90%;height: 70%}
#criteria {
    margin-left: 10px
}
</style>