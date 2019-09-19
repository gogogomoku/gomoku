<template>
    <div>
        <div class="playerSb">
            <div v-if="playerInfo.Id === 1" class="playerPieceImgWrapper">
                <img class="playerPieceImg" src="1.png" />
            </div>
            <div v-else-if="playerInfo.Id === 2" class="playerPieceImgWrapper">
                <img class="playerPieceImg" src="2.png" />
            </div>
            <div class="sbText">Player: {{ playerInfo.Id }} <br>
            captures: {{ playerInfo.CapturedPieces }} <br>
            Pieces left: {{ playerInfo.PiecesLeft }} <br>
            </div>
        </div>
    </div>
</template>

<script>
export default {
    name: 'Scoreboard',
    props: ["aiStatus", "turn", "currentPlayer", "playerInfo"],
    methods: {
    }
}
</script>

<style scoped>
    .playerSb {
        display: flex;
        width: 100%;
        justify-content: space-evenly;
        align-items: center;
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
</style>
