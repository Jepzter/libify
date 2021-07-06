<template>
  <Sidebar />
  <router-view />
</template>

<script>
import Sidebar from "@/components/sidebar/Sidebar.vue";

export default {
  name: "App",
  components: {
    Sidebar,
  },
};
</script>
