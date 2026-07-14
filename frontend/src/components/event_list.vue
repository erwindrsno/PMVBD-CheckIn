<script setup>
import { ref } from 'vue';

// 1. Add these state variables so the template knows what "dialog" and "newEvent" are
const dialog = ref(false);
const newEvent = ref({ name: '' });

const headers = [
  { title: 'No.', key: 'id' },
  { title: 'Name', key: 'name' },
  { title: 'Status', key: 'status' },
  { title: 'Action', key: 'action' }
];

const events = ref([
  { id: 1, name: 'Science Fair', status: 'Upcoming', action: '-' }
]);

// 2. Define the saveEvent function
const saveEvent = () => {
  const newId = events.value.length + 1;
  events.value.push({
    id: newId,
    name: newEvent.value.name,
    status: 'Upcoming',
    action: '-'
  });

  // Reset and close
  newEvent.value = { name: '' };
  dialog.value = false;
};
</script>

<template>
  <div class="d-flex justify-space-between align-center mb-4">
    <h1>Event List</h1>
    <!-- Now 'dialog' is defined, so this will work -->
    <v-btn color="primary" @click="dialog = true">Add New Event</v-btn>
  </div>

  <v-data-table :headers="headers" :items="events" class="elevation-1"></v-data-table>

  <v-dialog v-model="dialog" max-width="500px">
    <v-card>
      <v-card-title>Add New Event</v-card-title>
      <v-card-text>
        <v-text-field v-model="newEvent.name" label="Event Name"></v-text-field>
      </v-card-text>
      <v-card-actions>
        <v-spacer></v-spacer>
        <v-btn color="grey" @click="dialog = false">Cancel</v-btn>
        <v-btn color="primary" @click="saveEvent">Save</v-btn>
      </v-card-actions>
    </v-card>
  </v-dialog>
</template>
