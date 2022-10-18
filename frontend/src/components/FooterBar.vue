<template>
  <v-footer app fixed height="50px" class="footer">
    <v-col class="text-center" cols="12">
      <v-btn
        href="https://github.com/thibault-cne/projet_nouveau_site_integration"
        icon
        variant="text"
      >
        <v-icon>mdi-github</v-icon>
      </v-btn>

      <strong>Made with ❤️ by BigBaz54, Thib and Yewolf</strong>
    </v-col>
  </v-footer>
</template>
<script lang="ts">
export default {
  name: "FooterBar",
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
