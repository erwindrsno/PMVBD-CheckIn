<script setup>
import Header from '@/components/header/Header.vue'
</script>

<template>
  <v-app>
    <Header />

    <v-main>
      <v-container fluid>
        <router-view />
      </v-container>
    </v-main>
  </v-app>
</template>
