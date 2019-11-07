<template>
  <div id="app">
    <GomokuHome msg="Welcome to GoGoGomoku" />
    <SettingsModal
      v-if="showModal"
      :suggestorOn="suggestorOn"
      :playerInfo="playerInfo"
      :showModal="showModal"
    />
    <GameContainer
      v-bind:buttonMessage="buttonMessage"
      v-bind:currentPlayer="currentPlayer"
      v-bind:gameStatus="gameStatus"
      v-bind:playerInfo="playerInfo"
      v-bind:size="size"
      v-bind:suggestedPosition="suggestedPosition"
      v-bind:suggestionTimer="suggestionTimer"
      v-bind:suggestorOn="suggestorOn"
      v-bind:tab="tab"
      v-bind:turn="turn"
      v-bind:winner="winner"
    />
  </div>
</template>

<script>
import GomokuHome from "./components/GomokuHome.vue";
import GameContainer from "./components/gameContainer/GameContainer.vue";
import SettingsModal from "./components/SettingsModal.vue";
import { TAB } from "./constants";
import axios from "axios";
import { cloneDeep } from "lodash";

const initialAppState = {
  buttonMessage: "Start Game",
  currentPlayer: -1,
  gameStatus: 0,
  http_endpoint: process.env.VUE_APP_SERVER_HTTP || "http://localhost:4242",
  size: 19,
  suggestedPosition: -1,
  suggestionTimer: 0,
  suggestorOn: false,
  tab: TAB,
  turn: 0,
  playerInfo: {
    p1: {
      AiStatus: 1,
      Id: 1,
      CapturedPieces: 0,
      PiecesLeft: 0
    },
    p2: {
      AiStatus: 0,
      Id: 2,
      CapturedPieces: 0,
      PiecesLeft: 0
    }
  },
  winner: 0,
  showModal: false
};

export default {
  name: "app",
  components: {
    GomokuHome,
    GameContainer,
    SettingsModal
  },
  data() {
    return cloneDeep(initialAppState);
  },
  mounted: function() {
    // for development, re-init state on page refresh
    Object.assign(this.$data, initialAppState);
  },
  methods: {
    playerById(playerId) {
      return this._data.playerInfo[`p${playerId}`] || null;
    },
    getTab() {
      axios
        .get(this._data.http_endpoint)
        .then(response => this.updateTab(response));
    },
    async updateTab(response) {
      var res = response.data;
      if (res.Goban != undefined) {
        var size = res.Goban.Size;
        var tab = res.Goban.Tab;
        var newTab = [];
        for (var row = 0; row < size; row++) {
          var line = [];
          for (var col = 0; col < size; col++) {
            line.push(tab[row * size + col]);
          }
          newTab.push(line);
        }
        this._data.tab = newTab;
        this._data.size = size;
        this._data.turn = res.Turn;
        this._data.playerInfo = {
          p1: res.P1,
          p2: res.P2
        };
        this._data.currentPlayer = res.CurrentPlayer.Id;
        this._data.gameStatus = res.Status;
        this._data.suggestedPosition = res.SuggestedPosition;
        this._data.suggestionTimer = res.SuggestionTimer;
        this._data.winner = res.Winner;
        if (res.Winner != 0) {
          alert("Winner: Player " + res.Winner);
          // todo: tell server to refresh too
          Object.assign(this.$data, initialAppState);
        } else if (res.CurrentPlayer.AiStatus === 1) {
          await sleep(100);
          this.makeMove(res.SuggestedPosition, res.CurrentPlayer.Id);
        }
      }
    },
    makeMove(tileId, currentPlayer) {
      if (this._data.gameStatus > 0 && this._data.winner == 0) {
        axios
          .get(
            this._data.http_endpoint +
              "/move/" +
              tileId +
              "/id/" +
              currentPlayer
          )
          .then(response => this.updateTab(response));
      }
    },
    startGame(selectedOptions = true) {
      if (!selectedOptions) this.openRestartDialog();
      else {
        axios
          .get(this._data.http_endpoint)
          .then(response => this.updateTab(response));
        if (typeof this._data.status == "undefined") {
          axios
            .post(this._data.http_endpoint + "/start", {
              AiStatus1: this.playerById(1).AiStatus,
              AiStatus2: this.playerById(2).AiStatus
            })
            .then(response => this.updateTab(response));
          this._data.buttonMessage = "Restart Game";
        }
      }
    },
    openRestartDialog() {
      this.showModal = true;
    },
    restartGame(selectedOptions = true) {
      if (!selectedOptions) this.openRestartDialog();
      else {
        axios
          .post(this._data.http_endpoint + "/restart", {
            AiStatus1: this.playerById(1).AiStatus,
            AiStatus2: this.playerById(2).AiStatus
          })
          .then(response => this.updateTab(response));
      }
    },
    toggleSuggestor() {
      this._data.suggestorOn = !this._data.suggestorOn;
    },
    toggleAiStatus(playerId) {
      const player = this.playerById(playerId);
      if (player) player.AiStatus = !player.AiStatus | 0;
    },
    closeModal() {
      this._data.showModal = false;
      if (this.gameStatus > 0) this.restartGame(true);
      else if (this.gameStatus == 0) this.startGame(true);
    }
  }
};

function sleep(ms) {
  return new Promise(resolve => setTimeout(resolve, ms));
}
</script>

<style>
#app {
  font-family: "Share Tech Mono", "Avenir", Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  margin-top: 20px;
}
body {
  background-color: #121315;
}
</style>
