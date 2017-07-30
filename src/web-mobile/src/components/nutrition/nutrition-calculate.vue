<template>
<div>
    <md-card md-with-hover>
        <md-card-header>
            <div class="md-title">您的营养分解</div>
            <div class="md-subtitle">您需要的卡路里总量为：{{calorie}}</div>
        </md-card-header>
        <md-card-content>
        <md-list>
            <md-list-item>
                <md-button-toggle md-single>
                    <md-button v-on:click.native="extra = 1">增肌饮食</md-button>
                    <md-button v-on:click.native="extra = -1">减脂饮食</md-button>
                    <md-button v-on:click.native="extra = 0">保持体重</md-button>
                </md-button-toggle>
       
            </md-list-item>
            <md-list-item>
                <md-icon>whatshot</md-icon>
                <span>蛋白质:{{protein}}</span>
           <!--      <md-list-expand>
                    <md-list>
                        <md-list-item class="md-inset">World</md-list-item>
                        <md-list-item class="md-inset">Americas</md-list-item>
                        <md-list-item class="md-inset">Europe</md-list-item>
                    </md-list>
                </md-list-expand> -->
            </md-list-item>
            <md-list-item>
                <md-icon>whatshot</md-icon>
                <span>碳水化合物:{{carbs}}</span>
           <!--      <md-list-expand>
                    <md-list>
                        <md-list-item class="md-inset">World</md-list-item>
                        <md-list-item class="md-inset">Americas</md-list-item>
                        <md-list-item class="md-inset">Europe</md-list-item>
                    </md-list>
                </md-list-expand> -->
            </md-list-item>
            <md-list-item>
                <md-icon>whatshot</md-icon>
                <span>脂肪：{{fat}}</span>
           <!--      <md-list-expand>
                    <md-list>
                        <md-list-item class="md-inset">World</md-list-item>
                        <md-list-item class="md-inset">Americas</md-list-item>
                        <md-list-item class="md-inset">Europe</md-list-item>
                    </md-list>
                </md-list-expand> -->
            </md-list-item>

            </md-list>
            <md-card-actions>
            <md-button class="md-raised" v-on:click.native="goToCollect()">
            		<md-icon>edit</md-icon>
            		重新填写
            	</md-button>
        	</md-card-actions>
        </md-card-content>
    </md-card>
</div>
</template>
<script type="text/javascript">
export default {
	name:'nutrition-track',
	data(){
		return {
			extra: 0,
			physique:{}
		}
	},
	mounted:function(){
		this.$http.get("/physique/" + this.$store.state.user.id).then((res)=>{
			if(res.body.data){
				console.info(res.body.data)
				this.physique = res.body.data
			}else{
				this.$router.replace('/working/nutrition/collect')
			}
		})
	},
	methods:{
		goToCollect:function(){
			this.$router.replace('/working/nutrition/collect')
		}
	},
	computed:{
		rmc: function(){
			var base = 10 * this.physique.weight + 6.25 * this.physique.height - 5 * this.physique.age;
			if(this.physique.gender == 1){
				//male
				return  base + 5
			}else if(this.physique.gender == 2){
				return base - 161
			}
		},
		calorie: function(){
			if(this.extra > 0){
				return this.rmc + 300
			}else if(this.extra < 0){
				return this.rmc - 500
			}
			return this.rmc
		},
		protein:function(){
			if(this.extra < 0){
				return Math.floor(this.calorie * 0.5 / 4)
			}else {
				return Math.floor(this.calorie * 0.4 / 4)
			}
		},
		carbs:function(){
			if(this.extra < 0){
				return Math.floor(this.calorie * 0.3 / 4)
			}else {
				return Math.floor(this.calorie * 0.4 / 4)
			}
		},
		fat:function(){
			return Math.floor(this.calorie * 0.2 / 9)
		}

	}
}
</script>	