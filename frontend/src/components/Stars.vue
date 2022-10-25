<template>
  <v-card class="scrollable" :onscroll="checkFetch">
    <v-card-text>
      <v-timeline align="start" :side="tlSide">
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
              {{ star.receiver_name || status.user.name }}
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
import { useDisplay } from "vuetify";

export default defineComponent({
  name: "i-stars",
  props: {
    status: { type: Object as () => LoggedIn, required: true },
    preload: Boolean,
  },
  data() {
    return {
      stars: [] as Stars[],
      page: -1,
      lock: false,
    };
  },
  mounted() {
    if (this.preload) {
      this.stars = this.status.user.stars;
      this.lock = true;
    } else {
      this.fetchStars();
    }
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
        this.fetchStars();
      }
      if (this.scrollPercentage() == 100) {
        setTimeout(() => {
          this.fetchStars();
        }, 500);
      }
    },
    fetchStars() {
      if (this.lock) return;
      this.lock = true;
      getRequest("/stars/" + (this.page + 1), "")
        .then((r) => {
          if (r.status == 200) {
            console.debug("Getting stars from page:", this.page + 1);
            this.page++;
            for (const star of r.data) {
              this.stars.push(star);
            }
          } else {
            console.debug("No more stars to fetch");
          }
        })
        .finally(() => {
          // timeout lock to prevent spamming
          setTimeout(() => {
            this.lock = false;
          }, 1000);
        });
    },
  },
  computed: {
    tlSide() {
      let { xs } = useDisplay();
      if (xs.value) {
        return "end";
      }
      return undefined;
    },
  },
});
</script>
<style scoped lang="scss">
.scrollable {
  overflow-y: scroll;
}
</style>
