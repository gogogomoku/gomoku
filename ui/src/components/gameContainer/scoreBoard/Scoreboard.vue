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
    </div>
</template>

<script>
import PlayerScoreBoard from './PlayerScoreBoard.vue'

export default {
    name: 'Scoreboard',
    components: {
        PlayerScoreBoard,
    },
    props: ["turn", "currentPlayer", "playerInfo"],
    methods: {
    }
}
</script>

<style scoped>
    #scoreboard {
        background-color: #47A;
        width: 25%;
        min-width: 200px;
        margin: 10px;
        padding: 5px;
    }
    #generalSb {
        background-color: #69C;
        margin: 10px;
        padding: 5px;
    }
    .playerSb {
        background-color: #8BD;
        margin: 10px;
        padding: 5px;
        color: #036;
    }
</style>
