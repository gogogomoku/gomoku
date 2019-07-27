<template>
    <div id="gameContainer">
        <Goban
            v-bind:size="size"
            v-bind:tab="tab"
            v-bind:turn="turn"
            v-bind:currentPlayer="currentPlayer"
        />
        <Scoreboard
            v-bind:turn="turn"
            v-bind:currentPlayer="currentPlayer"
            v-bind:playerInfo="playerInfo"
        />
    </div>
</template>

<script>
import Goban from './Goban.vue'
import Scoreboard from './scoreBoard/Scoreboard.vue'

export default {
    name: 'GameContainer',
    components: {
        Goban,
        Scoreboard,
    },
    props: ["size", "tab", "turn", "currentPlayer", "playerInfo"],
    methods: {
    }
}
</script>

<style scoped>
    #gameContainer {
        background-color: #258;

        margin: 5px auto;
        display: flex;
        flex-direction: row;
        flex-wrap: wrap;
        justify-content: center;
        width: 60%;
        min-width: 600px;
    }
</style>
