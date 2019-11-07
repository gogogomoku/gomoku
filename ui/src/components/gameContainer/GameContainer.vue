<template>
  <div id="gameContainer">
    <Goban
      v-bind:currentPlayer="currentPlayer"
      v-bind:gameStatus="gameStatus"
      v-bind:size="size"
      v-bind:suggestedPosition="suggestedPosition"
      v-bind:suggestorOn="suggestorOn"
      v-bind:tab="postgameInfo.inPostgame ? postgameInfo.tab : tab"
      v-bind:inPostgame="postgameInfo.inPostgame"
      v-bind:turn="turn"
    />
    <ControlPanel
      v-bind:currentPlayer="currentPlayer"
      v-bind:gameStatus="gameStatus"
      v-bind:playerInfo="playerInfo"
      v-bind:postgameInfo="postgameInfo"
      v-bind:suggestionTimer="suggestionTimer"
      v-bind:suggestorOn="suggestorOn"
      v-bind:turn="turn"
      v-bind:winner="winner"
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
  props: {
    "currentPlayer": {
      required: true,
      type: Number
    },
    "gameStatus": {
      required: true,
      type: Number
    },
    "playerInfo": {
      required: true,
      type: Object
    },
    "postgameInfo": {
      required: true,
      type: Object
    },
    "size": {
      required: true,
      type: Number
    },
    "suggestedPosition": {
      required: true,
      type: Number
    },
    "suggestionTimer": {
      required: true,
      type: Number
    },
    "suggestorOn": {
      required: true,
      type: Boolean
    },
    "tab": {
      required: true,
      type: Array
    },
    "turn": {
      required: true,
      type: Number
    },
    "winner": {
      required: true,
      type: Number
    }
  },
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
