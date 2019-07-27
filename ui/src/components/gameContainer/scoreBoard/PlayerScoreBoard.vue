<template>
    <div>
        <div class="playerSb">
            <div v-if="playerInfo.Id === 1">
                <img class="playerPieceImg" src="1.png" />
            </div>
            <div v-else-if="playerInfo.Id === 2">
                <img class="playerPieceImg" src="2.png" />
            </div>
            Player: {{ playerInfo.Id }} <br>
            CapturedPieces: {{ playerInfo.CapturedPieces }} <br>
            Pieces left: {{ playerInfo.PiecesLeft }} <br>
        </div>
    </div>
</template>

<script>
export default {
    name: 'Scoreboard',
    props: ["turn", "currentPlayer", "playerInfo"],
    methods: {
    }
}
</script>

<style scoped>
    .playerPieceImg {
        width: 25px;
        border-radius: 50%;
        display: inline-block;
        margin: 1px;
        cursor: pointer;
    }
</style>
