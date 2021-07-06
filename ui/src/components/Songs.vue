<template>
  <div class="songs">
    <div class="titles">
      <p class="name">Name</p>
      <p class="artist">Artist</p>
      <p class="duration">Duration</p>
    </div>
    <div v-for="song in songs" :key="song.id">
      <Song :song="song" :playlistId="playlistId" />
    </div>
  </div>
</template>

<script>
import Song from "./Song.vue";
export default {
  Name: "Songs",
  components: {
    Song,
  },
  props: {
    songs: {
      type: Array,
      required: true,
    },
    playlistId: Number,
  },
};
</script>

<style>
.songs {
  display: flex;
  flex-direction: column;
}

.titles {
  display: flex;
  flex-direction: row;
  justify-content: center;
  padding: 0.9em;
  font-weight: 900;
  font-size: 14pt;
}

.artist {
  min-width: 45%;
  max-width: 45%;
}

.name {
  min-width: 45%;
  max-width: 45%;
}

.duration {
  text-align: right;
  max-width: 10%;
  min-width: 10%;
}
</style>