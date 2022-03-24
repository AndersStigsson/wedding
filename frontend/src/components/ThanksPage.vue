<template>
  <div class="flex flex-col container max-w-lg mt-10 mx-auto w-full items-center justify-center bg-white dark:bg-gray-800 rounded-lg shadow">
    <div class="text-3xl">
      {{ translatedWord('thanks') }}
    </div>
    <div class="text-3xl">
      Amanda och Anders Bröllop
    </div>
    <div class="text-2xl">
      {{ translatedWord('happyThatYouWantToCome') }}
    </div>
    <div>
      {{ translatedWord('checkInformation') }}
      <a class="text-blue-500" href="mailto:brollop@skafteresort.se">brollop@skafteresort.se</a>
    </div>
    <div>
      Email: {{ sentData.email }}
    </div>
    <div>
      {{ translatedWord('amount') }}: {{ sentData.amount }}
    </div>
    <div v-for="person in sentData.guests"
        :key="person.name"
    >
      <div>
        {{ translatedWord('name') }}: {{ person.name }}
      </div>
      <div>
        {{ translatedWord('foodPreference') }}: {{ person.foodpreference }}
      </div>
    </div>
    <div v-if="sentData.parking">
      {{ translatedWord('youNeedParking') }}
    </div>
    <div v-else>
      {{ translatedWord('youDontNeedParking') }}
    </div>
    <div v-if="sentData.speech">
      {{ translatedWord('holdSpeech') }}
    </div>
    <div v-else>
      {{ translatedWord('dontHoldSpeech') }}
    </div>
  </div>

</template>

<script>
export default {
  props: {
    sentData: {
      type: Object,
      default: null
    },
    currentLanguage: {
      type: String,
      default: 'sv'
    }
  },
  computed: {
    translations() {
      return {
        'thanks' : {
          'en' : 'Thank you!',
          'sv' : 'Tack för din anmälan!',
        },
        'happyThatYouWantToCome' : {
          'en' : 'We are happy that you want to come to our wedding',
          'sv' : 'Vad roligt att Ni vill komma på vårt bröllop!'
        },
        'checkInformation': {
          'en' : 'Check below information and if anything is wrong, please send an email to ',
          'sv' : 'Kolla igenom nedan information och om något har blivit fel, skicka ett mail till ',
        },
        'amount' : {
          'en' : 'Amount',
          'sv' : 'Antal',
        },
        'name' : {
          'en' : 'Name',
          'sv' : 'Namn',
        },
        'foodPreference': {
          'en' : 'Food preference',
          'sv' : 'Specialkost',
        },
        'youNeedParking': {
          'en' : 'You need parking',
          'sv' : 'Ni behöver parkeringsplats',
        },
        'youDontNeedParking': {
          'en' : 'You don\t need parking',
          'sv' : 'Ni behöver ingen parkeringsplats',
        },
        'holdSpeech': {
          'en' : 'You plan to hold a speech or perform a show',
          'sv' : 'Ni vill hålla tal/framföra ett spex',
        },
        'dontHoldSpeech': {
          'en' : '',
          'sv' : 'Ni planerar inte att hålla tal/framföra ett spex',
        },
      }
    }

  },
  methods: {
    translatedWord(word) {
      return this.translations[word][this.currentLanguage];
    },
  },
}
</script>
