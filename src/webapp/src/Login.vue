<template>
    <div class="text-center" style="padding:50px 0">
        <div class="logo">登陆</div>
        <!-- Main Form -->
        <div class="login-form-1">
            <div id="login-form" class="text-left">
                <div class="login-form-main-message"></div>
                <div class="main-login-form">
                    <div class="login-group">
                        <div class="form-group">
                            <label for="lg_username" class="sr-only">用户名</label>
                            <input type="text" class="form-control" v-model="user" placeholder="用户名" @keyup.enter="login">
                        </div>
                    </div>
                    <button class="login-button" @click="login"><i class="fa fa-chevron-right"></i></button>
                </div>
            </div>
        </div>
        <!-- end:Main Form -->
    </div>
</template>
<script>
window.mobilecheck = function() {
    var check = false;
    (function(a) {
        if (/(android|bb\d+|meego).+mobile|avantgo|bada\/|blackberry|blazer|compal|elaine|fennec|hiptop|iemobile|ip(hone|od)|iris|kindle|lge |maemo|midp|mmp|mobile.+firefox|netfront|opera m(ob|in)i|palm( os)?|phone|p(ixi|re)\/|plucker|pocket|psp|series(4|6)0|symbian|treo|up\.(browser|link)|vodafone|wap|windows ce|xda|xiino/i.test(a) || /1207|6310|6590|3gso|4thp|50[1-6]i|770s|802s|a wa|abac|ac(er|oo|s\-)|ai(ko|rn)|al(av|ca|co)|amoi|an(ex|ny|yw)|aptu|ar(ch|go)|as(te|us)|attw|au(di|\-m|r |s )|avan|be(ck|ll|nq)|bi(lb|rd)|bl(ac|az)|br(e|v)w|bumb|bw\-(n|u)|c55\/|capi|ccwa|cdm\-|cell|chtm|cldc|cmd\-|co(mp|nd)|craw|da(it|ll|ng)|dbte|dc\-s|devi|dica|dmob|do(c|p)o|ds(12|\-d)|el(49|ai)|em(l2|ul)|er(ic|k0)|esl8|ez([4-7]0|os|wa|ze)|fetc|fly(\-|_)|g1 u|g560|gene|gf\-5|g\-mo|go(\.w|od)|gr(ad|un)|haie|hcit|hd\-(m|p|t)|hei\-|hi(pt|ta)|hp( i|ip)|hs\-c|ht(c(\-| |_|a|g|p|s|t)|tp)|hu(aw|tc)|i\-(20|go|ma)|i230|iac( |\-|\/)|ibro|idea|ig01|ikom|im1k|inno|ipaq|iris|ja(t|v)a|jbro|jemu|jigs|kddi|keji|kgt( |\/)|klon|kpt |kwc\-|kyo(c|k)|le(no|xi)|lg( g|\/(k|l|u)|50|54|\-[a-w])|libw|lynx|m1\-w|m3ga|m50\/|ma(te|ui|xo)|mc(01|21|ca)|m\-cr|me(rc|ri)|mi(o8|oa|ts)|mmef|mo(01|02|bi|de|do|t(\-| |o|v)|zz)|mt(50|p1|v )|mwbp|mywa|n10[0-2]|n20[2-3]|n30(0|2)|n50(0|2|5)|n7(0(0|1)|10)|ne((c|m)\-|on|tf|wf|wg|wt)|nok(6|i)|nzph|o2im|op(ti|wv)|oran|owg1|p800|pan(a|d|t)|pdxg|pg(13|\-([1-8]|c))|phil|pire|pl(ay|uc)|pn\-2|po(ck|rt|se)|prox|psio|pt\-g|qa\-a|qc(07|12|21|32|60|\-[2-7]|i\-)|qtek|r380|r600|raks|rim9|ro(ve|zo)|s55\/|sa(ge|ma|mm|ms|ny|va)|sc(01|h\-|oo|p\-)|sdk\/|se(c(\-|0|1)|47|mc|nd|ri)|sgh\-|shar|sie(\-|m)|sk\-0|sl(45|id)|sm(al|ar|b3|it|t5)|so(ft|ny)|sp(01|h\-|v\-|v )|sy(01|mb)|t2(18|50)|t6(00|10|18)|ta(gt|lk)|tcl\-|tdg\-|tel(i|m)|tim\-|t\-mo|to(pl|sh)|ts(70|m\-|m3|m5)|tx\-9|up(\.b|g1|si)|utst|v400|v750|veri|vi(rg|te)|vk(40|5[0-3]|\-v)|vm40|voda|vulc|vx(52|53|60|61|70|80|81|83|85|98)|w3c(\-| )|webc|whit|wi(g |nc|nw)|wmlb|wonu|x700|yas\-|your|zeto|zte\-/i.test(a.substr(0, 4))) check = true;
    })(navigator.userAgent || navigator.vendor || window.opera);
    return check;
};
var storage;
var fail;
var uid;
try {

    uid = new Date;
    (storage = window.localStorage).setItem(uid, uid);
    fail = storage.getItem(uid) != uid;
    storage.removeItem(uid);
    fail && (storage = false);
} catch (exception) {}
export default {
    name: 'login',
    data() {
        return {
            user: ''
        }
    },
    methods: {
        login: function() {
            if (storage) {
                var userIdentity = md5(this.$data.user);
                console.info("username:" + this.$data.user)
                window.localStorage.setItem('user', userIdentity);

                this.$http.post('/login', {
                    'username': this.$data.user,
                    'useridentity': userIdentity
                }).then((response) => {
                    if (response.body.success) {
                        if (window.mobilecheck()) {
                            console.info("using mobile device!")
                        } else {
                            window.location.href = "#/lift/workouts"
                        }
                    }
                })
            } else {
                console.warn("enable to use this browser to login!")
            }

        }
    }
}
</script>
<style scoped>
a {
    color: #aaaaaa;
    transition: all ease-in-out 200ms;
}

a:hover {
    color: #333333;
    text-decoration: none;
}


/*=== 4. Main Form ===*/

.login-form-1 {
    max-width: 300px;
    border-radius: 5px;
    display: inline-block;
}

.main-login-form {
    position: relative;
}

.login-form-1 .form-control {
    border: 0;
    box-shadow: 0 0 0;
    border-radius: 0;
    background: transparent;
    color: #555555;
    padding: 7px 0;
    font-weight: bold;
    height: auto;
}

.login-form-1 .form-control::-webkit-input-placeholder {
    color: #999999;
}

.login-form-1 .form-control:-moz-placeholder,
.login-form-1 .form-control::-moz-placeholder,
.login-form-1 .form-control:-ms-input-placeholder {
    color: #999999;
}

.login-form-1 .form-group {
    margin-bottom: 0;
    border-bottom: 2px solid #efefef;
    padding-right: 20px;
    position: relative;
}

.login-form-1 .form-group:last-child {
    border-bottom: 0;
}

.login-group {
    background: #ffffff;
    color: #999999;
    border-radius: 8px;
    padding: 10px 20px;
}

.login-group-checkbox {
    padding: 5px 0;
}


/*=== 5. Login Button ===*/

.login-form-1 .login-button {
    position: absolute;
    right: -25px;
    top: 50%;
    background: #ffffff;
    color: #999999;
    padding: 11px 0;
    width: 50px;
    height: 50px;
    margin-top: -25px;
    border: 5px solid #efefef;
    border-radius: 50%;
    transition: all ease-in-out 500ms;
}

.login-form-1 .login-button:hover {
    color: #555555;
    transform: rotate(450deg);
}

.login-form-1 .login-button.clicked {
    color: #555555;
}

.login-form-1 .login-button.clicked:hover {
    transform: none;
}

.login-form-1 .login-button.clicked.success {
    color: #2ecc71;
}

.login-form-1 .login-button.clicked.error {
    color: #e74c3c;
}


/*=== 6. Form Invalid ===*/

label.form-invalid {
    position: absolute;
    top: 0;
    right: 0;
    z-index: 5;
    display: block;
    margin-top: -25px;
    padding: 7px 9px;
    background: #777777;
    color: #ffffff;
    border-radius: 5px;
    font-weight: bold;
    font-size: 11px;
}

label.form-invalid:after {
    top: 100%;
    right: 10px;
    border: solid transparent;
    content: " ";
    height: 0;
    width: 0;
    position: absolute;
    pointer-events: none;
    border-color: transparent;
    border-top-color: #777777;
    border-width: 6px;
}


/*=== 7. Form - Main Message ===*/

.login-form-main-message {
    background: #ffffff;
    color: #999999;
    border-left: 3px solid transparent;
    border-radius: 3px;
    margin-bottom: 8px;
    font-weight: bold;
    height: 0;
    padding: 0 20px 0 17px;
    opacity: 0;
    transition: all ease-in-out 200ms;
}

.login-form-main-message.show {
    height: auto;
    opacity: 1;
    padding: 10px 20px 10px 17px;
}

.login-form-main-message.success {
    border-left-color: #2ecc71;
}

.login-form-main-message.error {
    border-left-color: #e74c3c;
}

.logo {
    padding: 15px 0;
    font-size: 25px;
    color: #aaaaaa;
    font-weight: bold;
}
</style>
