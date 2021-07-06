<template>
  <div class="sidebar">
    <Search />
    <Menu />
    <hr />
    <Playlists />
    <Player />
  </div>
</template>

<script>
import Search from "./Search.vue";
import Playlists from "./Playlists.vue";
import Player from "./Player.vue";
import Menu from "./Menu.vue";

export default {
  name: "Sidebar",
  components: {
    Search,
    Playlists,
    Player,
    Menu,
  },
};
</script>

<style>
.sidebar {
  flex-grow: 1;
  display: flex;
  flex-direction: column;
  background-color: #393e46;
  min-width: 18%;
  max-width: 18%;
}

.sidebar-item {
  color: #ffffff;
  margin-left: 1em;
  margin-right: 1em;
  text-decoration: none;
  padding: 1em;
  border-top: 2px solid transparent;
  border-bottom: 2px solid transparent;
}

.playlist {
  flex-grow: 7;
}
hr {
  border-color: #00adb5;
  background-color: #00adb5;
  min-width: 90%;
}

.sidebar-item:hover {
  background-color: #222831;
}

.selected {
  background-color: #222831;
  color: #00adb5;
  font-weight: 800;
}
</style>