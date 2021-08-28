<template>
  <button>
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

<style lang="scss" scoped>

button {
    background-color: green;
    color: white;
    border-radius: 5px;
}

</style>
