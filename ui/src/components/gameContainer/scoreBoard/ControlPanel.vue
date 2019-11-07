<template>
  <div id="scoreboard">
    <StartButton
      id="startButton"
      v-bind:buttonMessage="buttonMessage"
      v-bind:gameStatus="gameStatus"
      v-bind:disabled="gameStatus > 0 && playerInfo.p1.AiStatus == 10"
    />
    <PlayerScoreBoard
      id="player1Sb"
      v-bind:gameStatus="gameStatus"
      v-bind:currentPlayer="currentPlayer"
      v-bind:playerInfo="playerInfo.p1"
    />
    <PlayerScoreBoard
      id="player2Sb"
      v-bind:gameStatus="gameStatus"
      v-bind:currentPlayer="currentPlayer"
      v-bind:playerInfo="playerInfo.p2"
    />
    <div id="generalSb">
      Turn: {{ Math.floor(turn/2) }}
      <br />
      Winner: {{ winner }}
      <br />
      Game status: {{ gameStatus }}
    </div>

    <Timer
      v-if="suggestorOn && gameStatus > 0"
      v-bind:turn="turn"
      v-bind:suggestionTimer="suggestionTimer"
    />
  </div>
</template>

<script>
import PlayerScoreBoard from "./PlayerScoreBoard.vue";
import StartButton from "./StartButton.vue";
import Timer from "./Timer.vue";

export default {
  name: "Scoreboard",
  components: {
    PlayerScoreBoard,
    StartButton,
    Timer
  },
  props: {
    "buttonMessage": String,
    "currentPlayer": Number,
    "gameStatus": Number,
    "playerInfo": Object,
    "suggestorOn": Boolean,
    "suggestionTimer": Number,
    "turn": Number,
    "winner": Number
  },
};
</script>

<style scoped>
#scoreboard {
  font-family: "Rubik", "Avenir", Helvetica, Arial, sans-serif;
  color: #bbb;
  font-size: small;
  text-transform: lowercase;
  letter-spacing: 0.04em;

  background-color: #1c1d21;

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
  font-family: "Share Tech Mono", monospace;
  order: 4;
  width: 75%;
}

#generalSb {
  margin: 10px 0;
  padding: 5px;
  flex-grow: 0;
}
</style>
