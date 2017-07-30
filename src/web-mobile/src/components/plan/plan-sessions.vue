<template>
    <div>
        <md-toolbar class="md-transparent">
            <md-button class="md-icon-button md-dense md-accent" v-on:click.native="onClose('ok')">
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
        <div style="overflow-y: auto; max-height: 65vh">
            <md-list md-triple-line>
                <md-list-item v-for="(row, rowIndex) in sessions">
                    <div class="md-list-text-container" v-on:click="editSesssion(rowIndex)">
                        <span>{{ row.name }}</span>
                        <span>主要训练肌群：{{ row.targetMuscle }}</span>
                        <p>训练时间：每{{ timeMapping[row.weekly] }}</p>
                    </div>
                    <md-button class="md-icon-button md-accent md-list-action" v-if="$store.state.plan.planEditing" v-on:click.native="removeSession(row)">
                        <md-icon>delete</md-icon>
                    </md-button>
                    <md-divider></md-divider>
                </md-list-item>
            </md-list>
        </div>
        <md-button class="md-raised " v-on:click.native="showNewSessionModal()" v-if="$store.state.plan.planEditing">添加新训练课</md-button>
        <md-button class="md-raised" v-on:click.native="applyPlan()" v-if="!$store.state.plan.planEditing">应用计划</md-button>
        <md-dialog-prompt md-title="选择开始日期" md-ok-text="确定" md-cancel-text="取消" v-model="startDate" @close="confirmStart" ref="startDate">
        </md-dialog-prompt>
        <md-dialog ref="newSessionModal">
            <md-dialog-content>
                <md-list>
                    <md-list-item>
                        <md-input-container :class="showError && !activeSession.name?'md-input-invalid':''">
                            <label>训练课名称</label>
                            <md-input v-model="activeSession.name" required></md-input>
                        </md-input-container>
                    </md-list-item>
                    <md-list-item>
                        <md-input-container :class="showError && activeSession.muscle.length == 0?'md-input-invalid':''">
                            <label>训练目标</label>
                            <md-select name="muscle" id="muscle" v-model="activeSession.muscle" multiple required>
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
                        <md-input-container :class="showError && !activeSession.weekly?'md-input-invalid':''">
                            <label>训练时间</label>
                            <md-select name="weekly" id="weekly" v-model="activeSession.weekly" required>
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
                        <md-input-container :class="showError && !activeSession.repeat?'md-input-invalid':''">
                            <label>重复次数</label>
                            <md-input type="number" v-model="activeSession.repeat" required></md-input>
                        </md-input-container>
                    </md-list-item>
                </md-list>
                <md-dialog-actions>
                    <md-button class="md-primary" v-on:click.native="closeDialog()">取消</md-button>
                    <md-button class="md-primary" v-on:click.native="addSession()">添加</md-button>
                </md-dialog-actions>
            </md-dialog-content>
        </md-dialog>
        <loading-modal ref="loadingModal"></loading-modal>
    </div>
</template>
<script>
import moment from 'moment';
import LoadingModal from '../general/loading.vue'

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
			startDate:moment().format("YYYY-MM-DD"),
			showError:false
		}
	},
	methods:{
		addSession: function() {
			console.info("add session")
			if(!this.activeSession.name || !this.activeSession.weekly || !this.activeSession.repeat || this.activeSession.muscle.length == 0){
				this.showError = true;
				return
			}
            this.activeSession.workouts = [];
            this.activeSession.targetMuscle = _.join(this.activeSession.muscle,',')
			this.$store.commit('pushSession', _.clone(this.activeSession));
			this.activeSession.repeat = Number(this.activeSession.repeat)
			this.$http.post('/plan/' + this.planId + "/sessions", this.activeSession).then((res) =>{
				this.loadPlan();
				this.$store.commit('editingPlan')

			})
            this.activeSession = {
                muscle :[],
                name:'',
                weekly:'',
                repeat:''
            }
            this.closeDialog()

        },
        removeSession(row){
    		this.$store.commit('removeSession', row);
    		this.$http.delete('/plan/' + this.planId + "/session/" + row.id)
        },
        editSesssion(rowIndex){
            this.$store.commit('editSession', rowIndex);
    		window.location.href = "#/working/plan/"+ this.planId +"/session/" + rowIndex;
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
        },
        applyPlan(){
        	this.$refs['startDate'].open()
        },
        confirmStart(type){
        	console.info(type)
        	if(type == 'ok'){
	        	this.$refs['loadingModal'].open();
	        	var confirmInfo = {
	        		user:this.$store.state.user,
	        		date:this.startDate
	        	}
	        	this.$http.post('/plan/' + this.planId, confirmInfo).then((res)=>{
	    			this.$refs['loadingModal'].close();
	        	})
        	}
        },
        loadPlan(){
        	this.$http.get('/plan/' + this.planId).then((res) =>{
				var planDetail = res.body.data;
                console.info("load plan");
                console.info(planDetail);
				this.$store.commit('resetPlan', planDetail);
			})

        }
	},
	components:{
		'loading-modal': LoadingModal
	},
    computed:{
        sessions: function(){
            return this.$store.state.plan.sessions;
        },
        planId: function(){
        	return this.$route.params.id
        }
    },
    mounted:function(){
    	console.info('loaded!')
    	if(this.$store.state.plan.sessions.length == 0){
    		this.loadPlan()
    	}
    	
    }
}
</script>