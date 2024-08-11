<template>
  <div id="app">
    <h1>Hello from Vue.js!</h1>
    <p>Server response: {{ message }}</p>
  </div>
</template>

<script>
export default {
  data() {
    return {
      message: ''
    }
  },
  mounted() {
    fetch('http://localhost:8080')
      .then(response => response.text())
      .then(data => {
        this.message = data;
      });
  }
}
</script>
