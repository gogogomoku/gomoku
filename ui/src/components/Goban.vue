<template>
    <div>
        <div id="goban">
            <div class="row" v-for="line in tab">
                <div class="tile" v-for="tile in line">
                    <div class="tileAlpha" v-if="tile === 0">
                        <img src="0.png" />
                    </div>
                    <div v-else-if="tile === 1">
                        <img src="1.png" />
                    </div>
                    <div v-else-if="tile === 2">
                        <img src="2.png" />
                    </div>
                </div>
            </div>
            Size: {{ size }} <br>
            Turn: {{ turn }} <br>
        </div>
    </div>
</template>

<script>
export default {
    name: 'Goban',
    props: ["size", "tab", "turn"],
    methods: {

    }
}
</script>

<style scoped>
    #goban {
        background-color: #369;
        width: 50%;
        margin: 5px auto;
        border-radius: 5px;
        padding: 10px;
    }
    .tile {
        width: 25px;
        border-radius: 50%;
        display: inline-block;
        margin: 1px;
    }
    .tile img {
        width: 25px;
    }
    .tileAlpha img {
        opacity: 0.5;
    }


</style>
