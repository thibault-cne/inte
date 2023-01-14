<template>
  <span class="single-digit">
    <span v-if="d == '.'" data-test="digit">.</span>
    <span v-else-if="d == ' '" data-test="digit">&nbsp;</span>
    <span v-else :style="{ '--value': d }" data-test="digit"></span>
  </span>
</template>

<script lang="ts">
import { defineComponent } from "vue";
export default defineComponent({
  name: "i-digit",
  props: {
    digit: { type: String, required: true },
  },
  data() {
    return {
      d: "0",
    };
  },
  mounted() {
    setTimeout(() => {
      this.d = this.digit;
    }, 100);
  },
});
</script>

<style scoped lang="scss">
.single-digit {
  line-height: 1em;
}
.single-digit {
  display: inline-flex;
}
.single-digit > * {
  height: 1em;
  display: inline-block;
  overflow-y: hidden;
}
.single-digit > *:before {
  position: relative;
  content: "0\A 1\A 2\A 3\A 4\A 5\A 6\A 7\A 8\A 9";
  white-space: pre;
  top: calc(var(--value) * -1em);
  text-align: center;
  transition: all 1s cubic-bezier(1, 0, 0, 1);
}
</style>
