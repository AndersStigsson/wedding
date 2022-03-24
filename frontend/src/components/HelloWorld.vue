<template>
  <div
    v-if="!dataSent"
    class="flex flex-col container max-w-lg mt-10 mx-auto w-full items-center justify-center bg-white dark:bg-gray-800 rounded-lg shadow"
  >
    <div class="grid grid-cols-2">
      <div class="col-span-1">
        <button
          class="button text-sm"
          @click="currentLanguage = 'sv'"
        >
          Svenska
        </button>
      </div>
      <div class="col-span-1">
        <button
          class="button text-sm"
          @click="currentLanguage = 'en'"
        >
          English
        </button>
      </div>
    </div>

    <div class="text-3xl">
      {{ msg }}
    </div>
    <img
      alt="Wedding Rings"
      src="../assets/wedding-ring.png"
    >
    <div class="text-2xl">
      {{ translatedWord('happyThatYouWantToCome') }}
    </div>
    <div class="text">
      {{ translatedWord('pleaseFillForm') }}
    </div>
    <div class="text-lg italic">
      {{ translatedWord('unableToCome') }}
      <a
        class="text-blue-500"
        href="mailto:brollop@skafteresort.se"
      >
        brollop@skafteresort.se
      </a>
    </div>
    <div class="grid md:grid-cols-6 grid-cols-1">
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
        <div class="col-span-1 text-lg">
          {{ translatedWord('howMany') }}
        </div>
        <div class="md:col-span-5 md:ml-2 col-span-1">
          <select
            v-model="form.amount"
            class="mt-1 text-lg"
            :style="{border: '2px solid #777'}"
          >
            <option :value=1>
              {{ translatedWord('one') }}
            </option>
            <option :value=2>
              {{ translatedWord('two') }}
            </option>
            <option :value=3>
              {{ translatedWord('three') }}
            </option>
            <option :value=4>
              {{ translatedWord('four') }}
            </option>
          </select>
        </div>
      </div>
      <div
        v-for="n in Number(form.amount)"
        :key="n"
        :index="n"
        class="grid md:grid-cols-6 grid-cols-1 my-5"
      >
        <div class="text-lg col-span-1 md:col-span-2">
          {{ translatedWord('name') }}:
        </div>
        <div class="md:col-span-4 col-span-1">
          <input
            v-model="form.persons[n-1].fname"
            type="text"
            class=" input shadow appearance-none border rounded  py-2 px-3 text-gray-700 leading-tight focus:outline-none focus:shadow-outline"
            :placeholder="translatedWord('firstnameLastname')"
            required
          >
        </div>
        <div class="text-lg col-span-1 md:col-span-2">
          {{ translatedWord('foodPreferences') }}
        </div>
        <div class="md:col-span-4 col-span-1">
          <input
            v-model="form.persons[n-1].foodpreference"
            type="text"
            class=" input shadow appearance-none border rounded py-2 px-3 text-gray-700 leading-tight focus:outline-none focus:shadow-outline"
            :placeholder="translatedWord('foodPreference')"
            required
          >
        </div>
      </div>
      <div class="grid md:grid-cols-6 grid-cols-1">
        <div class="md:col-span-2 col-span-1 text-lg">
          {{ translatedWord('needParking') }}
        </div>
        <div class="md:col-span-4 col-span-1 md:text-right">
          <input
            v-model="form.parking"
            type="checkbox"
            class="checkbox"
          >
        </div>
        <div class="md:col-span-2 col-span-1 text-lg">
          {{ translatedWord('planToSpeak') }}

        </div>
        <div class="md:col-span-3 md:col-start-5 col-span-1 md:ml-3 md:text-right">
          <input
            v-model="form.speech"
            type="checkbox"
            class="checkbox"
          >
        </div>
          <div
            v-if="form.speech"
            class="md:col-span-6 col-span-1"
          >
            {{ translatedWord('ourToastmasterWillContact') }}
          </div>

      <SendButton
        :text="translatedWord('send')"
        class="sendbutton my-5"
        @click="submit"
      />
    </div>
  </div>
  <ThanksPage
    v-else
    :sentData="form"
    :currentLanguage="currentLanguage"
  />
</template>

<script>
import { defineComponent } from 'vue'
import SendButton from './SendButton.vue'
import ThanksPage from './ThanksPage.vue'
export default defineComponent({
  name: 'HelloWorld',
  components: {
    SendButton,
    ThanksPage
  },
  props: {
    msg: {
      type: String,
      default: ''
    }
  },
  data () {
    return {
      dataSent: false,
      form: {
        guests: [],
        persons: [
          {
            fname: '',
            lname: '',
            foodpreference: ''
          },
          {
            fname: '',
            lname: '',
            foodpreference: ''
          },
          {
            fname: '',
            lname: '',
            foodpreference: ''
          },
          {
            fname: '',
            lname: '',
            foodpreference: ''
          }],
        amount: 1,
        parking: false,
        speech: false,
        email: ''
      },
      currentLanguage: 'sv',
    }
  },
  computed: {

    translations() {
      return {
        'happyThatYouWantToCome' : {
          'en' : 'We are happy that you want to come to our wedding',
          'sv' : 'Vad roligt att Ni vill komma på vårt bröllop!'
        },
        'pleaseFillForm' : {
          'en' : 'Please fill the form and click send',
          'sv' : 'Vänligen fyll i formuläret och klicka på skicka in.',
        },
        'unableToCome' : {
          'en' : 'In case you can\'t come, please send an email to',
          'sv' : 'Om Ni inte kan komma på bröllopet, skicka gärna ett mail till',
        },
        'howMany' : {
          'en' : 'How many are you?',
          'sv' : 'Hur många är Ni?',
        },
        'one' : {
          'en' : 'One',
          'sv' : 'En',
        },
        'two' : {
          'en' : 'Two',
          'sv' : 'Två',
        },
        'three' : {
          'en' : 'Three',
          'sv' : 'Tre',
        },
        'four' : {
          'en' : 'Four',
          'sv' : 'Fyra',
        },
        'name' : {
          'en' : 'Name',
          'sv' : 'Namn',
        },
        'firstnameLastname': {
          'en' : 'Firstname and Surname',
          'sv' : 'Förnamn och Efternamn',
        },
        'foodPreference': {
          'en' : 'Food preference',
          'sv' : 'Specialkost',
        },
        'foodPreferences': {
          'en' : 'Potential food preferences',
          'sv' : 'Eventuell specialkost',
        },
        'needParking': {
          'en' : 'Do you need parking?',
          'sv' : 'Behöver Ni parkeringsplats?',
        },
        'planToSpeak': {
          'en' : 'Do you plan to hold a speech or perform a show?',
          'sv' : 'Planerar Ni att hålla tal eller framföra ett spex?',
        },
        'ourToastmasterWillContact': {
          'en' : 'Our toastmasters will contact you with further information about this.',
          'sv' : 'Våra toastmasters kommer att kontakta Er för ytterligare information gällande detta.',
        },
        'send' : {
          'en' : 'Send',
          'sv' : 'Skicka In',
        }

      }
    },
  },
  methods: {
    translatedWord(word) {
      return this.translations[word][this.currentLanguage];
    },
    checkForm () {
      let namesFilled = true

      for (let i = 0; i < this.form.amount; i++) {
        if (this.form.persons[i].fname === '') {
          namesFilled = false
          break
        }
      }
      return namesFilled && this.form.email !== ''
    },
    async submit () {
      //
      if (!this.checkForm()) {
        window.alert('Din email samt namn på samtliga gäster behövs')
        return
      }
      const body = { email: this.form.email, amount: this.form.amount, parking: this.form.parking, speech: this.form.speech, guests: [] }
      for (let i = 0; i < this.form.amount; i++) {
        body.guests.push(
          {
            name: this.form.persons[i].fname,
            foodpreference: this.form.persons[i].foodpreference
          }
        )
      }
      this.form.guests = body.guests
      const res = await fetch(`${process.env.VUE_APP_BACKEND}/add`, {
        method: 'POST',
        body: JSON.stringify(body),
        headers: {
          'x-api-token': process.env.VUE_APP_API_TOKEN
        }
      })

      const returnValue = await res.json()
      console.log(returnValue)
      if (returnValue === 'Success') {
        this.dataSent = true
      }
    }
  }
})
</script>
