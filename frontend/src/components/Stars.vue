<template>
  <v-card class="scrollable" :onscroll="checkFetch">
    <v-card-title>
      <h1 class="text-center">Les étoiles</h1>
    </v-card-title>
    <v-card-text>
      <v-timeline align="start">
        <v-timeline-item
          v-for="(star, i) in stars"
          :key="i"
          :dot-color="['#CD7F32', '#C0C0C0', '#FFD700'][star.type]"
          icon="mdi-star"
          icon-color="white"
          fill-dot
        >
          <v-card width="200px">
            <v-card-title
              :class="['text-h6']"
              :style="
                'background-color: ' +
                ['#CD7F32', '#C0C0C0', '#FFD700'][star.type]
              "
            >
              {{ star.receiver_name }}
            </v-card-title>
            <v-card-text class="white text--primary">
              <p>{{ star.message }}</p>
            </v-card-text>
          </v-card>
        </v-timeline-item>
      </v-timeline>
    </v-card-text>
  </v-card>
</template>
<script lang="ts">
import { LoggedIn } from "@/models/LoggedIn";
import { getRequest } from "@/requests/getRequests";
import { Stars } from "@/models/Stars";
import { defineComponent } from "vue";

export default defineComponent({
  name: "i-stars",
  props: {
    status: { type: Object as () => LoggedIn, required: true },
  },
  data() {
    return {
      stars: [] as Stars[],
      page: 0,
    };
  },
  mounted() {
    this.fetchStars();
  },
  methods: {
    scrollPercentage() {
      const scrollable = document.querySelector(".scrollable");
      if (scrollable) {
        return (
          (100 * (scrollable.scrollTop + scrollable.clientHeight)) /
          scrollable.scrollHeight
        );
      }
      return 0;
    },
    checkFetch() {
      if (this.scrollPercentage() > 90) {
        this.page += 1;
        this.fetchStars();
      }
    },
    fetchStars() {
      getRequest("/stars/" + this.page, "").then((r) => {
        console.log("Getting stars from page", this.page);
        if (r.data.length == 0) {
          this.page--;
        }
        for (const star of r.data) {
          this.stars.push(star);
        }
      });
    },
  },
});
</script>
<style scoped lang="scss">
.scrollable {
  overflow-y: scroll;
}
</style>
