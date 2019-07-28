<template>
    <button v-on:click="buttonAction">{{ buttonMessage }}</button>
</template>

<script>
export default {
    name: "StartButton",
    props: ["buttonMessage", "gameStatus"],
    methods: {
        buttonAction: function () {
            if (this.gameStatus == 0){
                this.$parent.$parent.$parent.startGame();
            } else {
                this.$parent.$parent.$parent.restartGame();
            }
        }
    }
}
</script>

<style scoped>
    button {
        width: 74%;
        padding: 20px;
        font-size: 2vmin;
        border-radius: 3px;
        background-color: #BBE;
    }
    button:hover {
        background-color: #CCF;
    }
</style>
