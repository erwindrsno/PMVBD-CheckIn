<script setup>
  import { ref } from 'vue';
  const headers = [
    { title: 'Name', key: 'name' },
    { title: 'Email', key: 'email' },
    { title: 'School', key: 'school' },
    { title: 'Action', key: 'action' },
  ];
  const attendees = ref([
    { name: 'Budi Santoso', email: 'budi@example.com', school: 'SMA Negeri 1 Karimun', action: '-' },
  ]);
</script>

<template>
  <h1>Attendee List</h1>
  <v-data-table :headers="headers" :items="attendees" class="elevation-1"></v-data-table>
</template>
