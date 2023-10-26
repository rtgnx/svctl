<template>
  <aside class="">
    <nav>
      <ul class="space-y-2 p-4">
        <li v-for="h in data" class="border-b border-dashed border-green-500 py-2">
          <a href="#" class="text-green-400 hover:text-red-600 px-1 py-1" @click="selectHost" :data-value="h.hostname">{{h.hostname}}</a>
        </li>
      </ul>
    </nav>
  </aside>
</template>

<script setup>
</script>
<style scoped></style>
<script>

import axios from 'axios';

export default {
  data() {
    return {
      data: null,
    };
  },
  methods: {
    selectHost(event) {
      const hostname = event.target.getAttribute('data-value');
      this.$emit("hostSelected", hostname);
    },
    fetchData() {
      axios.get('https://svd.fly.dev/api/v1/state')
        .then(response => {
          this.data = response.data;
          console.log(this.data)
        })
        .catch(error => {
          console.error('Error fetching data:', error);
        });
    },
  },
  mounted() {
    setInterval(() => {
      this.fetchData();
    }, 2000);
  }
};
</script>