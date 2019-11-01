<template>
  <div id="gameContainer">
    <Goban
      v-bind:currentPlayer="currentPlayer"
      v-bind:size="size"
      v-bind:suggestedPosition="suggestedPosition"
      v-bind:suggestorOn="suggestorOn"
      v-bind:tab="tab"
      v-bind:turn="turn"
    />
    <ControlPanel
      v-bind:buttonMessage="buttonMessage"
      v-bind:currentPlayer="currentPlayer"
      v-bind:gameStatus="gameStatus"
      v-bind:playerInfo="playerInfo"
      v-bind:suggestorOn="suggestorOn"
      v-bind:suggestionTimer="suggestionTimer"
      v-bind:turn="turn"
    />
  </div>
</template>

<script>
import Goban from "./Goban.vue";
import ControlPanel from "./scoreBoard/ControlPanel.vue";

export default {
  name: "GameContainer",
  components: {
    Goban,
    ControlPanel
  },
  props: [
    "buttonMessage",
    "currentPlayer",
    "gameStatus",
    "playerInfo",
    "size",
    "suggestedPosition",
    "suggestionTimer",
    "suggestorOn",
    "tab",
    "turn",
    "winner"
  ],
  methods: {}
};
</script>

<style scoped>
#gameContainer {
  display: flex;
  flex-direction: row;
  flex-wrap: nowrap;
  justify-content: space-around;

  margin: 0px auto;
  width: 900px;
}
</style>
