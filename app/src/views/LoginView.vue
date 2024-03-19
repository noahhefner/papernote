<template>
  <v-form>
    <v-container fluid>
      <v-row
        justify="center">
        <v-col
          cols="3"
          align="center"
          align-self="center">
          <v-text-field 
            label="Username"
            required
            v-model="username">
          </v-text-field>
        </v-col>
      </v-row>
      <v-row
        justify="center">
        <v-col
         cols="3"
         align="center"
         align-self="center">
          <v-text-field 
            label="Password"
            required
            v-model="password">
          </v-text-field>
        </v-col>
      </v-row>
      <v-row
        justify="center">
        <v-col
          cols="3"
          align="center"
          align-self="center">
          <v-btn 
            variant="flat"
            color="#5865f2"
            @click="loginUser">
            Login
          </v-btn>
        </v-col>
      </v-row>
    </v-container>
  </v-form>
</template>

<style>

</style>

<script>
import axios from 'axios';
import router from '@/router'; // Import Vue Router instance

export default {
  data() {
    return {
      username: '',
      password: '',
    };
  },
  methods: {
    async loginUser() {
      try {
        const response = await axios.post('http://0.0.0.0:8080/login', {
          username: this.username,
          password: this.password,
        });
        
        // Assuming the JWT token is returned in the response
        const jwtToken = response.data.token;

        // Save the JWT token to localStorage
        localStorage.setItem('token', jwtToken);

        router.push("/notes")

      } catch (error) {
        console.error('Error logging in:', error);
      }
    },
  },
};
</script>
