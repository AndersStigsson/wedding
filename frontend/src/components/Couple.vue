<template>
  <div
    class="flex flex-col container max-w-lg mt-10 mx-auto w-full bg-white dark:bg-gray-800 rounded-lg shadow"
  >
    <div class="row">
      <div class="col">
        Totalt antal: {{guests.length}}
      </div>
      <div class="col">
        Totalt antal parkeringar: {{ totalParkings }}
      </div>
    </div>
    <table class="table table-responsive table-auto">
      <thead>
        <tr>
          <th>
            Namn
          </th>
          <th>
            Parkering
          </th>
          <th>
            Matgnäll
          </th>
          <th>
            Email
          </th>
        </tr>
      </thead>
      <tbody>
        <tr v-for="guest in guests" :key="guest.id">
          <td>{{ guest.name }}</td>
          <td v-if="guest.parking === true">
            Ja
          </td>
          <td v-else>
            Nej
          </td>
          <td>{{ guest.foodpreference }}</td>
          <td>{{ guest.email }}</td>
        </tr>
      </tbody>
    </table>
  </div>
</template>

<script>
export default {
  data() {
    return {
      guests: {
        type: Array,
        default: []
      },
      totalParkings: {
        type: Number,
        default: 0
      }
    };
  },
  mounted() {
    this.fetchData();
  },
  methods: {
    async fetchData() {
      const res = await fetch(`${process.env.VUE_APP_BACKEND}/couple`, {
        method: "GET",
        headers: {
          "x-api-token": process.env.VUE_APP_API_TOKEN
        }
      });

      this.guests = await res.json();
      const parkings = await fetch(`${process.env.VUE_APP_BACKEND}/totalParkings`, {
        method: "GET",
        headers: {
          "x-api-token": process.env.VUE_APP_API_TOKEN
        }
      });

      this.totalParkings = await res.json();
    }
  },
};
</script>
