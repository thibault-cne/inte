<template>
  <div
    v-for="digit in split(num)"
    :key="digit"
    class="flex flex-col text-neutral-content justify-center"
  >
    <i-digit :digit="digit"></i-digit>
  </div>
</template>

<script lang="ts">
import { defineComponent } from "vue";
import { splitIntoDigits } from "@/plugins/utils";
import IDigit from "@/components/digit.vue";
export default defineComponent({
  name: "i-counter",
  props: {
    num: { type: Number, required: true },
  },
  components: {
    IDigit,
  },
  methods: {
    split(num: number) {
      console.log(splitIntoDigits(num));
      return splitIntoDigits(num);
    },
  },
});
</script>

<style scoped lang="scss"></style>
