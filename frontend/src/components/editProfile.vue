<template>
  <!-- Main modal -->
  <div
    id="authentication-modal"
    tabindex="-1"
    aria-hidden="true"
    class="fixed top-0 left-0 right-0 z-50 hidden w-full p-4 overflow-x-hidden overflow-y-auto md:inset-0 h-modal md:h-full"
  >
    <div class="relative w-full h-full max-w-md md:h-auto">
      <!-- Modal content -->
      <div class="relative bg-white rounded-lg shadow dark:bg-gray-700">
        <button
          type="button"
          class="absolute top-3 right-2.5 text-gray-400 bg-transparent hover:bg-gray-200 hover:text-gray-900 rounded-lg text-sm p-1.5 ml-auto inline-flex items-center dark:hover:bg-gray-800 dark:hover:text-white"
          data-modal-hide="authentication-modal"
        >
          <svg
            aria-hidden="true"
            class="w-5 h-5"
            fill="currentColor"
            viewBox="0 0 20 20"
            xmlns="http://www.w3.org/2000/svg"
          >
            <path
              fill-rule="evenodd"
              d="M4.293 4.293a1 1 0 011.414 0L10 8.586l4.293-4.293a1 1 0 111.414 1.414L11.414 10l4.293 4.293a1 1 0 01-1.414 1.414L10 11.414l-4.293 4.293a1 1 0 01-1.414-1.414L8.586 10 4.293 5.707a1 1 0 010-1.414z"
              clip-rule="evenodd"
            ></path>
          </svg>
          <span class="sr-only">Close modal</span>
        </button>
        <div class="px-6 py-6 lg:px-8">
          <h3 class="mb-4 text-xl font-medium text-gray-900 dark:text-white">
            Edit your profile
          </h3>
          <form class="space-y-6" action="#">
            <div>
              <label
                for="hometown"
                class="block mb-2 text-sm font-medium text-gray-900 dark:text-white"
                >Hometown</label
              >
              <input
                type="text"
                name="hometown"
                id="hometown"
                class="bg-gray-50 border border-gray-300 text-gray-900 text-sm rounded-lg focus:ring-blue-500 focus:border-blue-500 block w-full p-2.5 dark:bg-gray-600 dark:border-gray-500 dark:placeholder-gray-400 dark:text-white"
                placeholder="Paris"
                v-model="user.hometown"
              />
            </div>
            <div>
              <label
                for="studies"
                class="block mb-2 text-sm font-medium text-gray-900 dark:text-white"
                >Studies</label
              >
              <input
                type="text"
                name="studies"
                id="studies"
                placeholder="TELECOM Nancy"
                v-model="user.studies"
                class="bg-gray-50 border border-gray-300 text-gray-900 text-sm rounded-lg focus:ring-blue-500 focus:border-blue-500 block w-full p-2.5 dark:bg-gray-600 dark:border-gray-500 dark:placeholder-gray-400 dark:text-white"
              />
            </div>
            <div>
              <label
                for="description"
                class="block mb-2 text-sm font-medium text-gray-900 dark:text-white"
                >Description</label
              >
              <input
                type="text"
                name="description"
                id="description"
                placeholder="Elle est full option"
                v-model="user.personal_description"
                class="bg-gray-50 border border-gray-300 text-gray-900 text-sm rounded-lg focus:ring-blue-500 focus:border-blue-500 block w-full p-2.5 dark:bg-gray-600 dark:border-gray-500 dark:placeholder-gray-400 dark:text-white"
              />
            </div>
            <button
              type="button"
              :onclick="saveEdit"
              class="w-full text-white bg-blue-700 hover:bg-blue-800 focus:ring-4 focus:outline-none focus:ring-blue-300 font-medium rounded-lg text-sm px-5 py-2.5 text-center dark:bg-blue-600 dark:hover:bg-blue-700 dark:focus:ring-blue-800"
              data-modal-hide="authentication-modal"
            >
              Save edit
            </button>
          </form>
        </div>
      </div>
    </div>
  </div>
</template>
<script lang="ts">
import { LoggedIn } from "@/models/LoggedIn";
import { User } from "@/models/User";
import { postRequest } from "@/requests/postRequests";
import { defineComponent } from "vue";

export default defineComponent({
  name: "i-editProfile",
  props: {
    status: { type: Object as () => LoggedIn, required: true },
  },
  data() {
    return {
      user: {} as User,
    };
  },
  mounted() {
    this.user = this.status.user;
  },
  methods: {
    saveEdit() {
      postRequest(this.user, "user/modify/data", "json");
      this.$emit("edit", this.user);
    },
  },
});
</script>
<style scoped lang="scss"></style>
