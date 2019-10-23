<template>
    <div id="transLoading" class="show">
        <div class="loading-text">{{ text }}</div>
        <div class="spinner">
            <div class="bounce1"></div>
            <div class="bounce2"></div>
            <div class="bounce3"></div>
        </div>
    </div>
</template>
<script>
    export default {
        name: 'loading',
        props: {
            text: ""
        }
    }
</script>
<style scoped>
    #transLoading {
        position: fixed;
        top: 0;
        left: 0;
        width: 100%;
        height: 100%;
        background: rgba(255, 255, 255, 0.5);
        z-index: 9999;
        transition: .1s ease opacity;
        -webkit-transition: .1s ease opacity;
    }
    .loading-text {
        width: 150px;
        text-align: center;
        position: fixed;
        top: 40%;
        left: 50%;
        -webkit-transform: translate(-50%, -50%);
        transform: translate(-50%, -50%);
    }
</style>
