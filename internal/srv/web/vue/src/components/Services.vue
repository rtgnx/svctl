<template>
  <div class="">
    <h1 class="text-2xl font-bold mb-4 text-center">{{ hostname }}</h1>
    <table id="service-table" class="w-full table-auto border-2 border-dashed border-green-500 p-4">
      <thead>
        <tr>
          <th>Name</th>
          <th>PID</th>
          <th>Timestamp</th>
          <th>Duration</th>
          <th>State</th>
        </tr>
      </thead>
      <tbody id="service-body">
        <tr v-for="item in getServices" :key="item.Pid" :class="computeStyle(item)">
          <td class="px-4">{{ item.Name }}</td>
          <td class="px-4">{{ item.Pid }}</td>
          <td class="px-4">{{ item.Timestamp }}</td>
          <td class="px-4">{{ secondsToHumanReadable(item.Duration) }}</td>
          <td class="px-4">{{ stateToText(item) }}</td>
        </tr>
      </tbody>
    </table>
  </div>
</template>

<style scoped>
.table {
  max-height: 400px;
  overflow-y: auto;
}
</style>
<script setup></script>

<script>
export default {
  props: {
    hostname: String,
    data: Array
  },
  methods: {
    stateToText(item) {
      return (item.State == 1) ? "Running" : (item.State == 0) ? "Down" : "Finished";
    },
    computeStyle(item) {
      return {
        'hover:text-lime-800': true,
        'hover:underline': true,
        'text-green-400': item.Want == 1 && item.State == 1 && item.Duration > 60,
        'text-amber-400': (item.Want == 0 && item.State == 0),
        'text-red-600': (item.Want == 1 && item.State == 0) || (item.Want == 1 && item.State == 1 && item.Duration < 60),
      }
    },
    secondsToHumanReadable(seconds) {
      if (seconds <= 0) {
        return '0s';
      }

      const weeks = Math.floor(seconds / (7 * 24 * 60 * 60));
      const days = Math.floor((seconds % (7 * 24 * 60 * 60)) / (24 * 60 * 60));
      const hours = Math.floor((seconds % (24 * 60 * 60)) / (60 * 60));
      const minutes = Math.floor((seconds % (60 * 60)) / 60);
      const remainingSeconds = seconds % 60;

      const parts = [];

      if (weeks > 0) {
        parts.push(`${weeks}w`);
      }
      if (days > 0) {
        parts.push(`${days}d`);
      }
      if (hours > 0) {
        parts.push(`${hours}h`);
      }
      if (minutes > 0) {
        parts.push(`${minutes}m`);
      }
      if (remainingSeconds > 0) {
        parts.push(`${remainingSeconds}s`);
      }

      return parts.join('');
    },
  },
  computed: {
    getServices() {
      console.log(this.hostname)
      if (this.hostname == null) return [];
      const services = this.data.reduce((mergedServices, entry) => {
        if (entry.hostname === this.hostname) {
          // Merge the Services array for the matching hostname
          mergedServices.push(...entry.data[0].Services);
        }
        return mergedServices;
      }, []);
      console.log(services)
      return services;
    }
  },
}
</script>