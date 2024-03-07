<template>
  <div>
    <textarea v-model="inputValue" @input="handleInputChange"></textarea>
  </div>
  <div id="rendered">
  </div>
  <button @click="saveNote">Save</button>
</template>

<style>
</style>

<script>
import { marked } from 'marked';
import axios from "axios";

export default {
  data() {
    return {
      inputValue: ''
    };
  },
  methods: {
    handleInputChange() {
      document.getElementById('rendered').innerHTML = marked(this.inputValue)
    },
    saveNote: sendNote
  }
};

async function sendNote() {
  
  const data = {
    "filename": "randomFileName.md",
    "content": this.inputValue
  }

  console.log(data)

  await axios.post("/notes/", data)
    .then(response => {
      console.log(response)
    })
}

</script>