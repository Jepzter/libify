<template>
  <div class="menu">
    <router-link
      class="sidebar-item"
      :class="{ selected: $route.path == '/' }"
      to="/"
      >Home</router-link
    >
    <router-link
      class="sidebar-item"
      :class="{ selected: $route.path == '/settings' }"
      to="/settings"
      >Settings</router-link
    >
  </div>
</template>

<script>
export default {
  name: "Menu",
};
</script>

<style scoped>
.menu {
  margin-top: 1em;
  display: flex;
  flex-direction: column;
}
</style>