<template>
  <div class="home">
    <img alt="Wedding Rings" src="../assets/wedding-ring.png">
    <HelloWorld msg="Amanda och Anders Bröllop"/>
  </div>
</template>

<script lang="ts">
import { defineComponent } from 'vue'
import HelloWorld from '@/components/HelloWorld.vue' // @ is an alias to /src

export default defineComponent({
  name: 'Home',
  components: {
    HelloWorld
  }
})
</script>
