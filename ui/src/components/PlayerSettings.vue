<template>
  <div class="playerSettingsWrapper">
    <div v-if="id === 1" class="playerPieceImgWrapper">
      <img class="playerPieceImg" src="1.png" />
    </div>
    <div v-else-if="id === 2" class="playerPieceImgWrapper">
      <img class="playerPieceImg" src="2.png" />
    </div>
    <div class="sbText">
      <font-awesome-icon
        v-bind:icon="!aiStatus ? 'user' : 'robot'"
        class="aiStatusIndicator interactive"
        @click="onToggleAiStatus()"
      />
      Player {{ id }}
      <br />
    </div>
  </div>
</template>

<script>
export default {
  name: "PlayerSettings",
  props: {
    id: {
      type: Number,
      required: true,
      validator(val) {
        return val == 1 || val == 2;
      }
    },
    aiStatus: {
      type: Number,
      required: true,
      default: 0
    }
  },
  methods: {
    onToggleAiStatus: function() {
      this.$parent.$parent.toggleAiStatus(this.id);
    }
  }
};
</script>

<style scoped>
/* match PlayerScoreBoard.vue for now */
.playerSettingsWrapper {
  border: 1px solid #000000;
  border-radius: 2px;

  margin: 5px 15px;
  padding: 10px 5px 12px;

  background-color: #24252a;

  flex-grow: 0;
  box-sizing: border-box;
  display: flex;
  width: 100%;
  justify-content: space-evenly;
  align-items: center;
}

.playerSb.active {
  border-color: #f3a5a5;
  border-width: 2px;
  background-color: #35363f;
}

.playerPieceImgWrapper {
  flex-basis: 20%;
  height: auto;
}
.playerPieceImg {
  width: 25px;
  border-radius: 50%;
  display: inline-block;
  margin: 1px;
}

.sbText {
  text-align: left;
  flex-basis: 70%;
  line-height: 1.5em;
}

.aiStatusIndicator {
  margin-right: 3px;
}

.aiStatusIndicator.interactive {
  color: orange;
}

.aiStatusIndicator.interactive:hover {
  cursor: pointer;
}
</style>
