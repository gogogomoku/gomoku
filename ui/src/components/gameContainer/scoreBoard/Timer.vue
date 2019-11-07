<template>
  <div id="timer">
    <!-- <div v-if="suggestedPosition && !msElapsedServer"... -->
    <!-- <div v-if-else="msElapsedServer"...> -->
    <div id="time-up">Suggestion computed in</div>
    <div>{{ suggestionTimer }} ms</div>
  </div>
</template>

<script>
export default {
  name: "Timer",
  props: {
    turn: Number,
    suggestionTimer: Number
  }
};
</script>

<style scoped>
#timer {
  flex-basis: 70px;
  border: 1px solid #000000;
  border-radius: 8px;

  margin: 5px 15px;
  padding: 10px 5px 12px;

  background-color: #24252a;

  flex-grow: 0;
  box-sizing: border-box;
  order: 5;

  width: 100%;

  display: flex;
  align-items: center;
  justify-content: space-around;

  margin-top: 15px;
}

#time-up {
  text-align: left;
  font-size: smaller;
}

.lds-ripple {
  display: inline-block;
  position: relative;
  width: 64px;
  height: 64px;
  margin-right: 20px;
}
.lds-ripple div {
  position: absolute;
  border: 4px solid #fff;
  opacity: 1;
  border-radius: 50%;
  animation: lds-ripple 1s cubic-bezier(0, 0.2, 0.8, 1) infinite;
}
.lds-ripple div:nth-child(2) {
  animation-delay: -0.5s;
}
@keyframes lds-ripple {
  0% {
    top: 28px;
    left: 28px;
    width: 0;
    height: 0;
    opacity: 1;
  }
  100% {
    top: -1px;
    left: -1px;
    width: 58px;
    height: 58px;
    opacity: 0;
  }
}
</style>
