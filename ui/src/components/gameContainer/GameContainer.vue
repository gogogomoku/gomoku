<template>
    <div id="gameContainer">
        <Goban
            v-bind:size="size"
            v-bind:tab="tab"
            v-bind:turn="turn"
            v-bind:currentPlayer="currentPlayer"
            v-bind:suggestedPosition="suggestedPosition"
        />
        <Scoreboard
            v-bind:turn="turn"
            v-bind:currentPlayer="currentPlayer"
            v-bind:playerInfo="playerInfo"
            v-bind:buttonMessage="buttonMessage"
            v-bind:gameStatus="gameStatus"
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
    props: [
        "size",
        "tab",
        "turn",
        "currentPlayer",
        "playerInfo",
        "suggestedPosition",
        "buttonMessage",
        "gameStatus",
    ],
    methods: {
    }
}
</script>

<style scoped>
    #gameContainer {
        margin: 0px auto;
        display: flex;
        flex-direction: row;
        flex-wrap: nowrap;
        justify-content: space-around;
        align-items: stretch;
        width: 60%;
        min-width: 800px;
    }
</style>
