<template>
<div>
    <md-list>
        <md-list-item>
            <md-input-container>
                <label>课程名称</label>
                <md-input v-model="$store.state.newSession.name"></md-input>
            </md-input-container>
        </md-list-item>
        <md-list-item>
            <md-input-container>
                <label>所属计划</label>
                <md-input v-model="$store.state.newSession.plan"></md-input>
            </md-input-container>
        </md-list-item>
        <md-list-item>
            <md-input-container>
                <label>目标肌群</label>
                <md-select name="muscle" id="muscle" v-model="$store.state.newSession.muscle">
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
                <label>每周重复</label>
                <md-select name="repeat" id="repeat" v-model="$store.state.newSession.repeat">
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
    </md-list>
</div>

</template>
<script>
export default {
	name:'session',
	data(){
		return {
		}
	},
	methods:{
		
	}
}
	
</script>