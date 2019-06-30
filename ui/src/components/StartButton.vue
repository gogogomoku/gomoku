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
                this.$parent.startGame();
            } else {
                this.$parent.restartGame();
            }
        }
    }
}
</script>

<style scoped>
    button {
        padding: 20px;
        font-size: 2vmin;
        border-radius: 3px;
    }
    button :hover {
        background-color: #9CA;
        border-radius: 300px;
    }
</style>
