<template>
<div>
<md-button class="md-fab md-raised md-accent" v-on:click.native="openReturnConfirm()">
	<md-icon>reply</md-icon>
</md-button>
    <md-table-card>
        <md-toolbar>
            <h1 class="md-title">{{$store.state.plan.planName}}的训练课</h1>
        </md-toolbar>
        <md-table>
            <md-table-header>
                <md-table-row>
                    <md-table-head style="width:25%">训练课名称</md-table-head>
                    <md-table-head>训练肌群</md-table-head>
                    <md-table-head>训练时间</md-table-head>
                    <md-table-head>重复次数</md-table-head>
                    <md-table-head style="width:20%"></md-table-head>
                </md-table-row>
            </md-table-header>
            <md-table-body>
                <md-table-row v-for="(row, rowIndex) in sessions" :key="rowIndex" :md-item="row">
                    <md-table-cell>
                        {{ row.name }}
                    </md-table-cell>
                    <md-table-cell>
                        {{ row.muscle }}
                    </md-table-cell>
                    <md-table-cell>
                        {{ timeMapping[row.weekly] }}
                    </md-table-cell>
                    <md-table-cell md-numeric>
                        {{ row.repeats }}
                    </md-table-cell>
                    <md-table-cell >
                        <md-button class="md-fab md-primary" v-on:click.native="editSesssions(rowIndex)">
                            <md-icon>edit</md-icon>
                        </md-button>
                        <md-button class="md-fab md-accent" v-on:click.native="removeSession(row)" v-if="$store.state.plan.planEditing">
                            <md-icon>remove circle outline</md-icon>
                        </md-button>
                    </md-table-cell>

                </md-table-row>
            </md-table-body>
        </md-table>
    </md-table-card>
    <hr>
    <md-list v-if="$store.state.plan.planEditing">
        <md-list-item>
            <md-input-container>
                <label>训练课名称</label>
                <md-input v-model="activeSession.name"></md-input>
            </md-input-container>
        </md-list-item>
        <md-list-item>
            <md-input-container>
                <label>目标肌群</label>
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
        <md-list-item>
            <md-button class="md-raised" v-on:click.native="addSessions()">添加并开始编辑训练课</md-button>
            <md-button class="md-raised md-primary" v-on:click.native="addSessions()">创建计划</md-button>
        </md-list-item>
    </md-list>
    <md-dialog-confirm md-title="" md-content="未保存的计划，返回会丢失现有数据，确认放弃么？" md-ok-text="确认放弃" md-cancel-text="返回继续" @close="onClose" ref="confirmReturn">
    </md-dialog-confirm>
</div>

</template>
<script>
export default {
	name:'plan-sessions',
	data(){
		return {
			activeSession:{},
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
			this.$store.commit('pushSession', _.clone(this.activeSession));
            this.activeSession = {}
            this.editSesssions(this.$store.state.plan.sessions.length - 1)
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
        }

	},
    computed:{
        sessions: function(){
            return this.$store.state.plan.sessions;
        }
    }
}
</script>