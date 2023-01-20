<template>
  <v-container>
    <i-editProfile
      :status="status"
      @edit="(u: User) => edit(u)"
    ></i-editProfile>
    <div class="bg-gray-200 pt-4 card mt-4">
      <div class="grid justify-items-end">
        <div class="form-control w-32 p-4">
          <label class="cursor-pointer label">
            <button
              data-modal-target="authentication-modal"
              data-modal-toggle="authentication-modal"
              class="block text-white bg-blue-700 hover:bg-blue-800 focus:ring-4 focus:outline-none focus:ring-blue-300 font-medium rounded-lg text-sm px-5 py-2.5 text-center dark:bg-blue-600 dark:hover:bg-blue-700 dark:focus:ring-blue-800"
              type="button"
            >
              Edit
            </button>
          </label>
        </div>
      </div>
      <i-typeText
        class="p-6 mb-1 text-2xl md:text-4xl"
        :message="status.user.personal_description"
      ></i-typeText>
      <div
        class="container p-6 flex lg:flex-row flex-col justify-between justify-items-center"
      >
        <div class="flex-col lg:w-5/12 8/12 mb-5 lg:mb-0">
          <div class="flex-col">
            <div class="flex items-center">
              <div class="avatar online">
                <div class="w-36 rounded-full">
                  <img src="https://placeimg.com/192/192/people" />
                </div>
              </div>
              <div class="p-4">
                <div class="text-lg font-medium">
                  {{ status.user.name }}
                </div>
                <div class="text-sm font-medium">
                  Année : {{ status.user.current_year }}A
                </div>
                <div class="text-sm font-medium">
                  Promotion : {{ status.user.promotion_year }}
                </div>
              </div>
            </div>
            <div class="flex p-4">
              <div class="flex items-baseline pr-4">
                Nombre de points :
                <span class="flex pl-1"
                  ><i-counter :num="100"></i-counter
                ></span>
              </div>
              <div class="flex-col">
                <div class="flex items-center">
                  <span class="icon pr-2">
                    <v-icon icon="mdi-home" color="primary"></v-icon>
                  </span>
                  {{ status.user.hometown }}
                </div>
                <div class="flex items-center">
                  <span class="icon pr-2">
                    <v-icon icon="mdi-school" color="primary"></v-icon>
                  </span>
                  {{ status.user.studies }}
                </div>
              </div>
            </div>
          </div>
          <div class="divider"></div>
          <div class="flex items-center justify-center">
            <v-btn icon="mdi-facebook" class="p-2 m-2" color="blue"></v-btn>
            <v-btn icon="mdi-google" class="p-2 m-2" color="red"></v-btn>
            <v-btn icon="mdi-instagram" class="p-2 m-2" color="purple"></v-btn>
            <v-btn icon="mdi-snapchat" class="p-2 m-2" color="yellow"></v-btn>
          </div>
        </div>
        <div class="lg:w-6/12 w-12/12">
          <i-stars :status="status" :preload="true"></i-stars>
        </div>
      </div>
    </div>
  </v-container>
</template>
<script lang="ts">
import { defineComponent } from "vue";
import { LoggedIn } from "@/models/LoggedIn";
import ITypeText from "@/components/typingEffect.vue";
import ICounter from "@/components/counter.vue";
import IStars from "@/components/stars.vue";
import IEditProfile from "@/components/editProfile.vue";
import { initModals } from "flowbite";
import { User } from "@/models/User";

export default defineComponent({
  name: "DebugProfile",
  props: {
    status: { type: Object as () => LoggedIn, required: true },
  },
  data() {
    return {
      user: this.status.user,
    };
  },
  mounted() {
    initModals();
  },
  methods: {
    edit(u: User) {
      this.$emit("edit", u);
    },
  },
  components: { ITypeText, ICounter, IStars, IEditProfile },
});
</script>
<style scoped lang="scss"></style>
