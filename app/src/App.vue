<template>
  <v-app>
    <v-app-bar app>
      <v-app-bar-title>Notes</v-app-bar-title>
      <v-btn variant="flat" color="#5865f2" @click="saveNote">Save</v-btn>
    </v-app-bar>
    <v-main>
      <v-container fluid>
        <router-view></router-view>
      </v-container>
    </v-main>
  </v-app>
</template>

<style>

#app {
  width: 100%;
}

</style>
