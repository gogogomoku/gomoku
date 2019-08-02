<template>
    <div id="scoreboard">
        <StartButton
            id="startButton"
            v-bind:buttonMessage="buttonMessage"
            v-bind:gameStatus="gameStatus"
        />
        <PlayerScoreBoard id="player1Sb" class="playerSb"
            v-bind:playerInfo="playerInfo.p1"
        />
        <PlayerScoreBoard id="player2Sb" class="playerSb"
            v-bind:playerInfo="playerInfo.p2"
        />
        <div id="generalSb">
            Turn: {{ turn }} <br>
        </div>
        
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
        background-color:#222255;
        width: 100%;
        min-width: fit-content;
        display: flex;
        flex-wrap: wrap;
        flex-direction: row;
        align-items: center;
        justify-content: space-around;

        font-size: small;
    }

    #startButton {
        flex-basis: 100%;
        flex-grow: 1;
        flex-shrink: 0;
        order: 4;
    }

    #generalSb {
        flex-basis: 1;
        flex-grow: 1;
        flex-shrink: 0;
        margin: 10px auto;
        padding: 5px;
        color: #AAD;
    }

    .playerSb {
        flex-grow: 1;
        flex-basis: 1;
        flex-shrink: 0;
        color: #AAD;
        margin: 10px 0;
        padding: 5px;
    }

    @media (min-width: 800px) {
        #scoreboard {
            flex-basis: content;
            display: flex;
            flex-wrap: wrap;
            flex-direction: column;
            flex-wrap: nowrap;
            flex-grow: 0;
            flex-shrink: 1;
            min-width: 210px;
            /* color: #2a1a0a; <--- good dark */
            justify-content: flex-start;
            align-items: center;
        }
        #generalSb {
            width: 70%;
            margin: 10px auto;
            padding: 5px;
            flex-grow: 0;
        }
        .playerSb {
            margin: 10px 0;
            padding: 5px;
            flex-grow: 0;
        }
        #startButton {
            flex-basis: auto;
            flex-grow: 0;
        }
    }
</style>
