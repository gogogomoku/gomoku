<template>
  <div class="home">
    <h1>{{ msg }}</h1>
  </div>
</template>

<script>
export default {
  name: 'Home',
  props: {
    msg: String
  }
}
</script>

<style scoped>
  .home {
      background-color: #69C;
      width: 50%;
      padding: 10px;
      border-radius: 10px;
      margin: 0 auto;
      color: #CDE;
  }
</style>
