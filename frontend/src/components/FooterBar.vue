<template>
  <footer class="p-4 shadow md:px-6 md:py-8 bg-gray-900">
    <a
      href="https://github.com/thibault-cne/projet_nouveau_site_integration"
      class="block text-sm text-gray-500 sm:text-center dark:text-gray-400 hover:underline"
      >© Made with ❤️ by BigBaz54, Thibz and Yewolf
    </a>
  </footer>
</template>
<script lang="ts">
export default {
  name: "i-footer",
};
</script>
<style scoped lang="scss">
@use "@/assets/styles/scss/standars/colors";
.footer {
  z-index: 9;
  background-color: colors.$bg-footer;
  color: colors.$footer;
}
</style>
