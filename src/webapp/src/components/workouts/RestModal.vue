<template>
<div id="rest-modal" class="modal fade" tabindex="-1" role="dialog">
    <div class="modal-dialog modal-sm" role="document">
        <div class="modal-content" :class="{'flash':remaining < 10}">
            <div class="modal-header">
                <button type="button" class="close" data-dismiss="modal" aria-label="Close"><span aria-hidden="true">&times;</span></button>
                <h4>休息倒计时</h4>
            </div>
            <p>
                <span v-show="false">{{flag}}</span>
                <h3><span v-show="remaining < 0">Passed:</span>{{Math.abs(remaining)}}<h3>
            </p>

        </div>
    </div>
</div>
</template>
<script>
export default {
    name: 'rest-modal',
    data() {
        return {
            remaining: this.rest,
            flag: false
        }
    },
    props: ['rest', 'startFlag'],
    computed: {
        flag: function() {
            console.info("start flag changed")
            this.$data.remaining = this.rest
            let self = this
            clearInterval(self.interval)
            self.interval = setInterval(function() {
                self.$data.remaining = self.$data.remaining - 1
            }, 1000)
            return !this.startFlag
        }
    }
}
</script>
<style>
@-webkit-keyframes flash {
    0% {
        background-color: Red;
        opacity: 1;
    }
    50% {
        background-color: Red;
        opacity: 0;
    }
    100% {
        background-color: Red;
    }
}

.flash {
    animation: flash 2s linear infinite;
}
</style>
