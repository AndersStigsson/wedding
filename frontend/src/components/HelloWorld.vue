<template>
  <div class="flex flex-col container max-w-lg mt-10 mx-auto w-full items-center justify-center bg-white dark:bg-gray-800 rounded-lg shadow">
    <div class="text-3xl">
      {{ msg }}
    </div>
    <img
      alt="Wedding Rings"
      src="../assets/wedding-ring.png"
    >
    <div class="text-2xl">
      Vad roligt att Ni vill komma på vårt bröllop!
    </div>
    <div class="text">
      Vänligen fyll i formuläret och klicka på skicka in.
    </div>
    <!--<form id="send-form">-->
    <div class="grid md:grid-cols-6 md:gap-1 grid-cols-1">
      <!-- <p class="field has-addons has-addons-centered control"> -->
        <div class="col-span-1 text-lg">
          Email:
        </div>
        <div class="md:col-span-5 col-span-1">
          <input
            v-model="form.email"
            type="email"
            required
            class=" input shadow appearance-none border rounded w-full py-2 px-3 text-gray-700 leading-tight focus:outline-none focus:shadow-outline"
          >
        </div>
      <!-- </p> -->
        <div class="col-span-1 text-lg">
            Hur många är Ni?
        </div>
        <div class="md:col-span-5 col-span-1">
          <select
            v-model="form.amount"
            class="form-select block mt-1 text-lg"
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
        </div>
      <div
        v-for="n in Number(form.amount)"
        :key="n"
        :index="n"
        class="md:col-span-6 col-span-1"
      >
        <div class="text-lg col-span-1">
            Namn:
        </div>
        <div class="md:col-span-2 col-span-1">
          <input
            v-model="form.persons[n-1].fname"
            type="text"
            class=" input shadow appearance-none border rounded  py-2 px-3 text-gray-700 leading-tight focus:outline-none focus:shadow-outline"
            placeholder="Förnamn och Efternamn"
            required
          >
        </div>
        <div class="text-lg col-span-1">
            Eventuell specialkost
        </div>
        <div class="md:col-span-2 col-span-1">
          <input
            v-model="form.persons[n-1].foodpreference"
            type="text"
            class=" input shadow appearance-none border rounded py-2 px-3 text-gray-700 leading-tight focus:outline-none focus:shadow-outline"
            placeholder="Specialkost"
            required
          >
        </div>
      </div>
        <div class="md:col-span-3 col-span-1 text-lg">
            Behöver Ni parkeringsplats?
        </div>
        <div class="md:col-span-3 col-span-1">
          <input
            v-model="form.parking"
            type="checkbox"
            class="checkbox"
          >
        </div>
        <div class="md:col-span-3 col-span-1">
          <label class="label">
            Planerar Ni att hålla tal eller framföra ett spex?
          </label>
        </div>
        <div class="md:col-span-3 col-span-1">
          <input
            v-model="form.speech"
            type="checkbox"
            class="checkbox"
          >
          <p
            v-if="form.speech"
            class="help"
          >
            Våra toastmasters kommer att kontakta Er för ytterligare information gällande detta.
          </p>
        </div>

      <SendButton
        text="Skicka In"
        class="sendbutton"
      />
    <!--</form>-->
    </div>
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
      form: {
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
  },
  methods: {
    submit () {
      // Send to database to save.
      // Either using Laravel/PHP or a Node.js backend.
    }
  }
})
</script>
