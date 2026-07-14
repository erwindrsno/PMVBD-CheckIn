<script setup>
import { ref, onMounted } from 'vue';

// --- State ---
const dialog = ref(false);
const loading = ref(false);
const events = ref([]);
const newEvent = ref({ name: '' });

const headers = [
  { title: 'No.', key: 'idx' }, // Uses index slot
  { title: 'Name', key: 'name' },
  { title: 'Status', key: 'status' },
  { title: 'Created at', key: 'created_at' },
  { title: 'Started at', key: 'started_at' },
  { title: 'Action', key: 'action' }
];

// --- API Logic ---
const fetchEvents = async () => {
  loading.value = true;
  try {
    // Replace with your actual events endpoint
    const response = await fetch('http://localhost:8080/api/v1/events');
    const result = await response.json();
    events.value = result.data.events || [];
  } catch (error) {
    console.error('Error fetching events:', error);
  } finally {
    loading.value = false;
  }
};

const saveEvent = async () => {
  try {
    const response = await fetch('http://localhost:8080/api/v1/events', {
      method: 'POST',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify(newEvent.value),
    });

    if (response.ok) {
      await fetchEvents(); // Refresh list
      dialog.value = false;
      newEvent.value = { name: '' };
    }
  } catch (error) {
    console.error('Error saving event:', error);
  }
};

// --- Lifecycle ---
onMounted(() => {
  fetchEvents();
});
</script>

<template>
  <div class="d-flex justify-space-between align-center mb-4">
    <h1>Event List</h1>
    <v-btn color="primary" @click="dialog = true">Add New Event</v-btn>
  </div>

  <v-data-table
    :headers="headers"
    :items="events"
    :loading="loading"
    class="elevation-1"
  >
    <!-- Incrementing Index -->
    <template v-slot:item.idx="{ index }">
      {{ index + 1 }}
    </template>

    <!-- Handling Name mapping (Assuming Go struct has `json:"name"`) -->
    <template v-slot:item.action="{ item }">
      <v-btn size="small" variant="text" color="error">Delete</v-btn>
    </template>
  </v-data-table>

  <!-- Add Event Dialog -->
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
