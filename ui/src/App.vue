<template>
  <div id="app">
    <GomokuHome msg="Welcome to GoGoGomoku"/>
    <Goban v-bind:size="size" v-bind:tab="tab" v-bind:turn="turn" v-bind:currentPlayer="currentPlayer"/>
    <StartButton v-bind:buttonMessage="buttonMessage" v-bind:gameStatus="gameStatus"/>
  </div>
</template>

<script>
import GomokuHome from './components/GomokuHome.vue'
import Goban from './components/Goban.vue'
import axios from "axios"
import StartButton from './components/StartButton.vue'

export default {
    name: 'app',
    components: {
        GomokuHome,
        Goban,
        StartButton,
    },
    data() {
        return {
            turn: 0,
            size: 19,
            tab: [[]],
            currentPlayer: 1,
            buttonMessage: "Start Game",
            gameStatus: 0
        }
    },
    methods: {
        getTab() {
            axios.get("http://localhost:4242")
            .then(response => this.updateTab(response))
        },
        updateTab(response) {
            console.log(response.data);
            var res = response.data

                var size = res.Goban.Size
                var tab = res.Goban.Tab
                var newTab = []
                for (var row = 0; row < size; row++) {
                    var line = []
                    for (var col = 0; col < size; col++) {
                        line.push(tab[(row * size) + col])
                    }
                    newTab.push(line)
                }
                this._data.tab = newTab
                this._data.size = size
                this._data.turn = res.Turn
                this._data.currentPlayer = res.CurrentPlayer.Id
                this._data.gameStatus = res.Status
        },
        makeMove(tileId, currentPlayer) {
            axios.get("http://localhost:4242/move/" + tileId +"/id/" + currentPlayer)
            .then(response => this.updateTab(response))
        },
        startGame() {
            axios.get("http://localhost:4242")
            .then(response => this.updateTab(response))
            if (typeof(this._data.status) == "undefined") {
                axios.get("http://localhost:4242/start")
                .then(response => this.updateTab(response))
                this._data.buttonMessage = "Restart Game"
            }
        },
        restartGame() {
            axios.get("http://localhost:4242/restart")
            .then(response => this.updateTab(response))
        },
    }

}
</script>

<style>
    #app {
      font-family: 'Avenir', Helvetica, Arial, sans-serif;
      -webkit-font-smoothing: antialiased;
      -moz-osx-font-smoothing: grayscale;
      text-align: center;
      color: #2c3e50;
      margin-top: 60px;
    }
    body {
        background-color: #036;
    }
</style>
