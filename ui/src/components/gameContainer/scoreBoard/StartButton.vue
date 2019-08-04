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
        padding: 9px 18px 12px;
        background-color: rgb(255, 173, 79);
        border: none;
        font-size: 16px;
        font-weight: bolder;
    }
    button:hover {
        background-color: rgb(255, 145, 1);
    }

    @media (min-width: 800px) {
        button {
            padding: 9px 18px 12px;
            font-size: 16px;
            width: 74%;
            border-radius: 3px;
        }
    }
</style>
