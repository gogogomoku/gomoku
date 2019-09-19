<template>
    <div id="scoreboard">
        <StartButton
            id="startButton"
            v-bind:buttonMessage="buttonMessage"
            v-bind:gameStatus="gameStatus"
        />
        <PlayerScoreBoard id="player1Sb"
            v-bind:aiStatus="aiStatus"
            v-bind:currentPlayer="currentPlayer"
            v-bind:playerInfo="playerInfo.p1"
        />
        <PlayerScoreBoard id="player2Sb"
            v-bind:aiStatus="aiStatus"
            v-bind:currentPlayer="currentPlayer"
            v-bind:playerInfo="playerInfo.p2"
        />
        <div id="generalSb">
            Turn: {{ Math.floor(turn/2) }} <br>
        </div>

        <div id="gameOptions">
            <h2>Settings</h2>
            <input type="checkbox" id="checkbox" v-model="checked" @change="onToggleSuggestor($event)" disabled>
            <label for="checkbox">enable suggestor</label>
        </div>

        <Timer
            v-if="suggestorOn && gameStatus > 0"
            v-bind:turn="turn"
        />

    </div>
</template>

<script>
import PlayerScoreBoard from './PlayerScoreBoard.vue'
import StartButton from './StartButton.vue'
import Timer from './Timer.vue'

export default {
    name: 'Scoreboard',
    components: {
        PlayerScoreBoard,
        StartButton,
        Timer,
    },
    props: [
        "aiStatus",
        "buttonMessage",
        "currentPlayer",
        "gameStatus",
        "playerInfo",
        "suggestorOn",
        "turn",
    ],
    data() {
        return {
            checked: true
        }
    },
    methods: {
        onToggleSuggestor: function($event) {
            this.$parent.$parent.toggleSuggestor($event.returnValue)
        }
    }
}
</script>

<style scoped>
    #scoreboard {
        /* TODO: Organize control panel including duplicate child template styles */
        font-family: 'Rubik', 'Avenir', Helvetica, Arial, sans-serif;
        color: #bbb;
        font-size: small;
        text-transform: lowercase;
        letter-spacing: .04em;

        background-color:#1c1d21;

        display: flex;
        flex-direction: column;
        flex-wrap: nowrap;
        flex-grow: 0;
        flex-shrink: 0;
        flex-basis: 210px;
        justify-content: flex-start;
        align-items: center;

        padding: 10px;

        border: 1px solid #76767a;
        border-left-width: 0;
        border-radius: 0px 10px 10px 0px;
    }

    #startButton {
        font-family: 'Share Tech Mono', monospace;
        order: 4;
        width: 75%;
    }

    #generalSb {
        margin: 10px 0;
        padding: 5px;
        flex-grow: 0;
    }

    #gameOptions {
        color: #bbb;

        border: 1px solid #000000;
        border-radius: 8px;

        margin: 5px 15px;
        padding: 10px 5px 12px;

        background-color:#24252a;

        flex-grow: 0;
        width: 100%;
        box-sizing: border-box;
        order: 7;
        margin-top: auto;
        text-align: left;
        flex-grow: 0;
        justify-self: flex-end;
    }

    #gameOptions label {
        margin-left: 10px;
    }

    #gameOptions h2 {
        margin: 0px 5px 10px;
    }

</style>
