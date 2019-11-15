<template>
  <div id="gobanContainer">
    <div id="goban">
      <div class="row" v-for="(line, posY) in tab" :key="posY">
        <div class="tile" v-for="(tile, posX) in line" :key="posX + (posY * size)">
          <!-- {{posX + (posY * size)}} -->
          <div class="tileImage" v-if="gameStatus === 0 || gameStatus === 2">
            <div class="tileSvgContainer">
              <font-awesome-icon
                :icon="icon(tile)"
                :style="{ color: color(tile) }"
                class="tileSvg"
                :class="{postgameTile: inPostgame}"
                size="2x"
              />
            </div>
          </div>
          <div
            class="tileImage"
            v-else-if="gameStatus === 1 && posX + (posY * size) == suggestedPosition && suggestorOn"
          >
            <div
              class="tileSvgContainer"
              :style="{cursor: 'pointer'}"
              v-on:click="clickTile(posX + (posY * size), currentPlayer)"
            >
              <font-awesome-icon
                :icon="iconTileFilled"
                :color="color(currentPlayer)"
                :style="{ visibility: 'visible' }"
                :id="posX + (posY * size) + '-filled'"
                size="2x"
                class="tileSvgFilled svgSuggested"
              />
            </div>
          </div>
          <div class="tileImage" v-else-if="gameStatus === 1">
            <div v-if="tile === 1 || tile === 2" class="tileSvgContainer">
              <font-awesome-icon
                :icon="iconTileFilled"
                :style="{ color: color(tile) }"
                class="tileSvg"
                size="2x"
              />
            </div>
            <div
              v-else-if="tile === 0"
              class="tileSvgContainer"
              :style="{cursor: 'pointer'}"
              v-on:mouseover="mouseOverSvg(posX + (posY * size), currentPlayer)"
              v-on:mouseleave="mouseOutSvg(posX + (posY * size))"
              v-on:click="clickTile(posX + (posY * size), currentPlayer)"
            >
              <font-awesome-icon
                :icon="iconTileEmpty"
                :style="{ color: color0, visibility: 'visible' }"
                :id="posX + (posY * size) + '-empty'"
                class="tileSvg"
                size="2x"
              />
              <font-awesome-icon
                :icon="iconTileFilled"
                :style="{ visibility: 'hidden' }"
                :id="posX + (posY * size) + '-filled'"
                size="2x"
                class="tileSvgFilled"
              />
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
      color0: "#333",
      colorP1: "black",
      colorP2: "#a9a9a9"
    };
  },
  computed: {},
  methods: {
    icon: function(value) {
      return value === 0 ? this.iconTileEmpty : this.iconTileFilled;
    },
    color: function(value) {
      return value === 1
        ? this.colorP1
        : value === 2
        ? this.colorP2
        : this.color0;
    },
    mouseOverSvg: function(tileId, currentPlayer) {
      document.getElementById(`${tileId}-empty`).style.visibility = "hidden";
      document.getElementById(`${tileId}-filled`).style.visibility = "visible";
      document.getElementById(`${tileId}-filled`).style.color = this.color(
        currentPlayer
      );
    },
    mouseOutSvg: function(tileId) {
      document.getElementById(`${tileId}-filled`).style.visibility = "hidden";
      document.getElementById(`${tileId}-empty`).style.visibility = "visible";
      document.getElementById(`${tileId}-empty`).style.color = this.color0;
    },
    clickTile: function(tileId, currentPlayer) {
      // eslint-disable-next-line
      console.log(
        "Make move: \nID: " + tileId + " currentPlayer" + currentPlayer
      );
      document.getElementById(`${tileId}-filled`).style.opacity = 1;
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
  height: 700px;
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
  /* border: 1px pink solid; */
}

.tile {
  height: auto;
  width: 100%;
}

.tileSvgContainer {
  position: relative;
  width: 100%;
}

.tileSvg {
  position: absolute;
  top: 0px;
  left: 0px;
  z-index: 0;
}

.tileSvg.tileSvgFilled {
  z-index: 1;
  position: absolute;
}

.postgameTile {
  cursor: initial;
  opacity: 0.3;
}

.svgSuggested {
  opacity: 0.2;
}

.tileSvgSuggested:hover {
  opacity: 0.5;
}
</style>
