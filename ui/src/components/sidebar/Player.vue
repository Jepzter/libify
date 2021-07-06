<template>
  <div class="player" v-if="src != ''">
    <audio
      id="audio-player"
      ref="audioPlayer"
      :src="src"
      type="audio/mpeg"
    ></audio>
    <div class="controls">
      <span @click="play" v-if="!playing" class="fa fa-play fa-3x"></span>
      <span @click="pause" v-else class="fa fa-pause fa-3x" />
    </div>
    <div class="song-name" v-if="song">
      <p>{{ song.title }} - {{ song.artist }}</p>
    </div>
  </div>
  <div class="player" v-else>
    <p class="no-song">Nothing playing</p>
  </div>
</template>

<script>
export default {
  Name: "Player",
  data() {
    return {
      song: {},
      playing: false,
      src: "",
    };
  },
  mounted() {
    this.$store.watch(
      (state) => state.currentSong,
      (song) => {
        this.playing = false;
        this.song = song;
        this.src =
          "http://localhost:8085/api/libify/stream-v1/" +
          this.$store.state.currentSong.fileLocation;
        setTimeout(() => {
          this.play();
        }, 500); // TODO: Remove this stupid bugfix!
      }
    );
  },
  methods: {
    play() {
      if (this.playing) {
        return;
      }
      this.$refs.audioPlayer.play();
      this.playing = true;
    },
    pause() {
      if (!this.playing) {
        return;
      }
      this.$refs.audioPlayer.pause();
      this.playing = false;
    },
  },
};
</script>

<style scoped>
.player {
  background-color: #00adb5;
  color: #ffffff;
  display: flex;
  flex-direction: column;
  justify-content: space-around;
  padding: 1em;
}

.controls {
  display: flex;
  justify-content: space-around;
}

.song-name {
  display: flex;
  justify-content: space-around;
  margin-top: 0.5em;
}

.no-song {
  align-self: center;
}
</style>