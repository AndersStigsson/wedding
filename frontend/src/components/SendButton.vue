<template>
  <button class="button">
    {{ text }}
  </button>
</template>

<script>
import { defineComponent } from 'vue'

export default defineComponent({
  name: 'SendButton',
  props: {
    text: {
      type: String,
      default: ''
    }
  }
})
</script>

<style scoped>

button {
    background-color: green;
    color: white;
    border-radius: 5px;
}

</style>
