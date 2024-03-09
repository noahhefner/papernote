<template>
  <v-form>
    <v-container fluid>
      <v-row>
        <v-col
          cols="11"
          align-self="center">
          <v-text-field 
            label="File Name"
            required>
          </v-text-field>
        </v-col>
        <v-col
         cols="1"
         align-self="center">
          <v-btn 
            variant="flat" 
            color="#5865f2" 
            @click="saveNote">
            Save
          </v-btn>
        </v-col>
      </v-row>
    </v-container>
  </v-form>
  <v-row class="bg">
    <v-col cols="6">
      <v-textarea
        v-model="inputValue" 
        @input="handleInputChange"
        auto-grow
        :rows=16
        >
      </v-textarea>
    </v-col>
    <v-col id="rendered" cols="6">
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
