<template>
  <v-container>
    <div class="relative overflow-x-auto shadow-md sm:rounded-lg">
      <div class="flex items-center justify-end pb-4 bg-white dark:bg-gray-900">
        <label for="table-search" class="sr-only">Search</label>
        <div class="relative">
          <div
            class="absolute inset-y-0 left-0 flex items-center pl-3 pointer-events-none"
          >
            <svg
              class="w-5 h-5 text-gray-500 dark:text-gray-400"
              aria-hidden="true"
              fill="currentColor"
              viewBox="0 0 20 20"
              xmlns="http://www.w3.org/2000/svg"
            >
              <path
                fill-rule="evenodd"
                d="M8 4a4 4 0 100 8 4 4 0 000-8zM2 8a6 6 0 1110.89 3.476l4.817 4.817a1 1 0 01-1.414 1.414l-4.816-4.816A6 6 0 012 8z"
                clip-rule="evenodd"
              ></path>
            </svg>
          </div>
          <input
            type="text"
            id="table-search-users"
            class="block p-2 pl-10 text-sm text-gray-900 border border-gray-300 rounded-lg w-80 bg-gray-50 focus:ring-blue-500 focus:border-blue-500 dark:bg-gray-700 dark:border-gray-600 dark:placeholder-gray-400 dark:text-white dark:focus:ring-blue-500 dark:focus:border-blue-500"
            placeholder="Search for users"
          />
        </div>
      </div>
      <table class="w-full text-sm text-left text-gray-500 dark:text-gray-400">
        <thead
          class="text-xs text-gray-700 uppercase bg-gray-50 dark:bg-gray-700 dark:text-gray-400"
        >
          <tr>
            <th scope="col" class="px-6 py-3">Name</th>
            <th scope="col" class="px-6 py-3">Year</th>
            <th scope="col" class="px-6 py-3">Promo</th>
            <th scope="col" class="px-6 py-3">Statut</th>
            <th scope="col" class="px-6 py-3">Points</th>
            <th scope="col" class="px-6 py-3">Action</th>
          </tr>
        </thead>
        <tbody>
          <tr
            class="bg-white border-b dark:bg-gray-800 dark:border-gray-700 hover:bg-gray-50 dark:hover:bg-gray-600"
            v-for="user in users"
            :key="user.id"
          >
            <th
              scope="row"
              class="flex items-center px-6 py-4 text-gray-900 whitespace-nowrap dark:text-white"
            >
              <!-- <img class="w-10 h-10 rounded-full" src="#" :alt="user.name" /> -->
              <div class="pl-3">
                <div class="text-base font-semibold">{{ getName(user) }}</div>
                <div class="font-normal text-gray-500">
                  {{ user.email }}
                </div>
              </div>
            </th>
            <td class="px-6 py-4">{{ user.current_year }}</td>
            <td class="px-6 py-4">{{ user.promotion_year }}</td>
            <td class="px-6 py-4">{{ user.user_type }}</td>
            <td class="px-6 py-4">{{ user.points }}</td>
            <td class="px-6 py-4">
              <edit-user :props-user="user" />
              <a
                type="button"
                :data-modal-target="user.id"
                :data-modal-show="user.id"
                class="font-medium text-blue-600 dark:text-blue-500 hover:underline"
                >Edit user</a
              >
            </td>
          </tr>
        </tbody>
      </table>
    </div>
  </v-container>
</template>
<script lang="ts">
import { User, getName } from "@/models/User";
import { getRequest } from "@/requests/getRequests";
import { defineComponent } from "vue";
import { initModals } from "flowbite";
import editUser from "@/components/editUser.vue";

export default defineComponent({
  name: "userManagement",
  data() {
    return {
      users: [] as User[],
    };
  },
  mounted() {
    getRequest("/user/get/all/", "json")
      .then((r) => {
        if (r.status == 200) {
          for (let element of r.data) {
            this.users.push(element as User);
          }
        }
      })
      .finally(() => {
        initModals();
      });
  },
  methods: {
    getName(user: User) {
      return getName(user);
    },
  },
  components: { editUser },
});
</script>
<style scoped lang="scss"></style>
