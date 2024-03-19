<template>
    <div class="notes">
      <h2>Your Notes</h2>
      <ul>
        <li v-for="(note, index) in notes" :key="index">{{ note }}</li>
      </ul>
    </div>
  </template>
  
  <script>
  import axios from 'axios';
  
  export default {
    data() {
      return {
        notes: [],
      };
    },
    created() {
      // Fetch notes when the component is created
      this.fetchNotes();
    },
    methods: {
      async fetchNotes() {
        try {
          // Make a GET request to the /notes endpoint with the JWT token in the Authorization header
          const token = localStorage.getItem('token'); // Assuming the JWT token is stored in localStorage
          const response = await axios.get('http://localhost:8080/notes', {
            headers: {
              Authorization: `Bearer ${token}`,
            },
          });
          // Set the notes received from the backend in the component's data
          this.notes = response.data;
          
        } catch (error) {
          console.error('Error fetching notes:', error);
        }
      },
    },
  };
  </script>
  
  <style scoped>
  .notes {
    max-width: 600px;
    margin: 0 auto;
  }
  
  ul {
    list-style-type: none;
    padding: 0;
  }
  
  li {
    margin-bottom: 10px;
    border: 1px solid #ccc;
    padding: 10px;
    border-radius: 5px;
  }
  </style>
  