<template>
  <div class="home">
    <div class="header">
      <img src="noun_go_181275_duygu_ozkan.png" />
      <h1>gogogomoku</h1>
    </div>
  </div>
</template>

<script>
export default {
  name: 'GomokuHome',
  props: {

  }
}
</script>

<style scoped>
  .home {
      background-color: #000033;
      width: 100%;
      min-width: fit-content;
      padding: 0;
      margin: 0 auto;
      color: #CDE;
      display: flex;
      flex-direction: column;
      flex-wrap: nowrap;
  }
  
  .header {
    height: auto;
    width: 100%;
    display: flex;
    flex-flow: row nowrap;
    align-items: center;
    justify-content: center;
  }

      .header img {
        display: initial;
        object-fit: contain;
        height: 50px;
        padding-right: 20px;
      }

    @media (min-width: 800px) {
      .home {
        width: 60%;
        min-width: 800px;
      }

      .header img {
        display: initial;
        object-fit: contain;
        height: 50px;
        padding-right: 20px;
      }
    }
</style>
