<template>
  <v-container>
    <div class="bg-gray-200 pt-4 card mt-4">
      <div class="grid justify-items-end">
        <div class="form-control w-32 p-4">
          <label class="cursor-pointer label">
            <span class="label-text">Edit</span>
            <input
              type="checkbox"
              class="toggle toggle-success"
              v-model="edit"
            />
          </label>
        </div>
      </div>
      <i-typeText
        class="p-6 mb-1 text-2xl md:text-4xl"
        message="hi, I'm thibault-cne 👋"
      ></i-typeText>
      <div class="container p-6 flex justify-around">
        <div class="flex-col max-w-xs">
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
                <div class="flex items-center" v-if="!edit">
                  <span class="icon pr-2">
                    <v-icon icon="mdi-home" color="primary"></v-icon>
                  </span>
                  {{ status.user.hometown }}
                </div>
                <input
                  v-else
                  type="text"
                  placeholder="Ville d'origine"
                  v-model="user.hometown"
                  class="input input-bordered input-sm w-full max-w-xs text-white"
                />
                <div class="flex items-center" v-if="!edit">
                  <span class="icon pr-2">
                    <v-icon icon="mdi-school" color="primary"></v-icon>
                  </span>
                  {{ status.user.studies }}
                </div>
                <input
                  v-else
                  type="text"
                  placeholder="Etudes"
                  v-model="user.studies"
                  class="input input-bordered input-sm w-full max-w-xs text-white"
                />
                <input
                  v-if="edit"
                  type="text"
                  placeholder="Description"
                  v-model="user.personal_description"
                  class="input input-bordered input-sm w-full max-w-xs text-white"
                />
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
        <i-stars :status="status" :preload="true" class="md:h-96"></i-stars>
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

export default defineComponent({
  name: "DebugProfile",
  props: {
    status: { type: Object as () => LoggedIn, required: true },
  },
  data() {
    return {
      edit: false,
      user: this.status.user,
    };
  },
  components: { ITypeText, ICounter, IStars },
});
</script>
<style scoped lang="scss"></style>
