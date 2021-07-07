<template>
  <div class="settings">
    <h1 class="title">settings</h1>
    <div class="section">
      <p>Songs root directory</p>
      <input class="input" type="text" v-model="settings.songsDirectory" />
    </div>
    <div class="section">
      <p>Host</p>
      <input class="input" type="text" v-model="settings.host" />
    </div>
    <div class="section">
      <p>Port</p>
      <input class="input" type="text" v-model="settings.port" />
    </div>
    <div class="section">
      <p>Username</p>
      <input class="input" type="text" v-model="settings.username" />
    </div>
    <div class="section">
      <p>Password</p>
      <input class="input" type="text" v-model="settings.password" />
    </div>
    <div>
      <button @click="save" class="save-btn">Save</button>
      <span class="fa fa-check" v-if="saved" />
      <button @click="sync" class="sync-btn">Sync songs</button>
      <span class="fa fa-check" v-if="synced" />
    </div>
  </div>
</template>

<script>
import { SettingsService } from "@/services/SettingsService";
import { SongService } from "@/services/SongService";
export default {
  name: "Settings",
  data() {
    return {
      settings: {},
      saved: false,
      synced: false,
    };
  },
  mounted() {
    SettingsService.fetchSettings().then(
      (settings) => (this.settings = settings)
    );
  },
  methods: {
    save() {
      SettingsService.save(this.settings).then((res) => {
        res.status == 200 ? (this.saved = true) : (this.saved = false);
      });
    },
    sync() {
      SongService.sync().then((res) => {
        res.status == 200 ? (this.synced = true) : (this.synced = false);
      });
    },
  },
};
</script>


<style scoped>
.settings {
  flex-grow: 10;
  background-color: #222831;
  color: #ffffff;
  padding: 1em;
  display: flex;
  flex-direction: column;
}
.section {
  margin-top: 1em;
  padding: 0.5em;
}
.title {
  margin-top: 0;
  font-size: 60pt;
  align-self: center;
  margin-bottom: 0;
}
.input {
  margin-top: 1em;
  border-radius: 20px;
  text-decoration: none;
  text-transform: none;
  border: 1px solid #00adb5;
  box-shadow: none;
  outline: none;
  align-self: center;
  width: 50%;
  height: 25px;
  background-color: #393e46;
  color: whitesmoke;
  padding: 1em;
}
.save-btn {
  margin: 0.5em;
  background-color: #00b58e;
  color: whitesmoke;
  border: none;
  max-width: 100px;
  min-width: 100px;
  padding: 1em;
  border-radius: 20px;
}

.sync-btn {
  margin: 0.5em;
  background-color: #008bb5;
  color: whitesmoke;
  border: none;
  max-width: 100px;
  min-width: 100px;
  padding: 1em;
  border-radius: 20px;
}
</style>