<template>
  <v-row>
    <v-col cols="12" sm="6" md="6" lg="6">
      <v-textarea
        v-model="inputValue" 
        @input="handleInputChange"
        auto-grow>
      </v-textarea>
    </v-col>
    <v-col id="rendered" cols="6" sm="12" md="6" lg="6">
    </v-col>
  </v-row>
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
