<template>
    <div id="gameContainer">
        <div id="goban">
            <div class="row" v-for="(line, posY) in tab">
                <div class="tile" v-for="(tile, posX) in line">
                    <div class="tileAlpha" v-if="tile === 0">
                        <img
                            v-on:mouseover="mouseOver(posX + (posY * size), currentPlayer)"
                            v-on:mouseleave="mouseOut(posX + (posY * size), tile)"
                            v-on:click="clickTile(posX + (posY * size), currentPlayer)"
                            :id="posX + (posY * size)"
                            class="tile0"
                            src="0.png" />
                    </div>
                    <div v-else-if="tile === 1">
                        <img
                            :id="posX + (posY * size)"
                            class="tile1"
                            src="1.png" />
                    </div>
                    <div v-else-if="tile === 2">
                        <img
                            :id="posX + (posY * size)"
                            class="tile2"
                            src="2.png" />
                    </div>
                </div>
            </div>
        </div>
    </div>
</template>

<script>
// import Scoreboard from './Scoreboard.vue'

export default {
    name: 'Goban',
    components: {
    },
    props: ["size", "tab", "turn", "currentPlayer"],
    methods: {
        mouseOver: function(tileId, currentPlayer){
            document.getElementById(tileId).src=currentPlayer+".png";
            document.getElementById(tileId).opacity=0.5;
        },
        mouseOut: function(tileId, tile){
            document.getElementById(tileId).src=tile+".png";
            document.getElementById(tileId).opacity=1;
        },
        clickTile: function(tileId, currentPlayer){
            this.$parent.$parent.makeMove(tileId, currentPlayer)
        }
    }
}
</script>

<style scoped>
    #gameContainer {
        background-color: #258;
        width: 46%;
        margin: 5px auto;
    }
    #goban {
        background-color: #47A;
        width: 46%;
        min-width: 550px;
        margin: 5px auto;
        border-radius: 5px;
        padding: 10px;
    }
    .tile {
        width: 25px;
        border-radius: 50%;
        display: inline-block;
        margin: 1px;
        cursor: pointer;
    }
    .tile img {
        width: 25px;
    }
    .tileAlpha img {
        opacity: 0.5;
    }


</style>
