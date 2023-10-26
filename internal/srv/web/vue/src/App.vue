<script setup>
import Sidebar from './components/Sidebar.vue';
import Services from './components/Services.vue';

</script>

<template>
  <div class="flex h-screen">
    <div class="p-4 bg-black w-64 w-1/6">
      <Sidebar @hostSelected="selectedHost"></Sidebar>
    </div>
    <div class="p-4 flex-1 overflow-y-auto">
      <Services :hostname="hostname" :data="data"></Services>
    </div>
  </div>
</template>

<style scoped>
</style>

<script>
import axios from 'axios';

export default {
  components: {
    Sidebar, Services,
  },
  data() {
    return {
      data: null,
      hostname: null,
    };
  },
  methods: {
    selectedHost(data) {
      console.log(data);
      this.hostname = data;
    },
    fetchData() {
      axios.get('/api/v1/state')
        .then(response => {
          this.data = response.data;
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