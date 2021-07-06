<template>
  <div class="search">
    <input
      class="search-bar"
      type="text"
      placeholder="Search..."
      @input="search"
    />
  </div>
</template>

<script>
export default {
  Name: "Search",
  methods: {
    search(e) {
      console.log(e.target.value);
    },
  },
};
</script>

<style scoped>
.search {
  display: flex;
  justify-content: space-around;
}
.search-bar {
  margin-top: 1em;
  border-radius: 20px;
  text-decoration: none;
  text-transform: none;
  border: 2px solid #00adb5;
  box-shadow: none;
  outline: none;
  align-self: center;
  width: 85%;
  height: 25px;
  background-color: #393e46;
  color: whitesmoke;
  padding: 1em;
}
</style>