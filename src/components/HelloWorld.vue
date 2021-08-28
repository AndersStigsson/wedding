<template>
  <div class="hello">
    <h1>{{ msg }}</h1>
    <h2>
      Vad roligt att Ni vill komma på vårt bröllop!
    </h2>
    <p>
      Vänligen fyll i formuläret och klicka på skicka in.
    </p>
    <form id="send-form">
      <p>
        Hur många är Ni?
        <select v-model="amount">
          <option value="1">
            En
          </option>
          <option value="2">
            Två
          </option>
          <option value="3">
            Tre
          </option>
          <option value="4">
            Fyra
          </option>
        </select>
      </p>
      <div
        v-for="n in Number(amount)"
        :key="n"
        :index="n"
      >
        <p>
          Förnamn:
          <input
            v-model="fname"
            type="text"
          >
          Efternamn:
          <input
            v-model="lname"
            type="text"
          >
        </p>
      </div>
      <p>
        Behöver Ni parkeringsplats?
        <input
          v-model="parking"
          type="checkbox"
        >
      </p>
      <p>
        Planerar Ni att hålla tal eller framföra ett spex?
        <input
          v-model="speech"
          type="checkbox"
        >
      </p>

      <SendButton text="Skicka In" />
    </form>
  </div>
</template>

<script lang="ts">
import { defineComponent } from 'vue'
import SendButton from './SendButton.vue'
export default defineComponent({
  name: 'HelloWorld',
  components: {
    SendButton
  },
  props: {
    msg: {
      type: String,
      default: ''
    }
  },
  data () {
    return {
      fname: '',
      lname: '',
      amount: '1',
      parking: false,
      speech: false
    }
  }
})
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped lang="scss">
h3 {
  margin: 40px 0 0;
}
ul {
  list-style-type: none;
  padding: 0;
}
li {
  display: inline-block;
  margin: 0 10px;
}
a {
  color: #42b983;
}
form {
  align-content: left;
}
</style>
