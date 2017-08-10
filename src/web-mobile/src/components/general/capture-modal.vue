<template>
	<md-dialog  ref="captureDialog">
	  <md-dialog-content>
	  	<md-card style='display:inline-block'>
            <md-card-media-actions>
                <md-card-media v-show="capture">
                    <video id="video"></video> 
                </md-card-media>
                <md-card-media v-show="!capture">
                    <img id="img"></img> 
                </md-card-media>
                <md-card-actions>
                    <md-button class="md-icon-button" v-on:click.native="captureGif()">
                        <md-icon>play_circle_outline</md-icon>
                    </md-button>
                    <md-button class="md-icon-button" v-on:click.native="clear()">
                        <md-icon>cancel</md-icon>
                    </md-button>
                </md-card-actions>
            </md-card-media-actions>
        </md-card>
        <md-progress class="md-accent" :md-progress="progress"></md-progress>
       </md-dialog-content>

	  <md-dialog-actions>
	    <md-button class="md-primary" @click="close()">取消</md-button>
	    <md-button class="md-primary" @click="upload()">上传</md-button>
	  </md-dialog-actions>
</md-dialog>
</template>
<script>
import gifshot from 'gifshot'

export default {
	name:'capture-modal',
	data(){
		return {
			capture:true,
			formData:'',
			progress:0,
			stream:''
		}
	},
	methods:{
		close(){
			this.clear()
			localStream.getVideoTracks()[0].stop();
			console.info(localStream.getVideoTracks())
			this.$refs['captureDialog'].close()
		},
		upload(){
			this.formData = new FormData()
			this.formData.set('gif',document.getElementById('img').src)
			this.$http.post('/movement/upload/'+ this.$store.state.user.id + '/' + this.movement, this.formData)
				.then(()=>{
					this.close()
				})
		},
		clear(){
			this.capture = true;
			document.getElementById('img').src = '';
			this.progress = 0;
			this.initVideo()
		},
		open(){
			this.$refs['captureDialog'].open()
		},
		initVideo(){
			var self = this;
	        navigator.getMedia = (navigator.getUserMedia ||
	            navigator.webkitGetUserMedia ||
	            navigator.mozGetUserMedia ||
	            navigator.msGetUserMedia);

	        navigator.getMedia({
	                video: true,
	                audio: false
	            },
	            function(stream) {
					var video = document.getElementById('video');
					if(!video){
						stream.getVideoTracks()[0].stop()
					}
	                if (navigator.mozGetUserMedia) {
	                    video.mozSrcObject = stream;
	                } else {
	                    var vendorURL = window.URL || window.webkitURL;
	                    video.src = vendorURL.createObjectURL(stream);
	                }
	                window.localStream = stream;
	                video.play();
	            },
	            function(err) {
	                console.log("An error occured! " + err);
	            }
	        );
		},
		captureGif(){
			var self = this;
			this.initVideo()
			var progressInterval = setInterval(()=>{
				this.progress += 2
				if(this.progress > 100){
					clearInterval(progressInterval)
				}
			}, 200);	
			gifshot.createGIF({
			    interval: 0.2,
			    numFrames: 50,
			    frameDuration: 10
			},function(obj) {
				if(!obj.error) {
					self.capture = false;
					gifshot.stopVideoStreaming();
					var image = obj.image,
					animatedImage = document.getElementById('img');
					animatedImage.src = image;
				}
			});
		}
	},
	props:['movement']
}
</script>