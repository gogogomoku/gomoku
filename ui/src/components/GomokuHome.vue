<template>
  <div class="home">
    <div class="header">
      <!-- <img src="noun_go_181275_duygu_ozkan.png" /> -->
      <h1>gogogomoku</h1>
      <div id="author">
        <div id="authorImgWrapper">
          <img src="github-logo.svg" width="15" height="15" />
        </div>
        <div id="authorTxtWrapper">
          <a href="https://github.com/gogogomoku/gomoku">@eskombro + @ekelen</a>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
export default {
  name: "GomokuHome",
  props: {}
};
</script>

<style scoped>
.header {
  height: auto;
  width: 900px;
  margin: 0 auto;
  display: flex;
  flex-flow: row nowrap;
  align-items: baseline;
  justify-content: space-between;

  color: #ece2d0;
}

#author {
  display: flex;
  flex-flow: row nowrap;
  justify-content: center;
}

#author img {
  margin-right: 15px;
}

#author a {
  color: #9f998e;
  text-decoration: none;
  vertical-align: center;
}
</style>
