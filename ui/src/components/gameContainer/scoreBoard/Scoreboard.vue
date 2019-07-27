<template>
    <div id="scoreboard">
        <PlayerScoreBoard id="player1Sb" class="playerSb"
            v-bind:playerInfo="playerInfo.p1"
        />
        <PlayerScoreBoard id="player2Sb" class="playerSb"
            v-bind:playerInfo="playerInfo.p2"
        />
        <div id="generalSb">
            Turn: {{ turn }} <br>
        </div>
        <StartButton
            v-bind:buttonMessage="buttonMessage"
            v-bind:gameStatus="gameStatus"
        />
    </div>
</template>

<script>
import PlayerScoreBoard from './PlayerScoreBoard.vue'
import StartButton from './StartButton.vue'

export default {
    name: 'Scoreboard',
    components: {
        PlayerScoreBoard,
        StartButton,
    },
    props: [
        "turn",
        "currentPlayer",
        "playerInfo",
        "buttonMessage",
        "gameStatus",
    ],
    methods: {
    }
}
</script>

<style scoped>
    #scoreboard {
        background-color:  	#222255;
        min-width: 210px;
        flex-grow: 1;
        color: #2a1a0a;
    }
    #generalSb {
        width: 70%;
        background-color: #AAD;
        margin: 10px auto;
        padding: 5px;
        /* box-shadow: 0px 2px #036; */
    }
    .playerSb {
        background-color: #AAD;
        margin: 10px 0;
        padding: 5px;

        /* box-shadow: 0px 2px #036; */
    }
</style>
