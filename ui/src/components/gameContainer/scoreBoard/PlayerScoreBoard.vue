<template>
        <div class="playerSb" v-bind:class="{active: currentPlayer === playerInfo.Id }">
            <div v-if="playerInfo.Id === 1" class="playerPieceImgWrapper">
                <img class="playerPieceImg" src="1.png" />
            </div>
            <div v-else-if="playerInfo.Id === 2" class="playerPieceImgWrapper">
                <img class="playerPieceImg" src="2.png" />
            </div>
            <div class="sbText">
                <font-awesome-icon v-bind:icon="!playerInfo.AiStatus ? 'user' : 'robot'" class="aiStatusIndicator" v-bind:class="{interactive: gameStatus === 0}" @click="onToggleAiStatus()"/> Player {{ playerInfo.Id }} <br>
            captures: {{ playerInfo.CapturedPieces }} <br>
            Pieces left: {{ playerInfo.PiecesLeft }} <br>
            </div>
        </div>
</template>

<script>

export default {
    name: 'Scoreboard',
    props: ["gameStatus", "currentPlayer", "playerInfo"],
    methods: {
        onToggleAiStatus: function() {
            if (this.gameStatus === 0)
                this.$parent.$parent.$parent.toggleAiStatus(this.playerInfo.Id)
        }
    }
}
</script>

<style scoped>

    .playerSb {
        border: 1px solid #000000;
        border-radius: 8px;

        margin: 5px 15px;
        padding: 10px 5px 12px;

        background-color:#24252a;

        flex-grow: 0;
        box-sizing: border-box;
        display: flex;
        width: 100%;
        justify-content: space-evenly;
        align-items: center;
    }

    .playerSb.active {
        border-color: #f3a5a5;
        border-width: 2px;
        background-color: #35363f;
    }

    .playerPieceImgWrapper {
        flex-basis: 20%;
        height: auto;
    }
    .playerPieceImg {
        width: 25px;
        border-radius: 50%;
        display: inline-block;
        margin: 1px;
    }

    .sbText {
        text-align: left;
        flex-basis: 70%;
        line-height: 1.5em;
    }

    .aiStatusIndicator {
        margin-right: 3px;
    }

    .aiStatusIndicator.interactive {
        color: orange;
    }

    .aiStatusIndicator.interactive:hover {
        cursor: pointer;
    }

</style>
