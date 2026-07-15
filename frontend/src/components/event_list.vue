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

const statusMap = {
  0: { label: 'New', color: 'warning' },
  1: { label: 'Active', color: 'primary' },
  2: { label: 'Completed', color: 'success' },
  3: { label: 'Cancelled', color: 'error' },
};

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
    <template v-slot:header.status>
      <div class="text-center">Status</div>
    </template>

    <template v-slot:header.action>
      <div class="text-center">Action</div>
    </template>
    <!-- Incrementing Index -->
    <template v-slot:item.idx="{ index }">
      {{ index + 1 }}
    </template>

    <template v-slot:item.status="{ item }">
      <div class="d-flex justify-center">
        <v-chip
          :color="statusMap[item.status]?.color || 'grey'"
          size="small"
          variant="tonal"
        >
          {{ statusMap[item.status]?.label || 'Unknown' }}
        </v-chip>
      </div>
    </template>

    <!-- Refactored Action Column -->
    <template v-slot:item.action="{ item }">
      <div class="d-flex justify-center gap-2">
        <v-btn
          disabled
          icon
          variant="text"
          size="small"
          color="primary"
          @click="editEvent(item)"
        >
         <v-icon>mdi-pencil</v-icon>
        </v-btn>
        <v-btn
          icon
          variant="text"
          size="small"
          color="error"
          @click="deleteEvent(item.id)"
        >
        <v-icon>mdi-delete</v-icon>
        </v-btn>
      </div>
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
