<template>
  <div>
    <textarea v-model="inputValue" @input="handleInputChange"></textarea>
  </div>
  <div id="rendered">
  </div>
</template>

<style>
</style>

<script>
import { marked } from 'marked';

export default {
  data() {
    return {
      inputValue: ''
    };
  },
  methods: {
    handleInputChange() {
      document.getElementById('rendered').innerHTML = marked(this.inputValue)
    }
  }
};
</script>