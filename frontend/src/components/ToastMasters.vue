<template>
  <div
    class="flex flex-col container max-w-lg mt-10 mx-auto w-full items-center justify-center bg-white dark:bg-gray-800 rounded-lg shadow"
  >
    <table class="table table-responsive table-auto">
      <thead>
        <tr>
          <th>
            Namn
          </th>
          <th>
            Email
          </th>
          <th>
            Tal
          </th>
        </tr>
      </thead>
      <tbody>
        <tr
          v-for="guest in guests"
          :key="guest.id"
        >
          <td> {{ guest.name }} </td>
          <td> {{ guest.email }} </td>
          <td v-if="guest.speech === true">
            Ja
          </td>
          <td v-else>
            Nej
          </td>
        </tr>
      </tbody>
    </table>
  </div>
</template>

<script>
export default {
  data () {
    return {
      guests: {
        type: Array,
        default: []
      }
    }
  },
  mounted () {
    this.fetchData()
  },
  methods: {
    async fetchData () {
      const res = await fetch(`${process.env.VUE_APP_BACKEND}/toastmasters`, {
        method: 'GET',
        headers: {
          'x-api-token': process.env.VUE_APP_API_TOKEN
        }
      })

      this.guests = await res.json()
    }
  }
}
</script>
