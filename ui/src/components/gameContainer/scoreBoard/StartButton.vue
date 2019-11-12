<template>
  <button v-on:click="buttonAction">{{ buttonMessage }}</button>
</template>

<script>
import { NOT_STARTED, RUNNING, CONCLUDED } from '../../../constants';
export default {
  name: "StartButton",
  props: ["gameStatus"],
  methods: {
    buttonAction: function() {
      if (this.gameStatus === NOT_STARTED) {
        this.$parent.$parent.$parent.startGame(false);
      } else {
        this.$parent.$parent.$parent.restartGame(false);
      }
    }
  },
  computed: {
    buttonMessage() {
      switch (this.gameStatus) {
        case NOT_STARTED:
          return "Start Game"
        case RUNNING:
          return "Restart Game"
        case CONCLUDED:
          return "Rematch"
        default:
          return "Start Game?"
      }
    }
  }
};
</script>

<style scoped>
button {
  padding: 9px 18px 12px;
  background-color: #3cadad;
  border: none;
  font-size: 16px;
  font-weight: bolder;
  letter-spacing: normal;
  border-radius: 3px;
}
button:hover {
  background-color: #48cece;
}
button:active {
  background-color: #2b7c7c;
}
</style>
