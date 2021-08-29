<template>
  <div class="hello content">
    <h1>
      {{ msg }}
    </h1>
    <img
      alt="Wedding Rings"
      src="../assets/wedding-ring.png"
    >
    <h2>
      Vad roligt att Ni vill komma på vårt bröllop!
    </h2>
    <p>
      Vänligen fyll i formuläret och klicka på skicka in.
    </p>
    <form id="send-form">
      <!-- <p class="field has-addons has-addons-centered control"> -->
      <div class="field is-horizontal">
        <p class="field-label is-normal">
          <label class="label">
            Email:
          </label>
        </p>
        <p class="field-body">
          <input
            v-model="email"
            type="email"
            required
            class="input is-dark"
          >
        </p>
      </div>
      <!-- </p> -->
      <div class="field is-horizontal">
        <p class="field-label is-normal">
          <label class="label">
            Hur många är Ni?
          </label>
        </p>
        <p class="field-body">
          <select
            v-model="amount"
            class="select is-expanded"
          >
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
      </div>
      <div
        v-for="n in Number(amount)"
        :key="n"
        :index="n"
        class="field is-horizontal"
      >
        <p class="field-label is-normal">
          <label class="label">
            Namn:
          </label>
        </p>
        <p class="field-body">
          <input
            v-model="persons[n-1].fname"
            type="text"
            class="input is-dark is-expanded"
            placeholder="Förnamn och Efternamn"
            required
          >
        </p>
      </div>
      <div class="field is-horizontal">
        <p class="field-label is-normal">
          <label class="label">
            Behöver Ni parkeringsplats?
          </label>
        </p>
        <p class="field-body">
          <input
            v-model="parking"
            type="checkbox"
            class="checkbox"
          >
        </p>
      </div>
      <div class="field is-horizontal">
        <p class="field-label is-normal">
          <label class="label">
            Planerar Ni att hålla tal eller framföra ett spex?
          </label>
        </p>
        <div class="field-body">
          <input
            v-model="speech"
            type="checkbox"
            class="checkbox"
          >
          <p
            v-if="speech"
            class="help"
          >
            Våra toastmasters kommer att kontakta Er för ytterligare information gällande detta.
          </p>
        </div>
      </div>

      <div class="field is-horizontal">
        <p class="field-label is-normal">
          <label class="label">
            Specialkost?
          </label>
        </p>
        <div class="field-body">
          <p>
            <input
              v-model="foodpreference"
              type="checkbox"
              class="checkbox"
            >
          </p>
          <div
            v-if="foodpreference"
            class="field-body"
          >
            <!-- <p
              v-if="foodpreference"
              class="field"
            > -->
            <textarea
              v-model="foodpreferences"
              class="is-dark field-body"
              placeholder="Vänligen specificera vad för specialkost och hur många personer"
            />
            <!-- </p> -->
          </div>
        </div>
      </div>
      <SendButton
        text="Skicka In"
        class="sendbutton"
      />
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
      persons: [
        {
          fname: '',
          lname: ''
        },
        {
          fname: '',
          lname: ''
        },
        {
          fname: '',
          lname: ''
        },
        {
          fname: '',
          lname: ''
        }],
      amount: '1',
      parking: false,
      speech: false,
      foodpreference: false,
      foodpreferences: '',
      email: ''
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
.extrainfo {
  font-size:0.8rem;
  margin-block-start:0;
}
.speechbox {
  margin-block-end: 0;
}
.sendbutton {
  margin-top: 10px;
}
.leveled {
  align-content: center;
}
.half-width {
  width: 50%;
}
</style>
