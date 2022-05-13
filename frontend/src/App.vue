<template>
  <div>
  <!-- <div v-if="loggedIn">
    <NavBar />
  </div> -->
  <router-view />
 </div>
</template>

<script>
import { authStore } from "./store/authStore";
import NavBar from "./components/NavBar.vue";

export default {
    data() {
        return {
            loggedIn: authStore.getters.loggedIn,
        };
    },
    components: { NavBar }
}
</script>

<style lang="scss">
#app {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
}

nav {
  padding: 30px;

  a {
    font-weight: bold;
    color: #2c3e50;

    &.router-link-exact-active {
      color: #42b983;
    }
  }
}
</style>
