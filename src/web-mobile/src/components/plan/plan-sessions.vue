<template>
<div>
    <md-toolbar class="md-transparent">
        <md-button class="md-icon-button md-dense md-accent" v-on:click.native="openReturnConfirm()">
            <md-icon>reply</md-icon>
        </md-button>
        <span class="md-title" style="flex:1">{{$store.state.plan.planName}}的训练课</span>
        <md-button class="md-icon-button md-primary" v-on:click.native="startEditing()" v-if="!$store.state.plan.planEditing">
            <md-icon>edit</md-icon>
        </md-button>
        <md-button class="md-icon-button md-primary" v-on:click.native="saveEditing()" v-if="$store.state.plan.planEditing">
            <md-icon>save</md-icon>
        </md-button>
    </md-toolbar>
    <div style="overflow-y: auto; max-height: 25em;">
        <md-list md-triple-line>
            <md-list-item v-for="(row, rowIndex) in sessions">
                <div class="md-list-text-container" v-on:click="editSesssions(rowIndex)">
                    <span>{{ row.name }}</span>
                    <span>{{ row.targetMuscle }}</span>
                    <p>{{ timeMapping[row.weekly] }}</p>
                </div>
                <md-button class="md-icon-button md-accent md-list-action" v-if="$store.state.plan.planEditing" v-on:click.native="removeSession(row)">
                    <md-icon>delete</md-icon>
                </md-button>
                <md-divider></md-divider>
            </md-list-item>
        </md-list>
    </div>
    <md-button class="md-raised" v-on:click.native="showNewSessionModal()" v-if="$store.state.plan.planEditing">为课程添加新训练课</md-button>
    <md-dialog ref="newSessionModal">
        <md-dialog-content>
            <md-list>
                <md-list-item>
                    <md-input-container>
                        <label>训练课名称</label>
                        <md-input v-model="activeSession.name"></md-input>
                    </md-input-container>
                </md-list-item>
                <md-list-item>
                    <md-input-container>
                        <label>训练目标</label>
                        <md-select name="muscle" id="muscle" v-model="activeSession.muscle" multiple>
                            <md-option value="肩部">肩部</md-option>
                            <md-option value="胸部">胸部</md-option>
                            <md-option value="背部">背部</md-option>
                            <md-option value="手臂">手臂</md-option>
                            <md-option value="核心">核心</md-option>
                            <md-option value="腿部">腿部</md-option>
                            <md-option value="体能">体能(HIIT)</md-option>
                        </md-select>
                    </md-input-container>
                </md-list-item>
                <md-list-item>
                    <md-input-container>
                        <label>训练时间</label>
                        <md-select name="weekly" id="weekly" v-model="activeSession.weekly">
                            <md-option value="Monday">周一</md-option>
                            <md-option value="Tuesday">周二</md-option>
                            <md-option value="Wednesday">周三</md-option>
                            <md-option value="Thursday">周四</md-option>
                            <md-option value="Friday">周五</md-option>
                            <md-option value="Saturday">周六</md-option>
                            <md-option value="Sunday">周日</md-option>
                        </md-select>
                    </md-input-container>
                </md-list-item>
                <md-list-item>
                    <md-input-container>
                        <label>重复次数</label>
                        <md-input type="number" v-model="activeSession.repeats"></md-input>
                    </md-input-container>
                </md-list-item>
            </md-list>
            <md-dialog-actions>
                <md-button class="md-primary" v-on:click.native="closeDialog()">取消</md-button>
                <md-button class="md-primary" v-on:click.native="addSessions()">添加</md-button>
            </md-dialog-actions>
        </md-dialog-content>
    </md-dialog>
    <md-dialog-confirm md-title="" md-content="未保存的计划，返回会丢失现有数据，确认放弃么？" md-ok-text="确认放弃" md-cancel-text="返回继续" @close="onClose" ref="confirmReturn">
    </md-dialog-confirm>
</div>

</template>
<script>
export default {
	name:'plan-sessions',
	data(){
		return {
			activeSession:{
				muscle:[],
                weekly:''
			},
			timeMapping:{
				'Monday': '周一',
				'Tuesday': '周二',
				'Wednesday': '周三',
				'Thursday': '周四',
				'Friday': '周五',
				'Saturday': '周六',
				'Sunday': '周日'
			},
            editing: true
		}
	},
	methods:{
		addSessions: function() {
            this.activeSession.workouts = [];
            this.activeSession.targetMuscle = _.join(this.activeSession.muscle,',')
			this.$store.commit('pushSession', _.clone(this.activeSession));
            this.activeSession = {
                muscle :[]
            }
            this.closeDialog()
            //this.editSesssions(this.$store.state.plan.sessions.length - 1)
        },
        removeSession(row){
    		this.$store.commit('removeSession', row);
        },
        editSesssions(rowIndex){
            this.$store.commit('editSession', rowIndex);
    		window.location.href = "#/working/plan/per-session/" + rowIndex;
        },
        openReturnConfirm(){
        	if(this.$store.state.plan.planEditing){
            	this.$refs['confirmReturn'].open();
        	}else{
        		this.onClose('ok');
        	}
        },
        onClose(type){
            if(type == 'ok'){
                this.$store.commit('reset')
                this.$router.replace('/working/plan')
            }
        },
        startEditing(){
            this.$store.commit('editingPlan')
        },
        saveEditing(){
            this.$store.commit('readOnlyPlan')
        },
        showNewSessionModal(){
            this.$refs['newSessionModal'].open()
        },
        closeDialog(){
            this.$refs['newSessionModal'].close()
        }

	},
    computed:{
        sessions: function(){
            return this.$store.state.plan.sessions;
        }
    }
}
</script>