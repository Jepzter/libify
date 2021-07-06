<template>
  <div class="playlist">
    <div v-for="playlist in playlists" :key="playlist.id" class="playlists">
      <router-link
        :to="{ path: `/playlist/${playlist.id}` }"
        class="sidebar-item"
        :class="{
          selected: $route.params?.id == playlist.id,
          playing: currentPlaylistId == playlist.id,
        }"
      >
        {{ playlist.name }}
      </router-link>
    </div>
  </div>
</template>

<script>
import { PlaylistService } from "@/services/PlaylistService";

export default {
  Name: "Playlists",
  data() {
    return {
      playlists: [],
      currentPlaylistId: undefined,
    };
  },
  mounted() {
    PlaylistService.fetchPlaylists().then((playlists) => this.playlists = playlists);
    this.$store.watch(
      (state) => state.currentPlaylistId,
      (playlistId) => {
        this.currentPlaylistId = playlistId;
      }
    );
  },
};
</script>
<style scoped>
.header {
  margin-left: 1em;
  margin-right: 1em;
  color: whitesmoke;
}

.playlists {
  display: flex;
  flex-direction: column;
}

.playing {
  border-top: 2px solid #00adb5;
  border-bottom: 2px solid #00adb5;
  color: #00adb5;
}
</style>