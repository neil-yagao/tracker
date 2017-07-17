<template>
        <div class="jumbotron" >
            <h1>我们应该怎么称呼您?</h1>
            <p>
                <div class="row padding-top">
                    <div class="col-xs-8">
                        <input type="text" v-model="username" class="form-control input-lg" @keyup.enter="login()">
                    </div>
                    <div class="col-xs-2 col-xs-offset-1">
                        <a class="btn btn-info btn-lg" role="button" @click="login()">开始举铁!</a>
                    </div>
                </div>
            </p>
        </div>

</template>
<script>
import md5 from 'md5'
export default {
    name: 'login',
    data() {
        return {
            username: ''
        }
    },
    methods: {
        login() {
            var md5Encode = md5(this.username);
            console.info(md5Encode);
            /*this.$http.post('login',{
            	'username': this.username,
            	'usr':md5Encode
            }).then((res)=>{
           		this.$store.state.username = this.username;
            	window.location.href = '#/working';
            })*/
            this.$store.commit('setUser',{
                username:this.username,
                userIdentity: md5Encode
            })
			window.location.href = '#/working';
        }
    }
}
</script>
<style scoped>
.jumbotron {
    padding-left: 3%;
    position: absolute;
    width: 100%;
    top:40%;
    left:1%;
}

.vcenter {
    margin-top: 50%;
}
</style>
