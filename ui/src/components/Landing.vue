<template>
  <div class="landing">
    <h1 class="title">Libify</h1>
    <h3>What is Libify?</h3>
    <p>
      Libify allows you to make your local library be streamed just like with
      spotify with a nice player and option to create playlists.
    </p>
  </div>
</template>

<script>
export default {
  name: "Landing",
};
</script>

<style scoped>
.landing {
  flex-grow: 10;
  background-color: #222831;
  color: #ffffff;
  padding: 1em;
  display: flex;
  flex-direction: column;
}
.title {
  margin-top: 0;
  font-size: 60pt;
  align-self: center;
  margin-bottom: 0;
}
h3 {
  font-size: 20pt;
  margin-bottom: 0;
}
</style>
