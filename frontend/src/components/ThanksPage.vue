<template>
  <div class="flex flex-col container max-w-lg mt-10 mx-auto w-full items-center justify-center bg-white dark:bg-gray-800 rounded-lg shadow">
    <div class="text-3xl">
      Tack för din anmälan!
    </div>
    <div class="text-3xl">
      Amanda och Anders Bröllop
    </div>
    <div class="text-2xl">
      Vad roligt att Ni vill komma på vårt bröllop!
    </div>
    <div>
      Kolla igenom nedan information och om något har blivit fel, skicka ett mail till <a class="text-blue-500" href="mailto:brollop@skafteresort.se">brollop@skafteresort.se</a>
    </div>
    <div>
      Email: {{ sentData.email }}
    </div>
    <div>
      Antal: {{ sentData.amount }}
    </div>
    <div v-for="person in sentData.guests"
        :key="person.fname"
    >
      <div>
      Namn: {{ person.fname }}
      </div>
      <div>
      Specialkost: {{ person.foodpreference }}
      </div>
    </div>
    <div v-if="sentData.parking">
      Ni behöver parkeringsplats
    </div>
    <div v-else>
      Ni behöver ingen parkeringsplats
    </div>
    <div v-if="sentData.speech">
      Ni vill hålla tal/framföra ett spex
    </div>
    <div v-else>
      Ni planerar inte att hålla tal/framföra ett spex
    </div>
  </div>

</template>

<script>
export default {
  props: {
    sentData: {
      type: Object,
      default: null
    }
  }
}
</script>
