<template>
  <button v-on:click="buttonAction" v-bind:disabled="disabled">{{ buttonMessage }}</button>
</template>

<script>
export default {
  name: "StartButton",
  props: ["buttonMessage", "disabled", "gameStatus"],
  methods: {
    buttonAction: function() {
      if (this.gameStatus == 0) {
        this.$parent.$parent.$parent.startGame(false);
      } else {
        this.$parent.$parent.$parent.restartGame(false);
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
