<template>
  <div id="app">
    <GomokuHome :httpError="httpError" />
    <SettingsModal
      v-if="showModal"
      :suggestorOn="suggestorOn"
      :playerInfo="playerInfo"
      :showModal="showModal"
    />
    <GameContainer
      v-bind:currentPlayer="currentPlayer"
      v-bind:httpPending="httpPending"
      v-bind:httpError="httpError"
      v-bind:gameStatus="gameStatus"
      v-bind:playerInfo="playerInfo"
      v-bind:postgameInfo="postgameInfo"
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
import { TAB, NOT_STARTED, CONCLUDED, RUNNING } from "./constants";
import axios from "axios";
import { cloneDeep, merge } from "lodash";

const initialAppState = {
  currentPlayer: -1,
  gameStatus: NOT_STARTED,
  httpError: "",
  httpPending: true,
  http_endpoint: process.env.VUE_APP_SERVER_HTTP || "http://localhost:4242",
  size: 19,
  suggestedPosition: -1,
  suggestionTimer: 0,
  suggestorOn: false,
  tab: cloneDeep(TAB),
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
  postgameInfo: {
    inPostgame: false,
    tab: cloneDeep(TAB),
    playerInfo: {
      p1: {
        AiStatus: -1,
        Id: 1,
        CapturedPieces: -1,
        PiecesLeft: -1
      },
      p2: {
        AiStatus: -1,
        Id: 2,
        CapturedPieces: -1,
        PiecesLeft: -1
      }
    },
    totalTurns: -1,
    winner: 0
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
    // merge(this.$data, initialAppState);
    this.httpPending = true;
    axios
      .get(this._data.http_endpoint)
      .then(response => this.updateTab(response))
      .catch(err => (this.httpError = err.message))
      .finally(
        sleep(2000).then(() => {
          this.httpPending = false;
        })
      );
  },
  methods: {
    playerById(playerId) {
      return this._data.playerInfo[`p${playerId}`] || null;
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
          const postgameInfo = {
            inPostgame: true,
            tab: cloneDeep(this.tab),
            playerInfo: cloneDeep(this.playerInfo),
            totalTurns: this.turn,
            winner: this.winner
          };
          merge(this.$data, initialAppState);
          this._data.gameStatus = CONCLUDED;
          this._data.postgameInfo = postgameInfo;
        } else if (res.CurrentPlayer.AiStatus === 1) {
          await sleep(100);
          this.makeMove(res.SuggestedPosition, res.CurrentPlayer.Id);
        }
      } else {
        // eslint-disable-next-line
        console.log(`(debug) res.Goban undefined`);
      }
    },
    makeMove(tileId, currentPlayer) {
      // console.log(`this.showModal:`, this.showModal);
      if (
        !this.showModal &&
        this._data.gameStatus === RUNNING &&
        this._data.winner === 0
      ) {
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
          .post(this._data.http_endpoint + "/start", {
            AiStatus1: this.playerById(1).AiStatus,
            AiStatus2: this.playerById(2).AiStatus
          })
          .then(response => this.updateTab(response));
      }
    },
    openRestartDialog() {
      merge(this._data.postgameInfo, initialAppState.postgameInfo);
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
      if (this.gameStatus === RUNNING || this.gameStatus === CONCLUDED)
        this.restartGame(true);
      else if (this.gameStatus === NOT_STARTED) this.startGame(true);
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
