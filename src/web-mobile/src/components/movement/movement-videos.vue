<template>
    <div class="scrolls" v-if="img">
        <md-card style='display:inline-block'>
            <md-card-media-actions>
                <md-card-media >
                    <img :src="img"/>
                </md-card-media>
    <!--    this version not available         <md-card-actions>
                    <md-button class="md-icon-button" >
                        <md-icon>favorite</md-icon>
                    </md-button>
                    <md-button class="md-icon-button">
                        <md-icon>bookmark</md-icon>
                    </md-button>
                    <md-button class="md-icon-button">
                        <md-icon>share</md-icon>
                    </md-button>
                </md-card-actions> -->
            </md-card-media-actions>
        </md-card>
        <capture-modal ref="capture" :movement="movement" @close="loadImg()"></capture-modal>
      
    </div>	
</template>
<script>
import CaptureModal from '../general/capture-modal.vue'
export default {
	name:'movement-videos',
	data(){
		return {
            img:''
		}
	},
	methods:{
		open(){
            this.$refs['capture'].open()
        },
        loadImg(){
            this.$http.get('/movement/img/' + this.movement).then((res)=>{
                console.info(res.body.data)
                var location = res.body.data.location
                if(location){
                    this.img = '/static/'+location
                }
            })
        }
	},
    components:{
        'capture-modal':CaptureModal
    },
	mounted:function(){
		this.loadImg()
	},
    props:['movement']
}
</script>
<style>
.scrolls { 
    overflow-x: scroll;
    overflow-y: hidden;
	white-space:nowrap
   } 
</style>