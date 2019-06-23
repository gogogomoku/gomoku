<template>
  <div id="app">
    <GomokuHome msg="Welcome to GoGoGomoku"/>
    <Goban v-bind:size="size" v-bind:tab="tab" v-bind:turn="turn"/>
    <button v-on:click="getTab">Update Board</button>
  </div>
</template>

<script>
import GomokuHome from './components/GomokuHome.vue'
import Goban from './components/Goban.vue'
import axios from "axios";

export default {
    name: 'app',
    components: {
        GomokuHome,
        Goban
    },
    data() {
        return {
            turn: 0,
            size: 19,
            tab: [[]]
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
            if (res.code == 1) {
                console.log("game hasn't started");
            } else {
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
                console.log(res.Goban.Turn);
                this._data.tab = newTab
                this._data.size = size
                this._data.turn = res.Turn
            }
            this.sleep(500);
            // this.getTab()
        },
        sleep(ms) {
            return new Promise(resolve => setTimeout(resolve, ms));
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
button {
    padding: 20px;
    font-size: 2vmin;
    border-radius: 3px;
}
</style>
