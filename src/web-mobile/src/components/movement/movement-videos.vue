<template>
    <div class="scrolls">
        <md-card style='display:inline-block'>
            <md-card-media-actions>
                <md-card-media v-show="capture">
                 
                </md-card-media>
                <md-card-actions>
                    <md-button class="md-icon-button" >
                        <md-icon>favorite</md-icon>
                    </md-button>
                    <md-button class="md-icon-button">
                        <md-icon>bookmark</md-icon>
                    </md-button>
                    <md-button class="md-icon-button">
                        <md-icon>share</md-icon>
                    </md-button>
                </md-card-actions>
            </md-card-media-actions>
        </md-card>
      
    </div>	
</template>
<script>
export default {
	name:'movement-videos',
	data(){
		return {
			capture:true
		}
	},
	methods:{
		
	},
	mounted:function(){
		
	}
}
</script>
<style>
.scrolls { 
    overflow-x: scroll;
    overflow-y: hidden;
	white-space:nowrap
   } 
</style>