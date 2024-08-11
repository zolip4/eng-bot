<template>
  <div id="app">
    <h1>Hello from Vue.js frontend!</h1>
    <button @click="fetchMessage">Fetch Message from Backend</button>
    <p>{{ message }}</p>
  </div>
</template>

<script>
export default {
  data() {
    return {
      message: '',
    };
  },
  methods: {
    async fetchMessage() {
      try {
        const response = await fetch('/api/hello');
        const data = await response.text();
        this.message = data;
      } catch (error) {
        console.error('Error fetching message:', error);
      }
    }
  }
}
</script>

<style>
#app {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  text-align: center;
  margin-top: 60px;
}
</style>
