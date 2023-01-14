<template>
  <div class="container">
    <span class="typed-text font-mono text-gray-700">{{ typeValue }}</span>
    <span class="cursor" :class="{ typing: typeStatus }">&nbsp;</span>
  </div>
</template>
<script lang="ts">
import { defineComponent } from "vue";
export default defineComponent({
  name: "i-typeText",
  props: {
    message: { type: String, required: true },
  },
  data() {
    return {
      typeValue: "",
      typeStatus: false,
      typingSpeed: 75,
      erasingSpeed: 100,
      newTextDelay: 6000,
      charIndex: 0,
    };
  },
  mounted() {
    setTimeout(this.typeText, 200);
  },
  methods: {
    typeText() {
      if (this.charIndex < this.message.length) {
        if (!this.typeStatus) this.typeStatus = true;
        this.typeValue += this.message.charAt(this.charIndex);
        this.charIndex += 1;
        setTimeout(this.typeText, this.typingSpeed);
      } else {
        this.typeStatus = false;
        setTimeout(this.eraseText, this.newTextDelay);
      }
    },
    eraseText() {
      if (this.charIndex > 0) {
        if (!this.typeStatus) this.typeStatus = true;
        this.typeValue = this.message.substring(0, this.charIndex - 1);
        this.charIndex -= 1;
        setTimeout(this.eraseText, this.erasingSpeed);
      } else {
        this.typeStatus = false;
        setTimeout(this.typeText, this.typingSpeed + 1000);
      }
    },
  },
});
</script>
<style scoped lang="scss">
@use "@/assets/styles/scss/standars/keyframes";
.container {
  display: flex;
  justify-content: center;
  align-items: center;
  span.typed-text {
    color: #d2b94b;
  }
  span.cursor {
    display: inline-block;
    margin-left: 3px;
    width: 4px;
    background-color: black;
    animation: cursorBlink 1s infinite;
  }
  span.cursor.typing {
    animation: none;
  }
}
</style>
