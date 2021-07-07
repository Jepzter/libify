<template>
  <div class="playlist">
    <p @click="createPlaylist" class="create-playlist">
      <span class="fa fa-plus"></span>New playlist
    </p>
    <div v-for="playlist in playlists" :key="playlist.id" class="playlists">
      <router-link
        :to="{ path: `/playlist/${playlist.id}` }"
        class="sidebar-item"
        :class="{
          selected: $route.params?.id == playlist.id,
          playing: currentPlaylistId == playlist.id,
        }"
        @dblclick="editName(playlist)"
      >
        {{ playlist.name }}
      </router-link>
    </div>
    <div class="modal" v-if="edit.isEditingName">
      <div class="flex">
        <h1>Change name</h1>
        <span class="fa fa-times" @click="closeEdit"></span>
      </div>
      <div class="flex">
        <input
          type="text"
          class="input-bar"
          v-model="edit.editingPlaylist.name"
        />
        <span class="fa fa-save" @click="saveEdit"></span>
      </div>
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
      edit: {
        isEditingName: false,
        editingPlaylist: {},
        oldName: "",
      },
    };
  },
  mounted() {
    PlaylistService.fetchPlaylists().then(
      (playlists) => (this.playlists = playlists)
    );
    this.$store.watch(
      (state) => state.currentPlaylistId,
      (playlistId) => {
        this.currentPlaylistId = playlistId;
      }
    );
  },
  methods: {
    createPlaylist() {
      const playlist = {
        name: `New playlist #${this.playlists.length + 1}`,
      };
      PlaylistService.createPlaylist(playlist).then((res) => {
        if (res.status === 201) {
          PlaylistService.fetchPlaylists().then(
            (playlists) => (this.playlists = playlists)
          );
        }
      });
    },
    editName(playlist) {
      this.edit.isEditingName = true;
      this.edit.editingPlaylist = playlist;
      this.edit.oldName = playlist.name;
    },
    saveEdit() {
      PlaylistService.updatePlaylist(this.edit.editingPlaylist).then((res) => {
        this.edit.isEditingName = false;
        if (res.status === 200) {
          PlaylistService.fetchPlaylists().then(
            (playlists) => (this.playlists = playlists)
          );
        }
      });
    },
    closeEdit() {
      this.edit.isEditingName = false;
      this.edit.editingPlaylist.name = this.edit.oldName;
    },
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

.create-playlist {
  margin-left: 2em;
  margin-right: 1em;
  color: #00adb5;
  margin-bottom: 0.5em;
}
.create-playlist span {
  margin-right: 0.3em;
}
.modal {
  position: absolute;
  padding: 0.5em;
  left: 40%;
  top: 40%;
  background-color: #00adb5;
  width: 400px;
  height: 70px;
  color: whitesmoke;
  box-shadow: 0px 0px 5px whitesmoke;
  display: flex;
  justify-content: space-between;
  flex-direction: column;
}
.flex {
  display: flex;
  flex-direction: row;
  justify-content: space-between;
  align-items: baseline;
}

.input-bar {
  text-decoration: none;
  text-transform: none;
  border: none;
  box-shadow: none;
  outline: none;
  height: 25px;
  background-color: #393e46;
  color: whitesmoke;
  padding: 1em;
  margin-top: 0.5em;
  min-width: 90%;
}
</style>