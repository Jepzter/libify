<template>
  <div :class="[isPlaying(), 'song']" @dblclick="playSong">
    <p class="song-text name">{{ song.title }}</p>
    <p class="song-text artist">{{ song.artist }}</p>
    <p class="song-text duration">{{ song.duration }}</p>
  </div>
</template>

<script>
export default {
  name: "Song",
  props: {
    song: Object,
    playlistId: Number,
  },
  methods: {
    isPlaying() {
      return this.$store.state.currentSong.id == this.song.id ? "playing" : "";
    },
    playSong() {
      this.$store.dispatch("setCurrentPlaylistId", this.playlistId);
      this.$store.dispatch("setCurrentSong", this.song);
    },
  },
};
</script>

<style scoped>
.song {
  display: flex;
  flex-direction: row;
  justify-content: center;
  padding: 1em;
  border: 2px solid transparent;
}
.playing {
  border-top: 2px solid #00adb5;
  border-bottom: 2px solid #00adb5;
  color: #00adb5;
}

.song-text {
  flex-grow: 1;
}

.artist {
  min-width: 45%;
  max-width: 45%;
}

.name {
  min-width: 45%;
  max-width: 45%;
}

.song:hover {
  background-color: #393e46;
}

.duration {
  text-align: right;
  max-width: 10%;
  min-width: 10%;
}
</style>