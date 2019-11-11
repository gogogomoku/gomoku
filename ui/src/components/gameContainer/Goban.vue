<template>
  <div id="gobanContainer">
    <div id="goban">
      <div class="row" v-for="(line, posY) in tab" :key="posY">
        <div class="tile" v-for="(tile, posX) in line" :key="posX">
          <!-- {{posX + (posY * size)}} -->
          <div class="tileImage" v-if="posX + (posY * size) == suggestedPosition && suggestorOn">
            <div v-if="currentPlayer === 1">
              <img
                v-on:mouseover="mouseOver(posX + (posY * size), currentPlayer)"
                v-on:mouseleave="mouseOutSuggested(posX + (posY * size), currentPlayer)"
                v-on:click="clickTile(posX + (posY * size), currentPlayer)"
                :id="posX + (posY * size)"
                class="tileSuggested"
                src="1.png"
              />
            </div>
            <div v-else-if="currentPlayer === 2">
              <img
                v-on:mouseover="mouseOver(posX + (posY * size), currentPlayer)"
                v-on:mouseleave="mouseOutSuggested(posX + (posY * size), currentPlayer)"
                v-on:click="clickTile(posX + (posY * size), currentPlayer)"
                :id="posX + (posY * size)"
                class="tileSuggested"
                src="2.png"
              />
            </div>
          </div>
          <div class="tileImage tileAlpha postgameTile" v-else-if="inPostgame">
            <div v-if="tile === 0">
              <img :id="posX + (posY * size)" class="tile0" src="0.png" />
            </div>
            <div v-else-if="tile === 1">
              <img :id="posX + (posY * size)" class="tile1" src="1.png" />
            </div>
            <div v-else-if="tile === 2">
              <img :id="posX + (posY * size)" class="tile2" src="2.png" />
            </div>
          </div>
          <div class="tileImage" v-else>
            <div v-if="tile === 0">
            <font-awesome-icon
                :icon="iconTileEmpty"
                :style="{ color: '#333' }"
                :id="posX + (posY * size)"
                size=2x
                v-on:mouseover="mouseOverSvg(posX + (posY * size), currentPlayer)"
                v-on:mouseleave="mouseOutSvg(posX + (posY * size), tile)"
                v-on:click="clickTile(posX + (posY * size), currentPlayer)"

              />
              <!-- v-on:mouseover="mouseOverSvg(posX + (posY * size), currentPlayer)" -->
            <!-- <div class="tileAlpha" v-if="tile === 0">
              <img
                v-on:mouseover="mouseOver(posX + (posY * size), currentPlayer)"
                v-on:mouseleave="mouseOut(posX + (posY * size), tile)"
                v-on:click="clickTile(posX + (posY * size), currentPlayer)"
                :id="posX + (posY * size)"
                class="tile0"
                src="0.png"
              /> -->
            </div>
            <div v-else-if="tile === 1">
              <img :id="posX + (posY * size)" class="tile1" src="1.png" />
            </div>
            <div v-else-if="tile === 2">
              <img :id="posX + (posY * size)" class="tile2" src="2.png" />
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
export default {
  name: "Goban",
  components: {},
  props: [
    "currentPlayer",
    "gameStatus",
    "inPostgame",
    "size",
    "suggestedPosition",
    "suggestorOn",
    "tab",
    "turn"
  ],
  data() {
    return {
      iconTileEmpty: ["far", "times-circle"],
      iconTileFilled: ["fas", "times-circle"],
      colorSpace: "#333",
      colorP1: "black",
      colorP2: "#a9a9a9",
    }
  },
  computed: {

  },
  methods: {
    playerColor: function() {
      console.log('this.currentPlayer: ', this.currentPlayer);
      return this.currentPlayer === 1 ? this.colorP1 : (this.currentPlayer === 2 ? this.colorP2 : this.colorSpace)
    },
    icon: function(value) {
      return value === 0 ? this.iconTileEmpty : this.iconTileFilled;
    },
    color: function(value) {
      return value === 1 ? this.colorP1 : (value === 2 ? this.colorP2 : this.colorSpace)
    },
    mouseOver: function(tileId, currentPlayer) {
      document.getElementById(tileId).src = currentPlayer + ".png";
      document.getElementById(tileId).opacity = 0.5;
    },
    mouseOverSvg: function(tileId, currentPlayer) {
      document.getElementById(tileId).style.color = this.color(currentPlayer);
      document.getElementById(tileId).style.opacity = 0.3;
    },
    mouseOut: function(tileId, tile) {
      document.getElementById(tileId).src = tile + ".png";
      document.getElementById(tileId).opacity = 1;
    },
    mouseOutSvg: function(tileId, tile) {
      // this.$refs[tileId][0].icon = this.icon(tile); // todo - make it go
      document.getElementById(tileId).style.color = this.color(tile);
      document.getElementById(tileId).style.opacity = 1;
    },
    mouseOutSuggested: function(tileId, tile) {
      document.getElementById(tileId).src = tile + ".png";
      document.getElementById(tileId).opacity = 0.5;
    },
    clickTile: function(tileId, currentPlayer) {
      // eslint-disable-next-line
      console.log(
        "Make move: \nID: " + tileId + " currentPlayer" + currentPlayer
      );
      this.$parent.$parent.makeMove(tileId, currentPlayer);
    }
  }
};
</script>

<style scoped>
#gobanContainer {
  background-color: #454649;
  display: flex;
  flex-basis: 700px;
  margin: 0;
  padding: 0;
  flex-grow: 0;
  flex-shrink: 0;

  border: 1px solid #76767a;
  border-radius: 10px 0px 0px 10px;
  box-sizing: border-box;
  border-right-width: 0px;
}

#goban {
  padding: 10px;
  margin: 0 auto;
  flex-basis: 700px;
  flex-grow: 0;
  display: flex;
  flex-wrap: wrap;
  box-sizing: border-box;
}

.row {
  flex-basis: 100%;
  display: flex;
  align-items: stretch;
  justify-content: center;
}

.tile {
  /* width: 4.7%; */
  display: inline-block;
  margin: 0.25%;
  cursor: pointer;
}
.tile img {
  width: 100%;
}
.tileAlpha img {
  opacity: 0.5;
}
.tileSuggested {
  filter: contrast(70%);
  /* filter: blur(1px); */
  opacity: 0.3;
}
.postgameTile {
  cursor: initial;
}
</style>
